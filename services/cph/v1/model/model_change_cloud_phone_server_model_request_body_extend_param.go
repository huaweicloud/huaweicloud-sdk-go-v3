package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCloudPhoneServerModelRequestBodyExtendParam 购买模式参数
type ChangeCloudPhoneServerModelRequestBodyExtendParam struct {

	// 否自动付款。默认不自动付款。如果是规格升配的变更，则需要支付额外的费用，如果是降配的变更，则会自动退款。 - 1 表示自动付款 - 0 表示不自动付款
	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`
}

func (o ChangeCloudPhoneServerModelRequestBodyExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCloudPhoneServerModelRequestBodyExtendParam struct{}"
	}

	return strings.Join([]string{"ChangeCloudPhoneServerModelRequestBodyExtendParam", string(data)}, " ")
}
