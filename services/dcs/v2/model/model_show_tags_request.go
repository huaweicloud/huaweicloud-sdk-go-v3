package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTagsRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowTagsRequest", string(data)}, " ")
}
