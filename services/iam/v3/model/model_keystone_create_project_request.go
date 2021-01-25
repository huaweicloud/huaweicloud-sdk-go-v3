package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneCreateProjectRequest struct {
	Body *KeystoneCreateProjectRequestBody `json:"body,omitempty"`
}

func (o KeystoneCreateProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateProjectRequest struct{}"
	}

	return strings.Join([]string{"KeystoneCreateProjectRequest", string(data)}, " ")
}
