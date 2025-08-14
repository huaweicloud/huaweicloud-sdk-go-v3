package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectConfigRequest Request Object
type ShowProjectConfigRequest struct {

	// 云存储ID。
	CloudStorageId string `json:"cloud_storage_id"`
}

func (o ShowProjectConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectConfigRequest", string(data)}, " ")
}
