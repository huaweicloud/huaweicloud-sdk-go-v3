package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIndicatorDetailIndicatorType 指标类型统计数据
type CreateIndicatorDetailIndicatorType struct {

	// 指标类型
	IndicatorType string `json:"indicator_type"`

	// 情报类型ID
	Id string `json:"id"`

	// 目录
	Category string `json:"category"`

	// 布局ID
	LayoutId string `json:"layout_id"`
}

func (o CreateIndicatorDetailIndicatorType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIndicatorDetailIndicatorType struct{}"
	}

	return strings.Join([]string{"CreateIndicatorDetailIndicatorType", string(data)}, " ")
}
