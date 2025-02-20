package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountMultiResourcesRequest Request Object
type CountMultiResourcesRequest struct {

	// 厂商来源（默认RMS，可填RMS/ALI/AWS）
	Vendor string `json:"vendor"`

	// 视图 id，视图模式下必填
	ViewId *string `json:"view_id,omitempty"`

	// 是否为资源模块
	IsResource *bool `json:"is_resource,omitempty"`

	// 区域 id
	RegionId *string `json:"region_id,omitempty"`
}

func (o CountMultiResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountMultiResourcesRequest struct{}"
	}

	return strings.Join([]string{"CountMultiResourcesRequest", string(data)}, " ")
}
