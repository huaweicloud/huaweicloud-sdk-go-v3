package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeInstanceVolumeRequestBody struct {
	Volume *ResizeInstanceVolumeOption `json:"volume"`

	// 扩容包年包月实例的存储容量时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。 - true，表示自动从账户中支付。 - false，表示手动从账户中支付，默认为该方式。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o ResizeInstanceVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeInstanceVolumeRequestBody", string(data)}, " ")
}
