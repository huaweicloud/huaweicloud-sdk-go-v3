package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncentiveAndDiscountPolicy struct {

	// 云服务类型列表。
	ServiceTypeCode *string `json:"service_type_code,omitempty" xml:"service_type_code"`

	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。
	ServiceTypeName *string `json:"service_type_name,omitempty" xml:"service_type_name"`

	// 激励策略。 0：非特定产品1：特定产品2：无业绩无返点13：有业绩无返点
	IncentivePolicy *string `json:"incentive_policy,omitempty" xml:"incentive_policy"`

	// 是否允许应用伙伴授予折扣。 YES：允许应用NO：不许应用
	AllowDiscount *string `json:"allow_discount,omitempty" xml:"allow_discount"`
}

func (o IncentiveAndDiscountPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncentiveAndDiscountPolicy struct{}"
	}

	return strings.Join([]string{"IncentiveAndDiscountPolicy", string(data)}, " ")
}
