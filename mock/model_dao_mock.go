// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: model_dao.go

package mock

import (
	sync "sync"

	github_com_google_uuid "github.com/google/uuid"
	github_com_stkr89_modelsvc_models "github.com/stkr89/go-k8s-cognito-crud-microservice-template/models"
)

// MockModelDao is a mock of ModelDao interface
type MockModelDao struct {
	lockCreate sync.Mutex
	CreateFunc func(obj *github_com_stkr89_modelsvc_models.Model) (*github_com_stkr89_modelsvc_models.Model, error)

	lockGet sync.Mutex
	GetFunc func(id github_com_google_uuid.UUID) (*github_com_stkr89_modelsvc_models.Model, error)

	lockList sync.Mutex
	ListFunc func() ([]*github_com_stkr89_modelsvc_models.Model, error)

	lockUpdate sync.Mutex
	UpdateFunc func(obj *github_com_stkr89_modelsvc_models.Model) (*github_com_stkr89_modelsvc_models.Model, error)

	lockDelete sync.Mutex
	DeleteFunc func(id github_com_google_uuid.UUID) error

	calls struct {
		Create []struct {
			Obj *github_com_stkr89_modelsvc_models.Model
		}
		Get []struct {
			Id github_com_google_uuid.UUID
		}
		List []struct {
		}
		Update []struct {
			Obj *github_com_stkr89_modelsvc_models.Model
		}
		Delete []struct {
			Id github_com_google_uuid.UUID
		}
	}
}

// Create mocks base method by wrapping the associated func.
func (m *MockModelDao) Create(obj *github_com_stkr89_modelsvc_models.Model) (*github_com_stkr89_modelsvc_models.Model, error) {
	m.lockCreate.Lock()
	defer m.lockCreate.Unlock()

	if m.CreateFunc == nil {
		panic("mocker: MockModelDao.CreateFunc is nil but MockModelDao.Create was called.")
	}

	call := struct {
		Obj *github_com_stkr89_modelsvc_models.Model
	}{
		Obj: obj,
	}

	m.calls.Create = append(m.calls.Create, call)

	return m.CreateFunc(obj)
}

// CreateCalled returns true if Create was called at least once.
func (m *MockModelDao) CreateCalled() bool {
	m.lockCreate.Lock()
	defer m.lockCreate.Unlock()

	return len(m.calls.Create) > 0
}

// CreateCalls returns the calls made to Create.
func (m *MockModelDao) CreateCalls() []struct {
	Obj *github_com_stkr89_modelsvc_models.Model
} {
	m.lockCreate.Lock()
	defer m.lockCreate.Unlock()

	return m.calls.Create
}

// Get mocks base method by wrapping the associated func.
func (m *MockModelDao) Get(id github_com_google_uuid.UUID) (*github_com_stkr89_modelsvc_models.Model, error) {
	m.lockGet.Lock()
	defer m.lockGet.Unlock()

	if m.GetFunc == nil {
		panic("mocker: MockModelDao.GetFunc is nil but MockModelDao.Get was called.")
	}

	call := struct {
		Id github_com_google_uuid.UUID
	}{
		Id: id,
	}

	m.calls.Get = append(m.calls.Get, call)

	return m.GetFunc(id)
}

// GetCalled returns true if Get was called at least once.
func (m *MockModelDao) GetCalled() bool {
	m.lockGet.Lock()
	defer m.lockGet.Unlock()

	return len(m.calls.Get) > 0
}

// GetCalls returns the calls made to Get.
func (m *MockModelDao) GetCalls() []struct {
	Id github_com_google_uuid.UUID
} {
	m.lockGet.Lock()
	defer m.lockGet.Unlock()

	return m.calls.Get
}

// List mocks base method by wrapping the associated func.
func (m *MockModelDao) List() ([]*github_com_stkr89_modelsvc_models.Model, error) {
	m.lockList.Lock()
	defer m.lockList.Unlock()

	if m.ListFunc == nil {
		panic("mocker: MockModelDao.ListFunc is nil but MockModelDao.List was called.")
	}

	call := struct {
	}{}

	m.calls.List = append(m.calls.List, call)

	return m.ListFunc()
}

// ListCalled returns true if List was called at least once.
func (m *MockModelDao) ListCalled() bool {
	m.lockList.Lock()
	defer m.lockList.Unlock()

	return len(m.calls.List) > 0
}

// ListCalls returns the calls made to List.
func (m *MockModelDao) ListCalls() []struct {
} {
	m.lockList.Lock()
	defer m.lockList.Unlock()

	return m.calls.List
}

// Update mocks base method by wrapping the associated func.
func (m *MockModelDao) Update(obj *github_com_stkr89_modelsvc_models.Model) (*github_com_stkr89_modelsvc_models.Model, error) {
	m.lockUpdate.Lock()
	defer m.lockUpdate.Unlock()

	if m.UpdateFunc == nil {
		panic("mocker: MockModelDao.UpdateFunc is nil but MockModelDao.Update was called.")
	}

	call := struct {
		Obj *github_com_stkr89_modelsvc_models.Model
	}{
		Obj: obj,
	}

	m.calls.Update = append(m.calls.Update, call)

	return m.UpdateFunc(obj)
}

// UpdateCalled returns true if Update was called at least once.
func (m *MockModelDao) UpdateCalled() bool {
	m.lockUpdate.Lock()
	defer m.lockUpdate.Unlock()

	return len(m.calls.Update) > 0
}

// UpdateCalls returns the calls made to Update.
func (m *MockModelDao) UpdateCalls() []struct {
	Obj *github_com_stkr89_modelsvc_models.Model
} {
	m.lockUpdate.Lock()
	defer m.lockUpdate.Unlock()

	return m.calls.Update
}

// Delete mocks base method by wrapping the associated func.
func (m *MockModelDao) Delete(id github_com_google_uuid.UUID) error {
	m.lockDelete.Lock()
	defer m.lockDelete.Unlock()

	if m.DeleteFunc == nil {
		panic("mocker: MockModelDao.DeleteFunc is nil but MockModelDao.Delete was called.")
	}

	call := struct {
		Id github_com_google_uuid.UUID
	}{
		Id: id,
	}

	m.calls.Delete = append(m.calls.Delete, call)

	return m.DeleteFunc(id)
}

// DeleteCalled returns true if Delete was called at least once.
func (m *MockModelDao) DeleteCalled() bool {
	m.lockDelete.Lock()
	defer m.lockDelete.Unlock()

	return len(m.calls.Delete) > 0
}

// DeleteCalls returns the calls made to Delete.
func (m *MockModelDao) DeleteCalls() []struct {
	Id github_com_google_uuid.UUID
} {
	m.lockDelete.Lock()
	defer m.lockDelete.Unlock()

	return m.calls.Delete
}

// Reset resets the calls made to the mocked methods.
func (m *MockModelDao) Reset() {
	m.lockCreate.Lock()
	m.calls.Create = nil
	m.lockCreate.Unlock()
	m.lockGet.Lock()
	m.calls.Get = nil
	m.lockGet.Unlock()
	m.lockList.Lock()
	m.calls.List = nil
	m.lockList.Unlock()
	m.lockUpdate.Lock()
	m.calls.Update = nil
	m.lockUpdate.Unlock()
	m.lockDelete.Lock()
	m.calls.Delete = nil
	m.lockDelete.Unlock()
}
