package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckoutAndUpdateResponse Response Object
type BatchCheckoutAndUpdateResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]VersionModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchCheckoutAndUpdateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckoutAndUpdateResponse struct{}"
	}

	return strings.Join([]string{"BatchCheckoutAndUpdateResponse", string(data)}, " ")
}
