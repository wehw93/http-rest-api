package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"github.com/wehw93/http-rest-api/internal/app/model"
	"github.com/wehw93/http-rest-api/internal/app/store/teststore"
)

func TestServer_HandleUserCreate(t *testing.T) {
	s := newServer(teststore.New(),sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.org",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		}, {
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		}, {
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSessiondCreate(t *testing.T) {
	u := model.TestUser(t)
	store := teststore.New()
	store.User().Create(u)
	s := newServer(store,sessions.NewCookieStore([]byte("hello")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid payload",
			payload:"invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid Email",
			payload: map[string]string{
				"email":    "invalid",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_authentificateUser(t * testing.T){
	store:=teststore.New()
	u:=model.TestUser(t)
	store.User().Create(u)
	testCases:= []struct {
		name string
		cookieValues map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "authentificated",
			cookieValues: map[interface{}]interface{}{
				"user_id":u.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "not authentificated",
			cookieValues: nil,
			expectedCode: http.StatusUnauthorized,
		},
	}
	secretKey:=[]byte("secret")
	s:=newServer(store,sessions.NewCookieStore(secretKey))
	sc:=securecookie.New(secretKey,nil)
	handler:=http.HandlerFunc(func(w http.ResponseWriter,r*http.Request){
		w.WriteHeader(http.StatusOK)
	})
	for _,tc:=range testCases{
		t.Run(tc.name,func(t*testing.T){
			rec:=httptest.NewRecorder()
			req,_:=http.NewRequest(http.MethodGet,"/",nil)
			cookieStr,_ :=sc.Encode(sesssionName,tc.cookieValues)
			req.Header.Set("Cookie",fmt.Sprintf("%s=%s", sesssionName,cookieStr))
			s.authentificateUser(handler).ServeHTTP(rec,req)
			fmt.Printf("Test case: %s\nExpected code: %d, Actual code: %d\n", tc.name, tc.expectedCode, rec.Code)
			fmt.Printf("Request headers: %v\n", req.Header)

			assert.Equal(t,tc.expectedCode,rec.Code)	
		})
	}
}