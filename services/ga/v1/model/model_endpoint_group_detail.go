package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointGroupDetail 终端节点组实例。
type EndpointGroupDetail struct {

	// 终端节点组ID。
	Id *string `json:"id,omitempty"`

	// 终端节点组名称，取值范围：1~64个字符之间，只能由数字、字母、中划线和中文组成。
	Name *string `json:"name,omitempty"`

	// 终端节点组描述信息，取值范围：0~255个字符之间，禁止输入字符：<>。
	Description *string `json:"description,omitempty"`

	Status *ConfigStatus `json:"status,omitempty"`

	// 流量拨分到此组的百分比。
	TrafficDialPercentage *int32 `json:"traffic_dial_percentage,omitempty"`

	// 终端节点组所属区域ID。
	RegionId *string `json:"region_id,omitempty"`

	// 关联监听器列表。
	Listeners *[]Id `json:"listeners,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 租户ID。
	DomainId *string `json:"domain_id,omitempty"`

	FrozenInfo *FrozenInfo `json:"frozen_info,omitempty"`
}

func (o EndpointGroupDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointGroupDetail struct{}"
	}

	return strings.Join([]string{"EndpointGroupDetail", string(data)}, " ")
}
