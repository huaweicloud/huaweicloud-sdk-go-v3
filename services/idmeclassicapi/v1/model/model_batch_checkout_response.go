package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckoutResponse Response Object
type BatchCheckoutResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]VersionModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchCheckoutResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckoutResponse struct{}"
	}

	return strings.Join([]string{"BatchCheckoutResponse", string(data)}, " ")
}
