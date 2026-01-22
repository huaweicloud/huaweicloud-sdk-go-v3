package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcProtectsVo **参数解释**： vpc protects vo **取值范围**： 不涉及
type VpcProtectsVo struct {

	// **参数解释**： 总防护VPC数 **取值范围**： 不涉及
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 防火墙东西向防护可进行跨账号防护VPC，self_total表示本项目防护VPC总数。 **取值范围**： 不涉及
	SelfTotal *int32 `json:"self_total,omitempty"`

	// **参数解释**： 防火墙东西向防护可进行跨账号防护VPC，other_total表示其他项目防护VPC数 **取值范围**： 不涉及
	OtherTotal *int32 `json:"other_total,omitempty"`

	// **参数解释**： 防火墙东西向防护可进行跨账号防护VPC，protect_vpcs表示总体防护VPC列表 **取值范围**： 不涉及
	ProtectVpcs *[]VpcAttachmentDetail `json:"protect_vpcs,omitempty"`

	// **参数解释**： 防火墙东西向防护可进行跨账号防护VPC，self_protect_vpcs表示本项目防护VPC列表 **取值范围**： 不涉及
	SelfProtectVpcs *[]VpcAttachmentDetail `json:"self_protect_vpcs,omitempty"`

	// **参数解释**： 防火墙东西向防护可进行跨账号防护VPC，other_protect_vpcs表示其他项目防护VPC列表 **取值范围**： 不涉及
	OtherProtectVpcs *[]VpcAttachmentDetail `json:"other_protect_vpcs,omitempty"`

	// **参数解释**： 租户的所有VPC资产数量 **取值范围**： 不涉及
	TotalAssets *int32 `json:"total_assets,omitempty"`
}

func (o VpcProtectsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcProtectsVo struct{}"
	}

	return strings.Join([]string{"VpcProtectsVo", string(data)}, " ")
}
