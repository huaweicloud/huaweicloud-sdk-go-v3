package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 函数流执行请求body体
type StartSyncWorkflowExecutionRequestBody struct {

	// 函数执行时需要的Header
	Headers *interface{} `json:"headers,omitempty"`

	// 定义函数执行时的入参，支持使用JSONPATH进行映射，以及指定默认值
	Input *interface{} `json:"input"`
}

func (o StartSyncWorkflowExecutionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSyncWorkflowExecutionRequestBody struct{}"
	}

	return strings.Join([]string{"StartSyncWorkflowExecutionRequestBody", string(data)}, " ")
}
