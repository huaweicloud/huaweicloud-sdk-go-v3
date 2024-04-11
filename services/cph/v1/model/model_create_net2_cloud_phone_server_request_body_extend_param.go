package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNet2CloudPhoneServerRequestBodyExtendParam 计费模式参数。
type CreateNet2CloudPhoneServerRequestBodyExtendParam struct {

	// 计费类型。  [- 0：包周期](tag:hws,hws_hk,cmcc) [- 1：按需](tag:fcs)
	ChargingMode int32 `json:"charging_mode"`

	// 订购周期类型。 - 2：月 - 3：年
	PeriodType int32 `json:"period_type"`

	// 订购周期数。 当订购周期为月时，取值范围[1, 9]。 当订购周期为年时，取值范围[1,10]。
	PeriodNum int32 `json:"period_num"`

	// 是否自动付款。默认不自动付款。 - 1：自动付款 - 0：不自动付款
	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`

	// 是否自动续订。默认不自动续订。 - 1：自动续订 - 0：不自动续订
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`

	// 企业项目ID。 该字段不传（或传为字符串“0”），则将资源绑定给默认企业项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateNet2CloudPhoneServerRequestBodyExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNet2CloudPhoneServerRequestBodyExtendParam struct{}"
	}

	return strings.Join([]string{"CreateNet2CloudPhoneServerRequestBodyExtendParam", string(data)}, " ")
}
