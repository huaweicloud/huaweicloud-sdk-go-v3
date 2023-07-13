package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecretVersionsResponse Response Object
type ListSecretVersionsResponse struct {

	// version_metadata对象。
	VersionMetadatas *[]VersionMetadata `json:"version_metadatas,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSecretVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListSecretVersionsResponse", string(data)}, " ")
}
