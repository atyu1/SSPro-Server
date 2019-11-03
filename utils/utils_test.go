package utils

import (
	"testing"
	"reflect"
)

func TestMessage(t *testing.T) {
	testStatus := false
	testMessage := "This is a test!"

	gotMessage := Message(testStatus, testMessage)
	expectedMessage  := map[string]interface{}{"status": testStatus, "message": testMessage}

	if eq := reflect.DeepEqual(gotMessage, expectedMessage); !eq {
		t.Errorf("Message does not provide required output! %v|%v", gotMessage, expectedMessage)
	}
}

