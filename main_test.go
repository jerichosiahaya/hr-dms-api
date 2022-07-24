package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var jsonMap map[string]interface{}
var server *gin.Engine

func performRequest(r http.Handler, method, path string, data io.Reader) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, data)
	// for _, h := range headers {
	// 	req.Header.Add(h.Key, h.Value)
	// }
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.ReleaseMode)
	server = SetupServer()
	code := m.Run()
	os.Exit(code)
}

func TestGetEmployee(t *testing.T) {
	w := performRequest(server, "GET", "/api/v1/employees", nil)
	if err := json.Unmarshal([]byte(w.Body.Bytes()), &jsonMap); err != nil {
		fmt.Println(err.Error())
	}
	assert.Equal(t, 200, w.Code)
	// if recorder.Code != 200 {
	// 	t.Fatalf("bad status code: %d", recorder.Code)
	// }
	// fmt.Println(w.Body)
	// server = SetupServer()
	// w := httptest.NewRecorder()
	// req := httptest.NewRequest("GET", "/api/v1/employees", nil)
	// server.ServeHTTP(w, req)
	// fmt.Print(w.Body)
}

func TestGetEmployeeById(t *testing.T) {
	w := performRequest(server, "GET", "/api/v1/employees/7", nil)
	if err := json.Unmarshal([]byte(w.Body.Bytes()), &jsonMap); err != nil {
		fmt.Println(err.Error())
	}
	assert.Equal(t, 200, w.Code)
}

func TestDeleteEmployeeById(t *testing.T) {
	w := performRequest(server, "DELETE", "/api/v1/employees/7", nil)
	if err := json.Unmarshal([]byte(w.Body.Bytes()), &jsonMap); err != nil {
		fmt.Println(err.Error())
	}
	assert.Equal(t, 200, w.Code)
}

func TestCreateEmployee(t *testing.T) {
	body := strings.NewReader(`{"first_name": "first-name", "last_name": "last_name", "phone_number": 628736261621, "office_id" : 1, "job_id": 1, "hire_date":"2022-07-24"}`)
	w := performRequest(server, "POST", "/api/v1/employees", body)
	if err := json.Unmarshal([]byte(w.Body.Bytes()), &jsonMap); err != nil {
		fmt.Println(err.Error())
	}
	assert.Equal(t, 200, w.Code)
}

func TestUpdateEmployeeById(t *testing.T) {
	body := strings.NewReader(`{"first_name": "first-name-updated", "last_name": "last_name-updated", "phone_number": 628736261621, "office_id" : 1, "job_id": 1, "hire_date":"2022-07-24"}`)
	w := performRequest(server, "PUT", "/api/v1/employees/8", body)
	if err := json.Unmarshal([]byte(w.Body.Bytes()), &jsonMap); err != nil {
		fmt.Println(err.Error())
	}
	assert.Equal(t, 200, w.Code)
}
