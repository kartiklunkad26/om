package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	StatusRunning   = "running"
	StatusSucceeded = "succeeded"
	StatusFailed    = "failed"
)

type InstallationsServiceOutput struct {
	ID         int
	Status     string
	Logs       string
	StartedAt  *time.Time `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at"`
	UserName   string     `json:"user_name"`
}

func (a Api) ListInstallations() ([]InstallationsServiceOutput, error) {
	req, err := http.NewRequest("GET", "/api/v0/installations", nil)
	if err != nil {
		return []InstallationsServiceOutput{}, err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return []InstallationsServiceOutput{}, fmt.Errorf("could not make api request to installations endpoint: %s", err)
	}

	defer resp.Body.Close()

	if err = validateStatusOK(resp); err != nil {
		return []InstallationsServiceOutput{}, err
	}

	var responseStruct struct {
		Installations []InstallationsServiceOutput
	}
	err = json.NewDecoder(resp.Body).Decode(&responseStruct)
	if err != nil {
		return []InstallationsServiceOutput{}, fmt.Errorf("failed to decode response: %s", err)
	}

	return responseStruct.Installations, nil
}

func (a Api) CreateInstallation(ignoreWarnings bool, deployProducts bool) (InstallationsServiceOutput, error) {
	deployProductsVal := "none"
	if deployProducts {
		deployProductsVal = "all"
	}

	data, err := json.Marshal(&struct {
		IgnoreWarnings string `json:"ignore_warnings"`
		DeployProducts string `json:"deploy_products"`
	}{
		IgnoreWarnings: fmt.Sprintf("%t", ignoreWarnings),
		DeployProducts: deployProductsVal,
	})
	if err != nil {
		return InstallationsServiceOutput{}, err
	}

	req, err := http.NewRequest("POST", "/api/v0/installations", bytes.NewReader(data))
	if err != nil {
		return InstallationsServiceOutput{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return InstallationsServiceOutput{}, fmt.Errorf("could not make api request to installations endpoint: %s", err)
	}

	defer resp.Body.Close()

	if err = validateStatusOK(resp); err != nil {
		return InstallationsServiceOutput{}, err
	}

	var installation struct {
		Install struct {
			ID int
		}
	}
	err = json.NewDecoder(resp.Body).Decode(&installation)
	if err != nil {
		return InstallationsServiceOutput{}, fmt.Errorf("failed to decode response: %s", err)
	}

	return InstallationsServiceOutput{ID: installation.Install.ID}, nil
}

// TODO: extract to helper package?
func (a Api) RunningInstallation() (InstallationsServiceOutput, error) {
	installationOutput, err := a.ListInstallations()
	if err != nil {
		return InstallationsServiceOutput{}, err
	}
	if len(installationOutput) > 0 && installationOutput[0].Status == StatusRunning {
		return installationOutput[0], nil
	}
	return InstallationsServiceOutput{}, nil
}

func (a Api) GetInstallation(id int) (InstallationsServiceOutput, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/v0/installations/%d", id), nil)
	if err != nil {
		return InstallationsServiceOutput{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return InstallationsServiceOutput{}, fmt.Errorf("could not make api request to installations status endpoint: %s", err)
	}

	defer resp.Body.Close()

	if err = validateStatusOK(resp); err != nil {
		return InstallationsServiceOutput{}, err
	}

	var output struct {
		Status string
	}
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return InstallationsServiceOutput{}, fmt.Errorf("failed to decode response: %s", err)
	}

	return InstallationsServiceOutput{Status: output.Status}, nil
}

func (a Api) GetInstallationLogs(id int) (InstallationsServiceOutput, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/v0/installations/%d/logs", id), nil)
	if err != nil {
		return InstallationsServiceOutput{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return InstallationsServiceOutput{}, fmt.Errorf("could not make api request to installations logs endpoint: %s", err)
	}

	defer resp.Body.Close()

	if err = validateStatusOK(resp); err != nil {
		return InstallationsServiceOutput{}, err
	}

	var output struct {
		Logs string
	}
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return InstallationsServiceOutput{}, fmt.Errorf("failed to decode response: %s", err)
	}

	return InstallationsServiceOutput{Logs: output.Logs}, nil
}
