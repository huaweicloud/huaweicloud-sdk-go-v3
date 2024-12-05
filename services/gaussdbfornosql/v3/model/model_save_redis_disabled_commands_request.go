package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveRedisDisabledCommandsRequest Request Object
type SaveRedisDisabledCommandsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *SaveRedisDisabledCommandsRequestBody `json:"body,omitempty"`
}

func (o SaveRedisDisabledCommandsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveRedisDisabledCommandsRequest struct{}"
	}

	return strings.Join([]string{"SaveRedisDisabledCommandsRequest", string(data)}, " ")
}
