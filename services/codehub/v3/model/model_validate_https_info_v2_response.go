package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateHttpsInfoV2Response Response Object
type ValidateHttpsInfoV2Response struct {
	Error *Error `json:"error,omitempty"`

	// 响应结果
	Result *string `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ValidateHttpsInfoV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateHttpsInfoV2Response struct{}"
	}

	return strings.Join([]string{"ValidateHttpsInfoV2Response", string(data)}, " ")
}
