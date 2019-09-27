// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import catalog "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/catalog"
import context "context"
import core "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"

import handler "github.com/lyft/flytepropeller/pkg/controller/nodes/handler"
import io "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/io"
import mock "github.com/stretchr/testify/mock"

// TaskNodeHandler is an autogenerated mock type for the TaskNodeHandler type
type TaskNodeHandler struct {
	mock.Mock
}

// Abort provides a mock function with given fields: ctx, executionContext
func (_m *TaskNodeHandler) Abort(ctx context.Context, executionContext handler.NodeExecutionContext) error {
	ret := _m.Called(ctx, executionContext)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, handler.NodeExecutionContext) error); ok {
		r0 = rf(ctx, executionContext)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Finalize provides a mock function with given fields: ctx, executionContext
func (_m *TaskNodeHandler) Finalize(ctx context.Context, executionContext handler.NodeExecutionContext) error {
	ret := _m.Called(ctx, executionContext)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, handler.NodeExecutionContext) error); ok {
		r0 = rf(ctx, executionContext)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FinalizeRequired provides a mock function with given fields:
func (_m *TaskNodeHandler) FinalizeRequired() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Handle provides a mock function with given fields: ctx, executionContext
func (_m *TaskNodeHandler) Handle(ctx context.Context, executionContext handler.NodeExecutionContext) (handler.Transition, error) {
	ret := _m.Called(ctx, executionContext)

	var r0 handler.Transition
	if rf, ok := ret.Get(0).(func(context.Context, handler.NodeExecutionContext) handler.Transition); ok {
		r0 = rf(ctx, executionContext)
	} else {
		r0 = ret.Get(0).(handler.Transition)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, handler.NodeExecutionContext) error); ok {
		r1 = rf(ctx, executionContext)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Setup provides a mock function with given fields: ctx, setupContext
func (_m *TaskNodeHandler) Setup(ctx context.Context, setupContext handler.SetupContext) error {
	ret := _m.Called(ctx, setupContext)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, handler.SetupContext) error); ok {
		r0 = rf(ctx, setupContext)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateOutputAndCacheAdd provides a mock function with given fields: ctx, i, r, tr, m
func (_m *TaskNodeHandler) ValidateOutputAndCacheAdd(ctx context.Context, i io.InputReader, r io.OutputReader, tr core.TaskReader, m catalog.Metadata) (*io.ExecutionError, error) {
	ret := _m.Called(ctx, i, r, tr, m)

	var r0 *io.ExecutionError
	if rf, ok := ret.Get(0).(func(context.Context, io.InputReader, io.OutputReader, core.TaskReader, catalog.Metadata) *io.ExecutionError); ok {
		r0 = rf(ctx, i, r, tr, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*io.ExecutionError)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, io.InputReader, io.OutputReader, core.TaskReader, catalog.Metadata) error); ok {
		r1 = rf(ctx, i, r, tr, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}