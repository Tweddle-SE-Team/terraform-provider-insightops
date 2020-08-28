package insightops

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getTestClient(requestMatcher TestRequestMatcher) *InsightOpsClient {
	testClientServer := TestClientServer{
		RequestMatcher: requestMatcher,
	}
	httpClient, httpServer := testClientServer.TestClientServer()
	c := &InsightOpsClient{InsightOpsUrl: httpServer.URL, ApiKey: "apikey", HttpClient: httpClient}
	return c
}

func TestInsightOpsClient_NewInsightOpsClient(t *testing.T) {
	_, err := NewInsightOpsClient("apiKey", "eu")
	assert.Nil(t, err)
}

func TestInsightOpsClient_NewInsightOpsClientMissing(t *testing.T) {
	_, err := NewInsightOpsClient("", "eu")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "ApiKey is mandatory to initialize Insight client")
	_, err = NewInsightOpsClient("apiKey", "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Region is mandatory to initialize Insight client")
}

type mockObject struct {
	Data string `json:"data"`
}

func TestInsightOpsClient_ClientGet(t *testing.T) {
	mockResponse := &mockObject{Data: "some data..."}
	requestMatcher := NewRequestMatcher(http.MethodGet, "/api/testing", nil, http.StatusOK, mockResponse)
	c := getTestClient(requestMatcher)
	expectedResponse := &mockObject{}
	err := c.get("/api/testing", expectedResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse.Data, mockResponse.Data)
}

func TestInsightOpsClient_ClientGetResponseNotStatusOk(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodGet, "/api/testing", nil, http.StatusUnauthorized, &mockObject{})
	c := getTestClient(requestMatcher)
	err := c.get("/api/testing", &mockObject{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}

func TestInsightOpsClient_ClientPost(t *testing.T) {
	mockRequestPayload := &mockObject{Data: "some req data..."}
	mockResponse := &mockObject{Data: "some data..."}
	requestMatcher := NewRequestMatcher(http.MethodPost, "/api/testing", mockRequestPayload, http.StatusCreated, mockResponse)

	c := getTestClient(requestMatcher)
	expectedResponse := &mockObject{}
	body, err := c.post("/api/testing", mockRequestPayload)
	err = json.Unmarshal(body, &expectedResponse)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse.Data, mockResponse.Data)
}

func TestInsightOpsClient_ClientPostResponseNotStatusCreated(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodPost, "/api/testing", &mockObject{}, http.StatusUnauthorized, &mockObject{})

	c := getTestClient(requestMatcher)
	_, err := c.post("/api/testing", &mockObject{})

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}

func TestInsightOpsClient_ClientPut(t *testing.T) {
	mockRequestPayload := &mockObject{Data: "some req data..."}
	mockResponse := &mockObject{Data: "some data..."}
	requestMatcher := NewRequestMatcher(http.MethodPut, "/api/testing", mockRequestPayload, http.StatusOK, mockResponse)

	c := getTestClient(requestMatcher)
	expectedResponse := &mockObject{}
	body, err := c.put("/api/testing", mockRequestPayload)
	err = json.Unmarshal(body, &expectedResponse)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse.Data, mockResponse.Data)
}

func TestInsightOpsClient_ClientPutResponseNotStatusCreated(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodPut, "/api/testing", &mockObject{}, http.StatusUnauthorized, &mockObject{})

	c := getTestClient(requestMatcher)
	_, err := c.put("/api/testing", &mockObject{})

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}

func TestInsightOpsClient_ClientDelete(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodDelete, "/api/testing", nil, http.StatusNoContent, nil)
	c := getTestClient(requestMatcher)
	err := c.delete("/api/testing")
	assert.Nil(t, err)
}

func TestInsightOpsClient_ClientGetResponseNotStatusNoContent(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodDelete, "/api/testing", nil, http.StatusUnauthorized, nil)
	c := getTestClient(requestMatcher)
	err := c.delete("/api/testing")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}
