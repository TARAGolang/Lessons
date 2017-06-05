// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import models "Lessons/GoBDD/models"

// IGithub is an autogenerated mock type for the IGithub type
type IGithub struct {
	mock.Mock
}

// GetPackageRepoInfo provides a mock function with given fields: owner, repositoryName
func (_m *IGithub) GetPackageRepoInfo(owner string, repositoryName string) (*models.Package, error) {
	ret := _m.Called(owner, repositoryName)

	var r0 *models.Package
	if rf, ok := ret.Get(0).(func(string, string) *models.Package); ok {
		r0 = rf(owner, repositoryName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Package)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(owner, repositoryName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
