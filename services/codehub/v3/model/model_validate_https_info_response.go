package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ValidateHttpsInfoResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	// 响应结果
	Result *string `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ValidateHttpsInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateHttpsInfoResponse struct{}"
	}

	return strings.Join([]string{"ValidateHttpsInfoResponse", string(data)}, " ")
}
