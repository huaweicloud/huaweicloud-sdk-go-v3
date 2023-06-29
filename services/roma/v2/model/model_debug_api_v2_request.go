package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DebugApiV2Request Request Object
type DebugApiV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// API的编号
	ApiId string `json:"api_id"`

	Body *ApiDebugInfo `json:"body,omitempty"`
}

func (o DebugApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugApiV2Request struct{}"
	}

	return strings.Join([]string{"DebugApiV2Request", string(data)}, " ")
}
