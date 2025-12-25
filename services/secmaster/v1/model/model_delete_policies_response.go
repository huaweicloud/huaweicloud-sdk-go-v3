package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePoliciesResponse Response Object
type DeletePoliciesResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 返回数据
	Data *string `json:"data,omitempty"`

	// 错误信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePoliciesResponse struct{}"
	}

	return strings.Join([]string{"DeletePoliciesResponse", string(data)}, " ")
}
