package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RestoreInstanceFromCollectionRequest struct {
	XLanguage  *string                                   `json:"X-Language,omitempty"`
	InstanceId string                                    `json:"instance_id"`
	Body       *RestoreInstanceFromCollectionRequestBody `json:"body,omitempty"`
}

func (o RestoreInstanceFromCollectionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreInstanceFromCollectionRequest struct{}"
	}

	return strings.Join([]string{"RestoreInstanceFromCollectionRequest", string(data)}, " ")
}
