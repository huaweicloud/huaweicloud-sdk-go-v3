package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReqAllocateDehExtendParam 专属主机扩展属性。
type ReqAllocateDehExtendParam struct {

	// 专属主机计费模式，不传该参数时默认为postPaid。 - prePaid：包周期 - postPaid：按需
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 订购周期类型。 - month：包月 - year：包年
	PeriodType *string `json:"period_type,omitempty"`

	// 订购周期数，大于0的整数，当charging_mode为prePaid时生效，且该字段必选。
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 是否自动支付。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	// 是否自动续费。
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ReqAllocateDehExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqAllocateDehExtendParam struct{}"
	}

	return strings.Join([]string{"ReqAllocateDehExtendParam", string(data)}, " ")
}
