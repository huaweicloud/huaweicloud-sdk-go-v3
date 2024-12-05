package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRedisDisabledCommandsRequest Request Object
type DeleteRedisDisabledCommandsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *DeleteRedisDisabledCommandsRequestBody `json:"body,omitempty"`
}

func (o DeleteRedisDisabledCommandsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRedisDisabledCommandsRequest struct{}"
	}

	return strings.Join([]string{"DeleteRedisDisabledCommandsRequest", string(data)}, " ")
}
