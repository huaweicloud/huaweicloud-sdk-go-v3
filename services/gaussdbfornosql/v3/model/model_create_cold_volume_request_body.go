package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateColdVolumeRequestBody struct {

	// 创建的冷数据存储大小，最小申请规格为500GB，最大申请规格为100000GB。单位：GB
	Size int32 `json:"size"`

	// 创建包年包月实例的冷数据存储时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。 - true，表示自动从账户中支付。 - false，表示手动从账户中支付，默认为该方式。
	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o CreateColdVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateColdVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"CreateColdVolumeRequestBody", string(data)}, " ")
}
