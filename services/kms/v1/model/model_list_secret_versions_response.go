package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSecretVersionsResponse struct {
	// version_metadata对象。

	VersionMetadatas *[]VersionMetadata `json:"version_metadatas,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o ListSecretVersionsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSecretVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListSecretVersionsResponse", string(data)}, " ")
}
