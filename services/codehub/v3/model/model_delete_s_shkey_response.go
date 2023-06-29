package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSShkeyResponse Response Object
type DeleteSShkeyResponse struct {
	Error *Error `json:"error,omitempty"`

	// 响应结果
	Result *string `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSShkeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSShkeyResponse struct{}"
	}

	return strings.Join([]string{"DeleteSShkeyResponse", string(data)}, " ")
}
