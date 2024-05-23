package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRedisBigKeysRequest Request Object
type ShowRedisBigKeysRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ShowRedisBigKeysRequestBody `json:"body,omitempty"`
}

func (o ShowRedisBigKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRedisBigKeysRequest struct{}"
	}

	return strings.Join([]string{"ShowRedisBigKeysRequest", string(data)}, " ")
}
