package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSecretVersionsResponse struct {
	// version_metadata对象。

	VersionMetadatas *[]VersionMetadata `json:"version_metadatas,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o ListSecretVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListSecretVersionsResponse", string(data)}, " ")
}
