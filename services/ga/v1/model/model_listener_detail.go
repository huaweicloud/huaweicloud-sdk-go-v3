package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListenerDetail 监听器实例。
type ListenerDetail struct {

	// 监听器ID。
	Id *string `json:"id,omitempty"`

	// 监听器名称，取值范围：1~64个字符之间，只能由数字、字母、中划线和中文组成。
	Name *string `json:"name,omitempty"`

	// 监听器的描述信息，取值范围：0~255个字符之间，禁止输入字符：<>。
	Description *string `json:"description,omitempty"`

	Protocol *ListenerProtocol `json:"protocol,omitempty"`

	Status *ConfigStatus `json:"status,omitempty"`

	// 监听端口范围列表。
	PortRanges *[]PortRange `json:"port_ranges,omitempty"`

	ClientAffinity *ClientAffinity `json:"client_affinity,omitempty"`

	// 全球加速实例ID。
	AcceleratorId *string `json:"accelerator_id,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 租户ID。
	DomainId *string `json:"domain_id,omitempty"`

	FrozenInfo *FrozenInfo `json:"frozen_info,omitempty"`

	// 标签列表。
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o ListenerDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerDetail struct{}"
	}

	return strings.Join([]string{"ListenerDetail", string(data)}, " ")
}
