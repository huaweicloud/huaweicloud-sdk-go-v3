package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteHooksResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	// 响应结果
	Result *interface{} `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteHooksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHooksResponse struct{}"
	}

	return strings.Join([]string{"DeleteHooksResponse", string(data)}, " ")
}
