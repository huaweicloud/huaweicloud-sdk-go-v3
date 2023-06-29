package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHooksResponse Response Object
type DeleteHooksResponse struct {
	Error *Error `json:"error,omitempty"`

	// 响应结果
	Result *interface{} `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteHooksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHooksResponse struct{}"
	}

	return strings.Join([]string{"DeleteHooksResponse", string(data)}, " ")
}
