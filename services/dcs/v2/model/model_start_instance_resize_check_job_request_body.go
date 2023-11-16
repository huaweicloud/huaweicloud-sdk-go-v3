package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartInstanceResizeCheckJobRequestBody 规格变更前置检查任务检查请求体
type StartInstanceResizeCheckJobRequestBody struct {

	// 新规格的容量，单位GB
	NewCapacity *int32 `json:"new_capacity,omitempty"`

	// 新的规格编码
	SpecCode *string `json:"spec_code,omitempty"`
}

func (o StartInstanceResizeCheckJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceResizeCheckJobRequestBody struct{}"
	}

	return strings.Join([]string{"StartInstanceResizeCheckJobRequestBody", string(data)}, " ")
}
