package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEipBandwidthsRequest Request Object
type ListEipBandwidthsRequest struct {

	// - 功能说明：每页返回的个数 - 取值范围：取值范围：1~[2000]，其中2000为局点差异项，具体取值由局点决定
	Limit *string `json:"limit,omitempty"`

	// - 功能说明：分页查询起始的资源ID，为空时为查询第一页
	Marker *string `json:"marker,omitempty"`

	// display in the form \"fields=id&fields=name&...\"  Supported fields：id/name/tenant_id/size/ratio_95peak_plus/ingress_size/bandwidth_type/admin_state/billing_info/charge_mode/type/publicip_info/enable_bandwidth_rules/rule_quota/bandwidth_rules/public_border_group/created_at/updated_at/lock_infos
	Fields *[]string `json:"fields,omitempty"`

	// - 功能说明：带宽唯一标识
	Id *string `json:"id,omitempty"`

	// - 功能说明：带宽类型，共享带宽默认为share。 - 取值范围：share，bgp，telcom，sbgp等。   - share：共享带宽   - bgp：动态bgp   - telcom ：联通   - sbgp：静态bgp
	BandwidthType *string `json:"bandwidth_type,omitempty"`

	// - 功能说明：宽带名称，按照宽带名称过滤
	Name *string `json:"name,omitempty"`

	// - 功能说明：根据宽带名称模糊查询过滤
	NameLike *string `json:"name_like,omitempty"`

	// - 功能说明：根据tenant_id过滤
	TenantId *string `json:"tenant_id,omitempty"`

	// - 功能说明：根据入云大小过滤
	IngressSize *string `json:"ingress_size,omitempty"`

	// - 功能说明：根据宽带状态过滤
	AdminState *string `json:"admin_state,omitempty"`

	// - 功能说明：根据计费信息过滤
	BillingInfo *string `json:"billing_info,omitempty"`

	// - 功能说明：根据标签过滤
	Tags *string `json:"tags,omitempty"`

	// - 功能说明：根据是否带宽分组使能过滤 - 取值范围：true、false
	EnableBandwidthRules *string `json:"enable_bandwidth_rules,omitempty"`

	// - 功能说明：根据规则数值过滤
	RuleQuota *int32 `json:"rule_quota,omitempty"`

	// - 功能说明：根据站点信息过滤
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// - 功能说明：按流量计费,按带宽计费还是按增强型95计费 - 取值范围：bandwidth（按带宽计费），traffic（按流量计费），95peak_plus（按增强型95计费），不返回或者为空时表示是bandwidth - 约束：只有共享带宽支持95peak_plus（按增强型95计费），按增强型95计费时需要指定保底百分比，默认是20%
	ChargeMode *string `json:"charge_mode,omitempty"`

	// - 功能说明：带宽大小。共享带宽的大小有最小值限制，默认为5M，可能因局点不同而不同。 - 取值范围：默认5Mbit/s~2000Mbit/s（具体范围以各区域配置为准，请参见控制台对应页面显示）。   调整带宽时的最小单位会根据带宽范围不同存在差异。 - 小于等于300Mbit/s：默认最小单位为1Mbit/s。 - 300Mbit/s~1000Mbit/s：默认最小单位为50Mbit/s。 - 大于1000Mbit/s：默认最小单位为500Mbit/s。
	Size *string `json:"size,omitempty"`

	// - 功能说明：带宽类型，标识是否是共享带宽 - 取值范围：WHOLE，PER。   - WHOLE表示共享带宽   - PER表示独享带宽
	Type *string `json:"type,omitempty"`
}

func (o ListEipBandwidthsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEipBandwidthsRequest struct{}"
	}

	return strings.Join([]string{"ListEipBandwidthsRequest", string(data)}, " ")
}
