// Package insight_goclient provides a insight client which allows the interaction with insight rest API
// via the seamless resource interfaces exposed. Examples include:
// - Logsets
// - Logs
// - Tags
// - Labels
package insightops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const INSIGHTOPS_API = "https://%s.rest.logs.insight.rapid7.com"

type InsightOpsClient struct {
	InsightOpsUrl string
	ApiKey     string
	HttpClient *http.Client
}

// NewInsightOpsClient creates a InsightOps client which exposes an interface with CRUD operations for each of the
// resources provided by InsightOps rest API
func NewInsightOpsClient(apiKey, region string) (*InsightOpsClient, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("ApiKey is mandatory to initialize InsightOps client")
	}
	if region == "" {
		return nil, fmt.Errorf("Region is mandatory to initialize InsightOps client")
	}
	client := &http.Client{}
	return &InsightOpsClient{fmt.Sprintf(INSIGHTOPS_API, region), apiKey, client}, nil
}

func (client *InsightOpsClient) sendRequest(request *http.Request, expectedResponseCode int) ([]byte, error) {
	if request.Body != nil {
		requestBody, err := request.GetBody()
		if err != nil {
			return nil, err
		}
		requestBodyBuffer := new(bytes.Buffer)
		requestBodyBuffer.ReadFrom(requestBody)
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("x-api-key", client.ApiKey)
	response, err := client.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(body)
	if response.StatusCode != expectedResponseCode {
		return nil, fmt.Errorf("Received a non expected response status code %d, expected code was %d. Response: %s", response.StatusCode, expectedResponseCode, bodyString)
	}
	return body, nil
}

func (client *InsightOpsClient) get(path string, resource interface{}) error {
	url := client.getInsightOpsUrl(path)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	body, err := client.sendRequest(request, http.StatusOK)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, resource)
}

func (client *InsightOpsClient) post(path string, requestBody interface{}) ([]byte, error) {
	url := client.getInsightOpsUrl(path)
	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	return client.sendRequest(request, http.StatusCreated)
}

func (client *InsightOpsClient) put(path string, requestBody interface{}) ([]byte, error) {
	url := client.getInsightOpsUrl(path)
	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	return client.sendRequest(request, http.StatusOK)
}

func (client *InsightOpsClient) delete(path string) error {
	url := client.getInsightOpsUrl(path)
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	_, err = client.sendRequest(request, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

func (client *InsightOpsClient) getInsightOpsUrl(path string) string {
	return fmt.Sprintf("%s%s", client.InsightOpsUrl, path)
}
