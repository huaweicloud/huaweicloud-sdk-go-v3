package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcProtectsVo vpc protects vo
type VpcProtectsVo struct {

	// 总防护VPC数
	Total *int32 `json:"total,omitempty"`

	// 防火墙东西向防护可进行跨账号防护VPC，self_total表示本项目防护VPC总数。
	SelfTotal *int32 `json:"self_total,omitempty"`

	// 防火墙东西向防护可进行跨账号防护VPC，other_total表示其他项目防护VPC数
	OtherTotal *int32 `json:"other_total,omitempty"`

	// 防火墙东西向防护可进行跨账号防护VPC，protect_vpcs指的是总体防护VPC列表
	ProtectVpcs *[]VpcAttachmentDetail `json:"protect_vpcs,omitempty"`

	// 防火墙东西向防护可进行跨账号防护VPC，self_protect_vpcs指的是本项目防护VPC列表
	SelfProtectVpcs *[]VpcAttachmentDetail `json:"self_protect_vpcs,omitempty"`

	// 防火墙东西向防护可进行跨账号防护VPC，other_protect_vpcs指的是其他项目防护VPC列表
	OtherProtectVpcs *[]VpcAttachmentDetail `json:"other_protect_vpcs,omitempty"`

	// 租户的所有VPC资产数量
	TotalAssets *int32 `json:"total_assets,omitempty"`
}

func (o VpcProtectsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcProtectsVo struct{}"
	}

	return strings.Join([]string{"VpcProtectsVo", string(data)}, " ")
}
