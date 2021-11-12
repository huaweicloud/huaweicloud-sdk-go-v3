package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchSetSubCustomerDiscountResponse struct {
	// 部分成功部分失败的时候返回的失败记录，如果全成功，该记录为空，具体参见表2。

	ErrorDetails   *[]ErrorDetail `json:"error_details,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o BatchSetSubCustomerDiscountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetSubCustomerDiscountResponse struct{}"
	}

	return strings.Join([]string{"BatchSetSubCustomerDiscountResponse", string(data)}, " ")
}
