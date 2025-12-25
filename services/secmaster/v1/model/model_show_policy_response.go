package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyResponse Response Object
type ShowPolicyResponse struct {

	// 返回编码
	Code *string `json:"code,omitempty"`

	// 返回信息
	Message *string `json:"message,omitempty"`

	// 返回数据
	Data *interface{} `json:"data,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 成功状态
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicyResponse", string(data)}, " ")
}
