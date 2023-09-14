package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVersionResponse Response Object
type UpdateVersionResponse struct {
	VersionMetadata *VersionMetadata `json:"version_metadata,omitempty"`
	HttpStatusCode  int              `json:"-"`
}

func (o UpdateVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVersionResponse struct{}"
	}

	return strings.Join([]string{"UpdateVersionResponse", string(data)}, " ")
}
