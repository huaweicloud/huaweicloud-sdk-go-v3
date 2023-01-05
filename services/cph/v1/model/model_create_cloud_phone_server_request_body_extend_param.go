package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 计费模式参数
type CreateCloudPhoneServerRequestBodyExtendParam struct {

	// 计费类型 0 表示包周期
	ChargingMode int32 `json:"charging_mode"`

	// 订购周期类型 - 2 表示月 - 3 表示年
	PeriodType int32 `json:"period_type"`

	// 订购周期数 当订购周期为月时，取值范围[1, 9]。 当订购周期为年时，取值范围[1,10]
	PeriodNum int32 `json:"period_num"`

	// 是否自动付款。默认不自动付款。 - 1 表示自动付款 - 0 表示不自动付款
	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`

	// 是否自动续订。默认不自动续订。 - 1 表示自动续订 - 0 表示不自动续订
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`

	// 企业项目ID。 该字段不传（或传为字符串“0”），则将资源绑定给默认企业项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateCloudPhoneServerRequestBodyExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudPhoneServerRequestBodyExtendParam struct{}"
	}

	return strings.Join([]string{"CreateCloudPhoneServerRequestBodyExtendParam", string(data)}, " ")
}
