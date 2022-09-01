package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 趋势图数据模型
type FrontLine struct {

	// 数据点集合
	PointList *[]FrontPoint `json:"point_list,omitempty" xml:"point_list"`

	// 标题
	Title *string `json:"title,omitempty" xml:"title"`

	// 单位
	Unit *string `json:"unit,omitempty" xml:"unit"`

	// 百分比
	Precision *int32 `json:"precision,omitempty" xml:"precision"`

	// 日期类型
	DataType *string `json:"data_type,omitempty" xml:"data_type"`

	// 是否可见
	Visible *bool `json:"visible,omitempty" xml:"visible"`
}

func (o FrontLine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrontLine struct{}"
	}

	return strings.Join([]string{"FrontLine", string(data)}, " ")
}
