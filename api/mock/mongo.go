// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"sync"

	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-topic-api/api"
)

var (
	lockMongoServerMockChecker sync.RWMutex
	lockMongoServerMockClose   sync.RWMutex
)

// Ensure, that MongoServerMock does implement api.MongoServer.
// If this is not the case, regenerate this file with moq.
var _ api.MongoServer = &MongoServerMock{}

// MongoServerMock is a mock implementation of api.MongoServer.
//
//     func TestSomethingThatUsesMongoServer(t *testing.T) {
//
//         // make and configure a mocked api.MongoServer
//         mockedMongoServer := &MongoServerMock{
//             CheckerFunc: func(ctx context.Context, state *healthcheck.CheckState) error {
// 	               panic("mock out the Checker method")
//             },
//             CloseFunc: func(ctx context.Context) error {
// 	               panic("mock out the Close method")
//             },
//         }
//
//         // use mockedMongoServer in code that requires api.MongoServer
//         // and then make assertions.
//
//     }
type MongoServerMock struct {
	// CheckerFunc mocks the Checker method.
	CheckerFunc func(ctx context.Context, state *healthcheck.CheckState) error

	// CloseFunc mocks the Close method.
	CloseFunc func(ctx context.Context) error

	// calls tracks calls to the methods.
	calls struct {
		// Checker holds details about calls to the Checker method.
		Checker []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// State is the state argument value.
			State *healthcheck.CheckState
		}
		// Close holds details about calls to the Close method.
		Close []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
}

// Checker calls CheckerFunc.
func (mock *MongoServerMock) Checker(ctx context.Context, state *healthcheck.CheckState) error {
	if mock.CheckerFunc == nil {
		panic("MongoServerMock.CheckerFunc: method is nil but MongoServer.Checker was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		State *healthcheck.CheckState
	}{
		Ctx:   ctx,
		State: state,
	}
	lockMongoServerMockChecker.Lock()
	mock.calls.Checker = append(mock.calls.Checker, callInfo)
	lockMongoServerMockChecker.Unlock()
	return mock.CheckerFunc(ctx, state)
}

// CheckerCalls gets all the calls that were made to Checker.
// Check the length with:
//     len(mockedMongoServer.CheckerCalls())
func (mock *MongoServerMock) CheckerCalls() []struct {
	Ctx   context.Context
	State *healthcheck.CheckState
} {
	var calls []struct {
		Ctx   context.Context
		State *healthcheck.CheckState
	}
	lockMongoServerMockChecker.RLock()
	calls = mock.calls.Checker
	lockMongoServerMockChecker.RUnlock()
	return calls
}

// Close calls CloseFunc.
func (mock *MongoServerMock) Close(ctx context.Context) error {
	if mock.CloseFunc == nil {
		panic("MongoServerMock.CloseFunc: method is nil but MongoServer.Close was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockMongoServerMockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	lockMongoServerMockClose.Unlock()
	return mock.CloseFunc(ctx)
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedMongoServer.CloseCalls())
func (mock *MongoServerMock) CloseCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockMongoServerMockClose.RLock()
	calls = mock.calls.Close
	lockMongoServerMockClose.RUnlock()
	return calls
}
