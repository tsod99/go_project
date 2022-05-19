package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/rs/xid"

	"github.com/tsod99/go_project/api"
	"github.com/tsod99/go_project/db"
)

// CorsMiddleware
func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, DELETE, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, user-token")
		w.Header().Set("Access-Control-Expose-Headers", "user-token")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// HandleListUsers
// @Summary List users
// @Description list users
// @Accept json
// @Produce json
// @Success 200 {object} []api.User
// @Failure 500
// @Failure 405
// @Router /list/users [get]
func HandleListUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error debug %v", r)
			return
		}
	}()

	conn, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	users, err := conn.ListUsers()
	if err != nil {
		panic(err)
	}

	var resp []api.User
	for _, u := range users {
		resp = append(resp, api.User{
			Username: u.Username,
			Password: u.Password,
			Email:    u.Email,
			Group:    u.Group,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

// HandleListGroups
// @Summary List groups
// @Description list groups
// @Accept json
// @Produce json
// @Success 200 {object} []api.Group
// @Failure 500
// @Failure 405
// @Router /list/groups [get]
func HandleListGroups(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error: debug %v", r)
			return
		}
	}()

	conn, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	groups, err := conn.ListGroups()
	if err != nil {
		panic(err)
	}

	var resp []api.Group

	for _, g := range groups {
		resp = append(resp, api.Group{
			Name: g.Name,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}

}

// HandleAddUser
// @Summary Create new user
// @Description create new user
// @Accept json
// @Param user body api.User true "user data"
// @Success 201
// @Failure 400
// @Failure 409
// @Failure 405
// @Failure 500
// @Router /add/user [post]
func HandleAddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("debug error %v", r)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var req api.User
	if err := json.Unmarshal(data, &req); err != nil {
		panic(err)
	}

	conn, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	userID := xid.New().String()
	err = conn.CreateUser(db.User{
		Id:       userID,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Group:    req.Group,
	})

	if err != nil && err == db.ErrConflict {
		w.WriteHeader(http.StatusConflict)
		return
	}

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
}

// HandleAddGroup
// @Summary Create new group
// @Description create new group
// @Accept json
// @Param group body api.Group true "group data"
// @Success 201
// @Failure 400
// @Failure 409
// @Failure 405
// @Failure 500
// @Router /add/group [post]
func HandleAddGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("debug Error %v", r)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var req api.Group

	if err := json.Unmarshal(data, &req); err != nil {
		panic(err)
	}

	conn, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	groupID := xid.New().String()
	err = conn.CreateGroup(db.Group{
		Id:   groupID,
		Name: req.Name,
	})

	if err != nil && err == db.ErrConflict {
		w.WriteHeader(http.StatusConflict)
		return
	}

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
}

// HandleUpdateUser
// @Summary Update user
// @Description update user
// @Accept json
// @Param update body api.UpdateUserRequest true "update user request"
// @Success 200
// @Failure 400
// @Failure 405
// @Failure 500
// @Router /update/user [patch]
func HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("debug error %v", r)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var req api.UpdateUserRequest
	if err := json.Unmarshal(data, &req); err != nil {
		panic(err)
	}

	if !req.Field.Check() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	conn, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if err := conn.UpdateUser(req.Username, req.Field.String(), req.Value); err != nil {
		panic(err)
	}

}

// HandleUpdateGroup
// @Summary Update group
// @Description update group
// @Accept json
// @Param update body api.UpdateGroupRequest true "update group request"
// @Success 200
// @Failure 400
// @Failure 405
// @Failure 500
// @Router /update/group [patch]
func HandleUpdateGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("debug error %v", r)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var req api.UpdateGroupRequest
	if err := json.Unmarshal(data, &req); err != nil {
		panic(err)
	}

	conn, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if err := conn.UpdateGroup(req.GroupName, req.NewGroupName); err != nil {
		panic(err)
	}
}

// HandleDeleteUser
// @Summary Delete User
// @Description delete user
// @Accept json
// @Param delete body api.DeleteUserRequest true "delete user request"
// @Success 200
// @Failure 400
// @Failure 405
// @Failure 500
// @Router /delete/user [delete]
func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("debug error %v", r)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var req api.DeleteUserRequest
	if err := json.Unmarshal(data, &req); err != nil {
		panic(err)
	}

	conn, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if err := conn.DeleteUser(req.Username); err != nil {
		panic(err)
	}

}

// HandleDeleteGroup
// @Summary Delete Group
// @Description delete group
// @Accept json
// @Param delete body api.DeleteGroupRequest true "delete group request"
// @Success 200
// @Failure 400
// @Failure 405
// @Failure 500
// @Router /delete/group [delete]
func HandleDeleteGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("debug error %v", r)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var req api.DeleteGroupRequest
	if err := json.Unmarshal(data, &req); err != nil {
		panic(err)
	}

	conn, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if err := conn.DeleteGroup(req.GroupName); err != nil {
		panic(err)
	}
}
