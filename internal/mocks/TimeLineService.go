// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	auth "github.com/bangumi/server/internal/auth"
	collection "github.com/bangumi/server/internal/collections/domain/collection"

	episode "github.com/bangumi/server/internal/episode"

	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"
)

// TimeLineService is an autogenerated mock type for the Service type
type TimeLineService struct {
	mock.Mock
}

type TimeLineService_Expecter struct {
	mock *mock.Mock
}

func (_m *TimeLineService) EXPECT() *TimeLineService_Expecter {
	return &TimeLineService_Expecter{mock: &_m.Mock}
}

// ChangeEpisodeStatus provides a mock function with given fields: ctx, u, sbj, _a3
func (_m *TimeLineService) ChangeEpisodeStatus(ctx context.Context, u auth.Auth, sbj model.Subject, _a3 episode.Episode) error {
	ret := _m.Called(ctx, u, sbj, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, auth.Auth, model.Subject, episode.Episode) error); ok {
		r0 = rf(ctx, u, sbj, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TimeLineService_ChangeEpisodeStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChangeEpisodeStatus'
type TimeLineService_ChangeEpisodeStatus_Call struct {
	*mock.Call
}

// ChangeEpisodeStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - u auth.Auth
//   - sbj model.Subject
//   - _a3 episode.Episode
func (_e *TimeLineService_Expecter) ChangeEpisodeStatus(ctx interface{}, u interface{}, sbj interface{}, _a3 interface{}) *TimeLineService_ChangeEpisodeStatus_Call {
	return &TimeLineService_ChangeEpisodeStatus_Call{Call: _e.mock.On("ChangeEpisodeStatus", ctx, u, sbj, _a3)}
}

func (_c *TimeLineService_ChangeEpisodeStatus_Call) Run(run func(ctx context.Context, u auth.Auth, sbj model.Subject, _a3 episode.Episode)) *TimeLineService_ChangeEpisodeStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(auth.Auth), args[2].(model.Subject), args[3].(episode.Episode))
	})
	return _c
}

func (_c *TimeLineService_ChangeEpisodeStatus_Call) Return(_a0 error) *TimeLineService_ChangeEpisodeStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

// ChangeSubjectCollection provides a mock function with given fields: ctx, u, sbj, collect, comment, rate
func (_m *TimeLineService) ChangeSubjectCollection(ctx context.Context, u uint32, sbj model.Subject, collect collection.SubjectCollection, comment string, rate uint8) error {
	ret := _m.Called(ctx, u, sbj, collect, comment, rate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, model.Subject, collection.SubjectCollection, string, uint8) error); ok {
		r0 = rf(ctx, u, sbj, collect, comment, rate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TimeLineService_ChangeSubjectCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChangeSubjectCollection'
type TimeLineService_ChangeSubjectCollection_Call struct {
	*mock.Call
}

// ChangeSubjectCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - u uint32
//   - sbj model.Subject
//   - collect collection.SubjectCollection
//   - comment string
//   - rate uint8
func (_e *TimeLineService_Expecter) ChangeSubjectCollection(ctx interface{}, u interface{}, sbj interface{}, collect interface{}, comment interface{}, rate interface{}) *TimeLineService_ChangeSubjectCollection_Call {
	return &TimeLineService_ChangeSubjectCollection_Call{Call: _e.mock.On("ChangeSubjectCollection", ctx, u, sbj, collect, comment, rate)}
}

func (_c *TimeLineService_ChangeSubjectCollection_Call) Run(run func(ctx context.Context, u uint32, sbj model.Subject, collect collection.SubjectCollection, comment string, rate uint8)) *TimeLineService_ChangeSubjectCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(model.Subject), args[3].(collection.SubjectCollection), args[4].(string), args[5].(uint8))
	})
	return _c
}

func (_c *TimeLineService_ChangeSubjectCollection_Call) Return(_a0 error) *TimeLineService_ChangeSubjectCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

// ChangeSubjectProgress provides a mock function with given fields: ctx, u, sbj, epsUpdate, volsUpdate
func (_m *TimeLineService) ChangeSubjectProgress(ctx context.Context, u uint32, sbj model.Subject, epsUpdate uint32, volsUpdate uint32) error {
	ret := _m.Called(ctx, u, sbj, epsUpdate, volsUpdate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, model.Subject, uint32, uint32) error); ok {
		r0 = rf(ctx, u, sbj, epsUpdate, volsUpdate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TimeLineService_ChangeSubjectProgress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChangeSubjectProgress'
type TimeLineService_ChangeSubjectProgress_Call struct {
	*mock.Call
}

// ChangeSubjectProgress is a helper method to define mock.On call
//   - ctx context.Context
//   - u uint32
//   - sbj model.Subject
//   - epsUpdate uint32
//   - volsUpdate uint32
func (_e *TimeLineService_Expecter) ChangeSubjectProgress(ctx interface{}, u interface{}, sbj interface{}, epsUpdate interface{}, volsUpdate interface{}) *TimeLineService_ChangeSubjectProgress_Call {
	return &TimeLineService_ChangeSubjectProgress_Call{Call: _e.mock.On("ChangeSubjectProgress", ctx, u, sbj, epsUpdate, volsUpdate)}
}

func (_c *TimeLineService_ChangeSubjectProgress_Call) Run(run func(ctx context.Context, u uint32, sbj model.Subject, epsUpdate uint32, volsUpdate uint32)) *TimeLineService_ChangeSubjectProgress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(model.Subject), args[3].(uint32), args[4].(uint32))
	})
	return _c
}

func (_c *TimeLineService_ChangeSubjectProgress_Call) Return(_a0 error) *TimeLineService_ChangeSubjectProgress_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewTimeLineService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTimeLineService creates a new instance of TimeLineService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTimeLineService(t mockConstructorTestingTNewTimeLineService) *TimeLineService {
	mock := &TimeLineService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}