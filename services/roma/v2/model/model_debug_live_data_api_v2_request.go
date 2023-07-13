package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DebugLiveDataApiV2Request Request Object
type DebugLiveDataApiV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 后端API的编号
	LdApiId string `json:"ld_api_id"`

	Body *LdApiTest `json:"body,omitempty"`
}

func (o DebugLiveDataApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugLiveDataApiV2Request struct{}"
	}

	return strings.Join([]string{"DebugLiveDataApiV2Request", string(data)}, " ")
}
