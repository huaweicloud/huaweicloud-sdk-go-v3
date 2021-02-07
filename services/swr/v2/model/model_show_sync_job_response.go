package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowSyncJobResponse struct {
	Body           *[]SyncJob `json:"body,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowSyncJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSyncJobResponse struct{}"
	}

	return strings.Join([]string{"ShowSyncJobResponse", string(data)}, " ")
}
