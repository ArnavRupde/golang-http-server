package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserHandler(t *testing.T) {
	requestData := User{
		Name: "AR",
		Age:  20,
	}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		t.Fatal(err)
	}
	reader := strings.NewReader(string(requestBody))

	req, err := http.NewRequest("POST", "/user", reader)

	responseRecorder := httptest.NewRecorder()

	userHandler(responseRecorder, req)

	status := responseRecorder.Result().StatusCode

	if status != http.StatusOK {
		t.Errorf(
			"Incorrect status code %v Expected %v",
			status,
			http.StatusOK,
		)
	}

	var responseData UserCreationResp
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &responseData)
	if err != nil {
		t.Errorf(
			"Error decoding response: %v",
			err,
		)
	}
	if responseData.Status != "Created" {
		t.Errorf(
			"Unexpected status: %s expected %s",
			responseData.Status,
			"Created",
		)
	}
}
