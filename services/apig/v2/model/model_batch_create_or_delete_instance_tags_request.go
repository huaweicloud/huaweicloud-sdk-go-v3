package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateOrDeleteInstanceTagsRequest Request Object
type BatchCreateOrDeleteInstanceTagsRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *TmsUpdatePublicReq `json:"body,omitempty"`
}

func (o BatchCreateOrDeleteInstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteInstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteInstanceTagsRequest", string(data)}, " ")
}
