package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wehw93/http-rest-api/internal/app/store/teststore"
)

func TestServer_HandleUserCreate(t *testing.T) {
	s := newServer(teststore.New())
	testCases:=[]struct{
		name string
		payload interface{}
		expectedCode int
	}{
		{
		name:"valid",
		payload:map[string]string{
			"email":"user@example.org",
			"password":"password",
		},
		expectedCode: http.StatusCreated,
		},{
			name:"invalid payload",
			payload:"invalid",
			expectedCode:http.StatusBadRequest,
		},{
			name:"invalid params",
			payload:map[string]string{
				"email":"invalid",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}
	for _,tc:=range testCases{
		t.Run(tc.name,func (t * testing.T){
			rec:=httptest.NewRecorder()
			b:=&bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req,_:=http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code) 
		})
	}

	
}
