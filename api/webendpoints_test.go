package api

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	dprequest "github.com/ONSdigital/dp-net/request"
	"github.com/ONSdigital/dp-topic-api/config"
	storetest "github.com/ONSdigital/dp-topic-api/store/mock"

	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	mu              sync.Mutex
	testContext     = context.Background()
	testTopicAPIURL = "http://localhost:1234"
)

func TestPublishedSubnetEndpointsAreDisabled(t *testing.T) {
	type testEndpoint struct {
		Method string
		URL    string
	}

	publishSubnetEndpoints := map[testEndpoint]int{
		// Topics endpoints
		{Method: "GET", URL: "http://localhost:25300/bad-endpoint"}: http.StatusNotFound,
	}

	Convey("When the API is started with private endpoints disabled and permission auth disabled", t, func() {
		for endpoint, expectedStatusCode := range publishSubnetEndpoints {
			Convey("The following endpoint "+endpoint.URL+"(Method:"+endpoint.Method+") should return 404", func() {
				request, err := createRequestWithAuth(endpoint.Method, endpoint.URL, nil)
				So(err, ShouldBeNil)

				cfg, err := config.Get()
				So(err, ShouldBeNil)
				cfg.EnablePrivateEndpoints = false
				cfg.EnablePermissionsAuth = false

				w := httptest.NewRecorder()
				mockedDataStore := &storetest.StorerMock{}
				api := GetAPIWithMocks(cfg, mockedDataStore)

				api.Router.ServeHTTP(w, request)

				So(w.Code, ShouldEqual, expectedStatusCode)
			})
		}
	})
}

func TestSetup(t *testing.T) {
	Convey("Given an API instance", t, func() {
		Convey("When created in Publishing mode", func() {
			cfg, err := config.Get()
			So(err, ShouldBeNil)
			cfg.EnablePrivateEndpoints = true
			mockedDataStore := &storetest.StorerMock{}

			api := GetAPIWithMocks(cfg, mockedDataStore)

			Convey("When created the following routes should have been added", func() {
				So(api, ShouldNotBeNil)
				So(hasRoute(api.Router, "/topics", "GET"), ShouldBeTrue)
				So(hasRoute(api.Router, "/topics/{id}", "GET"), ShouldBeTrue)
				So(hasRoute(api.Router, "/topics/{id}/subtopics", "GET"), ShouldBeTrue)
				So(hasRoute(api.Router, "/topics/{id}/content", "GET"), ShouldBeTrue)
				So(hasRoute(api.Router, "/navigation", "GET"), ShouldBeTrue)
			})
		})

		Convey("When created in Web mode", func() {
			cfg, err := config.Get()
			So(err, ShouldBeNil)
			cfg.EnablePrivateEndpoints = false
			cfg.EnablePermissionsAuth = true

			mockedDataStore := &storetest.StorerMock{}

			api := GetAPIWithMocks(cfg, mockedDataStore)

			Convey("Then the following routes should have been added", func() {
				So(api, ShouldNotBeNil)
				So(hasRoute(api.Router, "/topics", "GET"), ShouldBeTrue)
				So(hasRoute(api.Router, "/topics/{id}", "GET"), ShouldBeTrue)
				So(hasRoute(api.Router, "/topics/{id}/subtopics", "GET"), ShouldBeTrue)
				So(hasRoute(api.Router, "/topics/{id}/content", "GET"), ShouldBeTrue)
				So(hasRoute(api.Router, "/navigation", "GET"), ShouldBeTrue)
			})
		})
	})
}

func hasRoute(r *mux.Router, path, method string) bool {
	req := httptest.NewRequest(method, path, nil)
	match := &mux.RouteMatch{}
	return r.Match(req, match)
}

func createRequestWithAuth(method, url string, body io.Reader) (*http.Request, error) {
	request := httptest.NewRequest(method, url, body)

	ctx := request.Context()
	ctx = dprequest.SetCaller(ctx, "someone@ons.gov.uk")
	request = request.WithContext(ctx)
	return request, nil
}
