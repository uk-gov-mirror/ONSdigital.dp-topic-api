package api

import (
	"encoding/json"
	"net/http"

	dprequest "github.com/ONSdigital/dp-net/request"
	errs "github.com/ONSdigital/dp-topic-api/apierrors"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
)

// getTopicPublicHandler is a handler that gets a topic by its id from MongoDB
func (api *API) getTopicPublicHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	id := vars["id"]
	logdata := log.Data{
		"request_id": ctx.Value(dprequest.RequestIdKey),
		"topic_id":   id,
		"function":   "getTopicPublicHandler",
	}

	// get topic from mongoDB by id
	topic, err := api.dataStore.Backend.GetTopic(id)
	if err != nil {
		handleError(ctx, w, err, logdata)
		return
	}

	// Ensure the sub document has the main document ID
	topic.Current.ID = topic.ID

	// User is not authenticated and hence has only access to current sub document
	if err := WriteJSONBody(ctx, topic.Current, w, logdata); err != nil {
		return
	}
	log.Event(ctx, "request successful", log.INFO, logdata) // NOTE: name of function is in logdata
	// NOTE 1st log.Event() in CheckIdentity() needs removing, that looks like:
	// log.Event(ctx, "checking for an identity in request context", log.HTTP(r, 0, 0, nil, nil), logData)
}

// getTopicPrivateHandler is a handler that gets a topic by its id from MongoDB
func (api *API) getTopicPrivateHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	id := vars["id"]
	logdata := log.Data{
		"request_id": ctx.Value(dprequest.RequestIdKey),
		"topic_id":   id,
		"function":   "getTopicPrivateHandler",
	}

	// get topic from mongoDB by id
	topic, err := api.dataStore.Backend.GetTopic(id)
	if err != nil {
		handleError(ctx, w, err, logdata)
		return
	}

	// User has valid authentication to get raw topic document
	if err := WriteJSONBody(ctx, topic, w, logdata); err != nil {
		return
	}
	log.Event(ctx, "request successful", log.INFO, logdata) // NOTE: name of function is in logdata
}

func (api *API) getDataset(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	id := vars["id"]
	logdata := log.Data{
		"request_id": ctx.Value(dprequest.RequestIdKey),
		"topic_id":   id,
		"function":   "getDataset",
	}

	b, err := func() ([]byte, error) {
		dataset, err := api.dataStore.Backend.GetTopic(id)
		if err != nil {
			log.Event(ctx, "getDataset endpoint: dataStore.Backend.GetDataset returned an error", log.ERROR, log.Error(err), logdata)
			return nil, err
		}

		authorised := api.authenticate(req, logdata)

		var b []byte
		var datasetResponse interface{}

		if !authorised {
			// User is not authenticated and hence has only access to current sub document
			if dataset.Current == nil {
				log.Event(ctx, "getDataste endpoint: published dataset not found", log.INFO, logdata)
				return nil, errs.ErrTopicNotFound
			}

			log.Event(ctx, "getDataset endpoint: caller not authenticated returning dataset current sub document", log.INFO, logdata)

			dataset.Current.ID = dataset.ID
			datasetResponse = dataset.Current
		} else {
			// User has valid authentication to get raw dataset document
			if dataset == nil {
				log.Event(ctx, "getDataset endpoint: published or unpublished dataset not found", log.INFO, logdata)
				return nil, errs.ErrTopicNotFound
			}
			log.Event(ctx, "getDataset endpoint: caller authenticated returning dataset", log.INFO, logdata)
			datasetResponse = dataset
		}

		b, err = json.Marshal(datasetResponse)
		if err != nil {
			log.Event(ctx, "getDataset endpoint: failed to marshal dataset resource into bytes", log.ERROR, log.Error(err), logdata)
			return nil, err
		}

		return b, nil
	}()

	if err != nil {
		handleError(ctx, w, err, logdata)
		return
	}

	setJSONContentType(w)
	if _, err = w.Write(b); err != nil {
		log.Event(ctx, "getDataset endpoint: error writing bytes to response", log.ERROR, log.Error(err), logdata)
		handleError(ctx, w, err, logdata)
	}
	log.Event(ctx, "getDataset endpoint: request successful", log.INFO, logdata)
}

func setJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func (api *API) authenticate(r *http.Request, logData log.Data) bool {
	var authenticated bool

	if api.enablePrivateEndpoints {
		var hasCallerIdentity, hasUserIdentity bool

		// NOTE:
		// If the identity exists then the user has been authenticated.
		// There is an earlier step in the middleware which will call off to zebedee to
		// authenticate the request (user/service) and this will add the identity to the
		// request context for later use in the application ...
		// ... which happens to be here:

		callerIdentity := dprequest.Caller(r.Context())
		if callerIdentity != "" {
			logData["caller_identity"] = callerIdentity
			hasCallerIdentity = true
		}

		userIdentity := dprequest.User(r.Context())
		if userIdentity != "" {
			logData["user_identity"] = userIdentity
			hasUserIdentity = true
		}

		if hasCallerIdentity || hasUserIdentity {
			authenticated = true
		}
		logData["authenticated"] = authenticated
	}
	return authenticated
}