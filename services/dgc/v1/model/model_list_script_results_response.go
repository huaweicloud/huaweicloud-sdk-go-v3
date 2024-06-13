package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScriptResultsResponse Response Object
type ListScriptResultsResponse struct {

	// 执行状态。 - LAUNCHING ：提交中 - RUNNING ： 运行中 - FINISHED：执行成功 - FAILED：执行失败
	Status *string `json:"status,omitempty"`

	// 执行结果
	Results *[]Result `json:"results,omitempty"`

	// 执行失败消息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListScriptResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptResultsResponse struct{}"
	}

	return strings.Join([]string{"ListScriptResultsResponse", string(data)}, " ")
}
