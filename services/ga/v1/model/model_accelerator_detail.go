package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceleratorDetail 全球加速器实例。
type AcceleratorDetail struct {

	// 全球加速器ID。
	Id *string `json:"id,omitempty"`

	// 全球加速器名称，取值范围：1~64个字符之间，只能由数字、字母、中划线和中文组成。
	Name *string `json:"name,omitempty"`

	// 全球加速器描述信息，取值范围：0~255个字符之间，禁止输入字符：<>。
	Description *string `json:"description,omitempty"`

	Status *ConfigStatus `json:"status,omitempty"`

	// 全球加速器IP列表。
	IpSets *[]AccelerateIp `json:"ip_sets,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 租户ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 租户的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 规格ID。
	FlavorId *string `json:"flavor_id,omitempty"`

	FrozenInfo *FrozenInfo `json:"frozen_info,omitempty"`

	// 标签列表。
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o AcceleratorDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceleratorDetail struct{}"
	}

	return strings.Join([]string{"AcceleratorDetail", string(data)}, " ")
}
