package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowProjectRequest struct {
	ProjectId string `json:"project_id"`
}

func (o KeystoneShowProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowProjectRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowProjectRequest", string(data)}, " ")
}
