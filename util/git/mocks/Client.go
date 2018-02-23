// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Checkout provides a mock function with given fields: repoPath, sha
func (_m *Client) Checkout(repoPath string, sha string) error {
	ret := _m.Called(repoPath, sha)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(repoPath, sha)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CloneOrFetch provides a mock function with given fields: url, username, password, repoPath
func (_m *Client) CloneOrFetch(url string, username string, password string, repoPath string) error {
	ret := _m.Called(url, username, password, repoPath)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(url, username, password, repoPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}