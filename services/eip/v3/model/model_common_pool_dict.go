package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 公共池详情
type CommonPoolDict struct {

	// 公共池名字
	Name *string `json:"name,omitempty" xml:"name"`

	// 状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 公共池类型，如bgp，sbgp等
	Type *string `json:"type,omitempty" xml:"type"`

	// 已经使用的ip数量
	Used *int32 `json:"used,omitempty" xml:"used"`

	// 功能说明：表示中心站点资源或者边缘站点资源 取值范围： center、边缘站点名称 约束：publicip只能绑定该字段相同的资源
	PublicBorderGroup *string `json:"public_border_group,omitempty" xml:"public_border_group"`

	// 默认不展示，取值, 公共池ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 功能说明：表示此publicip可以加入的共享带宽类型列表，如果为空列表，则表示该           publicip不能加入任何共享带宽 约束：publicip只能加入到有该带宽类型的共享带宽中
	AllowShareBandwidthTypes *[]string `json:"allow_share_bandwidth_types,omitempty" xml:"allow_share_bandwidth_types"`
}

func (o CommonPoolDict) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonPoolDict struct{}"
	}

	return strings.Join([]string{"CommonPoolDict", string(data)}, " ")
}
