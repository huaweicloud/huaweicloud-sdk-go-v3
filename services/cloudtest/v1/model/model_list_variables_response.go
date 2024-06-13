package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVariablesResponse Response Object
type ListVariablesResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	Error *CommonResponseErrorOfApiTest `json:"error,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`

	Result *PageResults `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListVariablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVariablesResponse struct{}"
	}

	return strings.Join([]string{"ListVariablesResponse", string(data)}, " ")
}
