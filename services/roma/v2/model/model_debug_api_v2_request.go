package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DebugApiV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// API的编号
	ApiId string `json:"api_id" xml:"api_id"`

	Body *ApiDebugInfo `json:"body,omitempty" xml:"body"`
}

func (o DebugApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugApiV2Request struct{}"
	}

	return strings.Join([]string{"DebugApiV2Request", string(data)}, " ")
}
