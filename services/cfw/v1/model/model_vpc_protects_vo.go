package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcProtectsVo vpc protects vo
type VpcProtectsVo struct {

	// 总VPC数
	Total *int32 `json:"total,omitempty"`

	// 本项目防护VPC数
	SelfTotal *int32 `json:"self_total,omitempty"`

	// 其他项目防护VPC数
	OtherTotal *int32 `json:"other_total,omitempty"`

	// 防护VPC
	ProtectVpcs *[]VpcAttachmentDetail `json:"protect_vpcs,omitempty"`

	// 本项目防护VPC
	SelfProtectVpcs *[]VpcAttachmentDetail `json:"self_protect_vpcs,omitempty"`

	// 其他项目防护VPC
	OtherProtectVpcs *[]VpcAttachmentDetail `json:"other_protect_vpcs,omitempty"`

	// 所有资产数量
	TotalAssets *int32 `json:"total_assets,omitempty"`
}

func (o VpcProtectsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcProtectsVo struct{}"
	}

	return strings.Join([]string{"VpcProtectsVo", string(data)}, " ")
}
