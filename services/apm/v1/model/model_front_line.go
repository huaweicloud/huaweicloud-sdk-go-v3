package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FrontLine 趋势图数据模型。
type FrontLine struct {

	// 数据点集合。
	PointList *[]FrontPoint `json:"point_list,omitempty"`

	// 标题。
	Title *string `json:"title,omitempty"`

	// 单位。
	Unit *string `json:"unit,omitempty"`

	// 百分比。
	Precision *int32 `json:"precision,omitempty"`

	// 日期类型。
	DataType *string `json:"data_type,omitempty"`

	// 是否可见。
	Visible *bool `json:"visible,omitempty"`
}

func (o FrontLine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrontLine struct{}"
	}

	return strings.Join([]string{"FrontLine", string(data)}, " ")
}
