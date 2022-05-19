package api_testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tsod99/go_project/api"
	"github.com/tsod99/go_project/handlers"
)

var testGroup = api.Group{
	Name: "test-group",
}

var testUser = api.User{
	Username: "test-user",
	Password: "test-password",
	Email:    "test-email@foo.com",
	Group:    "test-group",
}

func TestAddGroup(t *testing.T) {
	data, err := json.Marshal(testGroup)
	if err != nil {
		t.Logf("JSON marshal %v", err)
		t.Fail()
		return
	}
	buff := bytes.NewBuffer(data)
	r, err := http.NewRequest(http.MethodPost, "/add/group", buff)
	if err != nil {
		t.Logf("NewRequest: Error %v", err)
		t.Fail()
		return
	}

	w := httptest.NewRecorder()
	handlers.HandleAddGroup(w, r)
	if w.Code != http.StatusCreated {
		t.Logf("failed to create new group")
		t.Fail()
		return
	}

	t.Logf("successfuly created a new group")
}

func TestAddUser(t *testing.T) {
	data, err := json.Marshal(testUser)
	if err != nil {
		t.Logf("JSON marshal %v", err)
		t.Fail()
		return
	}
	buff := bytes.NewBuffer(data)

	r, err := http.NewRequest(http.MethodPost, "/add/user", buff)
	if err != nil {
		t.Logf("NewRequest: Error %v", err)
		t.Fail()
		return
	}

	w := httptest.NewRecorder()
	handlers.HandleAddUser(w, r)

	if w.Code != http.StatusCreated {
		t.Logf("failed to create new user")
		t.Fail()
		return
	}

	t.Logf("successfuly created new user")
}

func TestUpdateUser(t *testing.T) {
	req := api.UpdateUserRequest{
		Username: testUser.Username,
		Field:    api.FieldEmail,
		Value:    "test-user@bar.com",
	}
	data, err := json.Marshal(req)
	if err != nil {
		t.Logf("JSON: Error %v", err)
		t.Fail()
		return
	}
	buff := bytes.NewBuffer(data)
	r, err := http.NewRequest(http.MethodPatch, "/update/user", buff)
	if err != nil {
		t.Logf("NewRequest: Error %v", err)
		t.Fail()
		return
	}

	w := httptest.NewRecorder()
	handlers.HandleUpdateUser(w, r)

	if w.Code != http.StatusOK {
		t.Logf("failed to update user %v %v", w.Code, w.Body.String())
		t.Fail()
		return
	}

	t.Logf("successfuly update user")
}

func TestUpdateGroup(t *testing.T) {
	req := api.UpdateGroupRequest{
		GroupName:    testGroup.Name,
		NewGroupName: "test-name",
	}
	testGroup.Name = "test-name"
	data, err := json.Marshal(req)
	if err != nil {
		t.Logf("JSON: Error %v", err)
		t.Fail()
		return
	}
	buff := bytes.NewBuffer(data)
	r, err := http.NewRequest(http.MethodPatch, "/update/group", buff)
	if err != nil {
		t.Logf("NewRequest: Error %v", err)
		t.Fail()
		return
	}

	w := httptest.NewRecorder()
	handlers.HandleUpdateGroup(w, r)

	if w.Code != http.StatusOK {
		t.Logf("failed to update group %v %v", w.Code, w.Body.String())
		t.Fail()
		return
	}

	t.Logf("successfuly update group")
}

func TestDeleteUser(t *testing.T) {
	data, err := json.Marshal(api.DeleteUserRequest{
		Username: testUser.Username,
	})
	if err != nil {
		t.Logf("JSON marshal %v", err)
		t.Fail()
		return
	}
	buff := bytes.NewBuffer(data)

	r, err := http.NewRequest(http.MethodDelete, "/delete/user", buff)
	if err != nil {
		t.Logf("NewRequest: Error %v", err)
		t.Fail()
		return
	}

	w := httptest.NewRecorder()
	handlers.HandleDeleteUser(w, r)

	if w.Code != http.StatusOK {
		t.Logf("failed to delete user %v %v", w.Code, w.Body.String())
		t.Fail()
		return
	}

	t.Logf("successfuly deleted user")
}

func TestDeleteGroup(t *testing.T) {
	data, err := json.Marshal(api.DeleteGroupRequest{
		GroupName: testGroup.Name,
	})
	if err != nil {
		t.Logf("JSON marshal %v", err)
		t.Fail()
		return
	}
	buff := bytes.NewBuffer(data)

	r, err := http.NewRequest(http.MethodDelete, "/delete/group", buff)
	if err != nil {
		t.Logf("NewRequest: Error %v", err)
		t.Fail()
		return
	}

	w := httptest.NewRecorder()
	handlers.HandleDeleteGroup(w, r)

	if w.Code != http.StatusOK {
		t.Logf("failed to delete group %v %v", w.Code, w.Body.String())
		t.Fail()
		return
	}

	t.Logf("successfuly deleted group")
}
