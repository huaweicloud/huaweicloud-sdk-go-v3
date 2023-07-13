package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateAssetRequest Request Object
type CreateOrUpdateAssetRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *EntityWithExtInfo `json:"body,omitempty"`
}

func (o CreateOrUpdateAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateAssetRequest struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateAssetRequest", string(data)}, " ")
}
