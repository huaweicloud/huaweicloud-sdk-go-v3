package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RegisterAgentResponse struct {
	// 状态信息

	Status *string `json:"status,omitempty"`
	// 返回结果

	Result *interface{} `json:"result,omitempty"`
	// 返回错误

	Error          *interface{} `json:"error,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o RegisterAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterAgentResponse struct{}"
	}

	return strings.Join([]string{"RegisterAgentResponse", string(data)}, " ")
}
