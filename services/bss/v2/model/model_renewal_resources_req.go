package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RenewalResourcesReq struct {

	// 资源ID列表。 只支持传入主资源ID，最多10个资源ID。 哪些资源是主资源请根据“查询客户包年/包月资源列表”接口响应参数中的“is_main_resource”来标识。
	ResourceIds []string `json:"resource_ids"`

	// 周期类型： 2：月3：年
	PeriodType int32 `json:"period_type"`

	// 周期数目： 如果是月，目前支持1-11如果是年，目前支持1-3
	PeriodNum int32 `json:"period_num"`

	// 到期策略（字段已废弃，请勿使用该字段。此字段必填，需携带，但携带的枚举实际并不生效）： 0：进入宽限期/保留期1：转按需2：自动退订3：自动续订
	ExpirePolicy *int32 `json:"expire_policy,omitempty"`

	// 是否自动支付。 0：否1：是 此参数不携带或携带值为null时，默认值为“0：否”，即不自动支付。
	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`
}

func (o RenewalResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenewalResourcesReq struct{}"
	}

	return strings.Join([]string{"RenewalResourcesReq", string(data)}, " ")
}
