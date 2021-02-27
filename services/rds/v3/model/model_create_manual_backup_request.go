package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateManualBackupRequest struct {
	XLanguage *string                        `json:"X-Language,omitempty"`
	Body      *CreateManualBackupRequestBody `json:"body,omitempty"`
}

func (o CreateManualBackupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateManualBackupRequest struct{}"
	}

	return strings.Join([]string{"CreateManualBackupRequest", string(data)}, " ")
}
