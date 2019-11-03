package utils

import (
	"testing"
	"reflect"
	"net/http"
	"net/http/httptest"
	"strings"
)

func TestMessage(t *testing.T) {
	testStatus := false
	testMessage := "This is a test!"

	gotMessage := Message(testStatus, testMessage)
	expectedMessage  := map[string]interface{}{"status": testStatus, "message": testMessage}

	if eq := reflect.DeepEqual(gotMessage, expectedMessage); !eq {
		t.Errorf("Message does not provide required output! got: %v want: %v", gotMessage, expectedMessage)
	}
}

//DummyHandlerForRespond is just a helper HandleFunc to be able to test Respong package for simulating http traffic
func DummyHandlerForRespond(w http.ResponseWriter, r *http.Request) {
	testStatus := false
        testMessage := "This is a test!"

        expectedMessage  := map[string]interface{}{"status": testStatus, "message": testMessage}

	Respond(w, expectedMessage)
}

func TestRespond(t *testing.T) {
	
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Errorf("New request cannot be created! %v", err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DummyHandlerForRespond)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got: %v want: %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expectedResponse := "{\"message\":\"This is a test!\",\"status\":false}"  //ToDo: Fix this, later convert to json and compare json as string can change like spaces or so
	if strings.Trim(rr.Body.String(), "\n") != expectedResponse {
		t.Errorf("handler returned unexpected body: got:%s want:%s", rr.Body.String(), expectedResponse)
	}
}
