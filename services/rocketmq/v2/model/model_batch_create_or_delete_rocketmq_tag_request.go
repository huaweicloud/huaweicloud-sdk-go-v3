package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateOrDeleteRocketmqTagRequest Request Object
type BatchCreateOrDeleteRocketmqTagRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *BatchCreateOrDeleteTagReq `json:"body,omitempty"`
}

func (o BatchCreateOrDeleteRocketmqTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteRocketmqTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteRocketmqTagRequest", string(data)}, " ")
}
