package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTagsRequest Request Object
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
