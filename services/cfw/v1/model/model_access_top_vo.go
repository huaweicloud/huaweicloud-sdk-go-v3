package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessTopVo struct {

	// **参数解释**： 阻断次数 **取值范围**： 不涉及
	DenyCount *int64 `json:"deny_count,omitempty"`

	// **参数解释**： 高频命中的阻断策略ID **取值范围**： 不涉及
	DenyTopOneAclId *string `json:"deny_top_one_acl_id,omitempty"`

	// **参数解释**： 高频命中的阻断策略ID **取值范围**： 不涉及
	DenyTopOneAclName *string `json:"deny_top_one_acl_name,omitempty"`

	// **参数解释**： 命中次数 **取值范围**： 不涉及
	HitCount *int64 `json:"hit_count,omitempty"`

	// **参数解释**： TOP出云阻断目的IP列表 **取值范围**： 不涉及
	In2outDenyDstIpList *[]AccessTopMemberVo `json:"in2out_deny_dst_ip_list,omitempty"`

	// **参数解释**： TOP出云阻断端口列表 **取值范围**： 不涉及
	In2outDenyDstPortList *[]AccessTopMemberVo `json:"in2out_deny_dst_port_list,omitempty"`

	// **参数解释**： TOP出云阻断目的地区列表 **取值范围**： 不涉及
	In2outDenyDstRegionList *[]AccessTopMemberVo `json:"in2out_deny_dst_region_list,omitempty"`

	// **参数解释**： TOP出云阻断源IP列表 **取值范围**： 不涉及
	In2outDenySrcIpList *[]AccessTopMemberVo `json:"in2out_deny_src_ip_list,omitempty"`

	// **参数解释**： TOP入云阻断目的IP列表 **取值范围**： 不涉及
	Out2inDenyDstIpList *[]AccessTopMemberVo `json:"out2in_deny_dst_ip_list,omitempty"`

	// **参数解释**： TOP入云阻断目的端口列表 **取值范围**： 不涉及
	Out2inDenyDstPortList *[]AccessTopMemberVo `json:"out2in_deny_dst_port_list,omitempty"`

	// **参数解释**： TOP入云阻断源IP列表 **取值范围**： 不涉及
	Out2inDenySrcIpList *[]AccessTopMemberVo `json:"out2in_deny_src_ip_list,omitempty"`

	// **参数解释**： TOP入云阻断源端口列表 **取值范围**： 不涉及
	Out2inDenySrcPortList *[]AccessTopMemberVo `json:"out2in_deny_src_port_list,omitempty"`

	// **参数解释**： TOP入云阻断源地区列表 **取值范围**： 不涉及
	Out2inDenySrcRegionList *[]AccessTopMemberVo `json:"out2in_deny_src_region_list,omitempty"`

	// **参数解释**： 放行次数 **取值范围**： 不涉及
	PermitCount *int64 `json:"permit_count,omitempty"`

	// **参数解释**： 高频命中的放行策略ID **取值范围**： 不涉及
	PermitTopOneAclId *string `json:"permit_top_one_acl_id,omitempty"`

	// **参数解释**： 高频命中的放行策略名称 **取值范围**： 不涉及
	PermitTopOneAclName *string `json:"permit_top_one_acl_name,omitempty"`

	// **参数解释**： 命中趋势 **取值范围**： 不涉及
	Records *[]AccessTopStatisticsVo `json:"records,omitempty"`

	// **参数解释**： TOP阻断规则列表 **取值范围**： 不涉及
	TopDenyRuleList *[]AccessTopMemberVo `json:"top_deny_rule_list,omitempty"`
}

func (o AccessTopVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessTopVo struct{}"
	}

	return strings.Join([]string{"AccessTopVo", string(data)}, " ")
}
