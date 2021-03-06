package insightops

import (
	"encoding/json"
	"fmt"
)

const (
	LOGS_PATH = "/management/logs"
)

// The Logs resource allows you to interact with Logs in your account. The following operations are supported:
// - Get details of an existing Log
// - Get details of a list of all Logs
// - Create a new Log
// - Update an existing Log
// - Delete a Log

// Log represents the entity used to get an existing log from the InsightOps API
type Log struct {
	Id              string       `json:"id,omitempty"`
	Name            string       `json:"name"`
	LogsetsInfo     []*Info      `json:"logsets_info,omitempty"`
	UserData        *LogUserData `json:"user_data"`
	Tokens          []string     `json:"tokens,omitempty"`
	SourceType      string       `json:"source_type,omitempty"`
	TokenSeed       string       `json:"token_seed,omitempty"`
	Structures      []string     `json:"structures,omitempty"`
	RetentionPeriod string       `json:"retention_period,omitempty"`
	Links           []*Link      `json:"links,omitempty"`
}

// LogUserData represents user metadata
type LogUserData struct {
	AgentFileName string     `json:"le_agent_filename"`
	AgentFollow   StringBool `json:"le_agent_follow"`
}

type Logs struct {
	Logs []*Log `json:"logs"`
}

type LogRequest struct {
	Log *Log `json:"log"`
}

// GetLogs lists all Logs for an account
func (client *InsightOpsClient) GetLogs() ([]*Log, error) {
	var logs Logs
	if err := client.get(LOGS_PATH, &logs); err != nil {
		return nil, err
	}
	return logs.Logs, nil
}

// GetLog gets a specific Log from an account
func (client *InsightOpsClient) GetLog(logId string) (*Log, error) {
	var logRequest LogRequest
	endpoint, err := client.getLogEndpoint(logId)
	if err != nil {
		return nil, err
	}
	if err := client.get(endpoint, &logRequest); err != nil {
		return nil, err
	}
	return logRequest.Log, nil
}

func (client *InsightOpsClient) GetLogToken(logsetName, logName string) (string, string, error) {
	logset, err := client.GetLogsetByName(logsetName)
	if err != nil {
		return "", "", err
	}

	for _, logInfo := range logset.LogsInfo {
		if logInfo.Name == logName {
			log, err := client.GetLog(logInfo.Id)
			if err != nil {
				return "", "", err
			}
			if len(log.Tokens) > 0 {
				return log.Tokens[0], logInfo.Id, nil
			} else {
				return "", "", fmt.Errorf("No tokens for log %s found", logInfo.Name)
			}
		}
	}
	return "", "", fmt.Errorf("No tokens found for logset %s", logsetName)
}

// PostTag creates a new Log
func (client *InsightOpsClient) PostLog(log *Log) error {
	logRequest := LogRequest{log}
	resp, err := client.post(LOGS_PATH, logRequest)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &logRequest)
	if err != nil {
		return err
	}
	return nil
}

// PutTag updates an existing Log
func (client *InsightOpsClient) PutLog(log *Log) error {
	logRequest := LogRequest{log}
	endpoint, err := client.getLogEndpoint(log.Id)
	if err != nil {
		return err
	}
	resp, err := client.put(endpoint, logRequest)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &logRequest)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTag deletes a specific Log from an account.
func (client *InsightOpsClient) DeleteLog(logId string) error {
	endpoint, err := client.getLogEndpoint(logId)
	if err != nil {
		return err
	}
	return client.delete(endpoint)
}

func (client *InsightOpsClient) getLogEndpoint(logId string) (string, error) {
	if logId == "" {
		return "", fmt.Errorf("logId input parameter is mandatory")
	} else {
		return fmt.Sprintf("%s/%s", LOGS_PATH, logId), nil
	}
}
