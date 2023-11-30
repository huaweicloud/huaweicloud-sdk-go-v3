package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DimensionHierarchiesVo struct {

	// 编码
	Id *int64 `json:"id,omitempty"`

	// 层级名称
	Name *string `json:"name,omitempty"`

	Attrs *[]HierarchiesAttrVo `json:"attrs,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`
}

func (o DimensionHierarchiesVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionHierarchiesVo struct{}"
	}

	return strings.Join([]string{"DimensionHierarchiesVo", string(data)}, " ")
}
