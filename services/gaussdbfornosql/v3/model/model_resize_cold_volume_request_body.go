package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeColdVolumeRequestBody struct {

	// 待扩容后冷存储空间大小,单位：GB。用户每次至少选择1GB扩容量，且必须为整数。待扩容后的最大存储空间为100000GB。
	Size int32 `json:"size"`

	// 扩容包年包月实例的冷数据存储容量时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。 ·true，表示自动从账户中支付。 ·false，表示手动从账户中支付，默认为该方式。
	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o ResizeColdVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeColdVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeColdVolumeRequestBody", string(data)}, " ")
}
