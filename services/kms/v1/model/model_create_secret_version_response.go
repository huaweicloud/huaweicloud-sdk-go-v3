package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateSecretVersionResponse struct {
	VersionMetadata *VersionMetadata `json:"version_metadata,omitempty"`
	HttpStatusCode  int              `json:"-"`
}

func (o CreateSecretVersionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSecretVersionResponse struct{}"
	}

	return strings.Join([]string{"CreateSecretVersionResponse", string(data)}, " ")
}
