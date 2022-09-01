package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ValidateHttpsInfoV2Response struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	// 响应结果
	Result *string `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ValidateHttpsInfoV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateHttpsInfoV2Response struct{}"
	}

	return strings.Join([]string{"ValidateHttpsInfoV2Response", string(data)}, " ")
}
