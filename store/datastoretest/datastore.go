// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package storetest

import (
	"github.com/ONSdigital/dp-topic-api/models"
	"github.com/ONSdigital/dp-topic-api/store"
	"sync"
)

// Ensure, that StorerMock does implement store.Storer.
// If this is not the case, regenerate this file with moq.
var _ store.Storer = &StorerMock{}

// StorerMock is a mock implementation of store.Storer.
//
//     func TestSomethingThatUsesStorer(t *testing.T) {
//
//         // make and configure a mocked store.Storer
//         mockedStorer := &StorerMock{
//             CheckTopicExistsFunc: func(id string) error {
// 	               panic("mock out the CheckTopicExists method")
//             },
//             GetContentFunc: func(id string, queryTypeFlags int) (*models.ContentResponse, error) {
// 	               panic("mock out the GetContent method")
//             },
//             GetTopicFunc: func(id string) (*models.TopicResponse, error) {
// 	               panic("mock out the GetTopic method")
//             },
//         }
//
//         // use mockedStorer in code that requires store.Storer
//         // and then make assertions.
//
//     }
type StorerMock struct {
	// CheckTopicExistsFunc mocks the CheckTopicExists method.
	CheckTopicExistsFunc func(id string) error

	// GetContentFunc mocks the GetContent method.
	GetContentFunc func(id string, queryTypeFlags int) (*models.ContentResponse, error)

	// GetTopicFunc mocks the GetTopic method.
	GetTopicFunc func(id string) (*models.TopicResponse, error)

	// calls tracks calls to the methods.
	calls struct {
		// CheckTopicExists holds details about calls to the CheckTopicExists method.
		CheckTopicExists []struct {
			// ID is the id argument value.
			ID string
		}
		// GetContent holds details about calls to the GetContent method.
		GetContent []struct {
			// ID is the id argument value.
			ID string
			// QueryTypeFlags is the queryTypeFlags argument value.
			QueryTypeFlags int
		}
		// GetTopic holds details about calls to the GetTopic method.
		GetTopic []struct {
			// ID is the id argument value.
			ID string
		}
	}
	lockCheckTopicExists sync.RWMutex
	lockGetContent       sync.RWMutex
	lockGetTopic         sync.RWMutex
}

// CheckTopicExists calls CheckTopicExistsFunc.
func (mock *StorerMock) CheckTopicExists(id string) error {
	if mock.CheckTopicExistsFunc == nil {
		panic("StorerMock.CheckTopicExistsFunc: method is nil but Storer.CheckTopicExists was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockCheckTopicExists.Lock()
	mock.calls.CheckTopicExists = append(mock.calls.CheckTopicExists, callInfo)
	mock.lockCheckTopicExists.Unlock()
	return mock.CheckTopicExistsFunc(id)
}

// CheckTopicExistsCalls gets all the calls that were made to CheckTopicExists.
// Check the length with:
//     len(mockedStorer.CheckTopicExistsCalls())
func (mock *StorerMock) CheckTopicExistsCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockCheckTopicExists.RLock()
	calls = mock.calls.CheckTopicExists
	mock.lockCheckTopicExists.RUnlock()
	return calls
}

// GetContent calls GetContentFunc.
func (mock *StorerMock) GetContent(id string, queryTypeFlags int) (*models.ContentResponse, error) {
	if mock.GetContentFunc == nil {
		panic("StorerMock.GetContentFunc: method is nil but Storer.GetContent was just called")
	}
	callInfo := struct {
		ID             string
		QueryTypeFlags int
	}{
		ID:             id,
		QueryTypeFlags: queryTypeFlags,
	}
	mock.lockGetContent.Lock()
	mock.calls.GetContent = append(mock.calls.GetContent, callInfo)
	mock.lockGetContent.Unlock()
	return mock.GetContentFunc(id, queryTypeFlags)
}

// GetContentCalls gets all the calls that were made to GetContent.
// Check the length with:
//     len(mockedStorer.GetContentCalls())
func (mock *StorerMock) GetContentCalls() []struct {
	ID             string
	QueryTypeFlags int
} {
	var calls []struct {
		ID             string
		QueryTypeFlags int
	}
	mock.lockGetContent.RLock()
	calls = mock.calls.GetContent
	mock.lockGetContent.RUnlock()
	return calls
}

// GetTopic calls GetTopicFunc.
func (mock *StorerMock) GetTopic(id string) (*models.TopicResponse, error) {
	if mock.GetTopicFunc == nil {
		panic("StorerMock.GetTopicFunc: method is nil but Storer.GetTopic was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGetTopic.Lock()
	mock.calls.GetTopic = append(mock.calls.GetTopic, callInfo)
	mock.lockGetTopic.Unlock()
	return mock.GetTopicFunc(id)
}

// GetTopicCalls gets all the calls that were made to GetTopic.
// Check the length with:
//     len(mockedStorer.GetTopicCalls())
func (mock *StorerMock) GetTopicCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetTopic.RLock()
	calls = mock.calls.GetTopic
	mock.lockGetTopic.RUnlock()
	return calls
}
