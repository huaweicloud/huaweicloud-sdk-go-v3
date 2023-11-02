package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndicatorDataObjectDetailIndicatorType 情报类型对象
type IndicatorDataObjectDetailIndicatorType struct {

	// 情报类型
	IndicatorType *string `json:"indicator_type,omitempty"`

	// 情报类型ID
	Id *string `json:"id,omitempty"`

	// 目录
	Category *string `json:"category,omitempty"`

	// 布局ID
	LayoutId *string `json:"layout_id,omitempty"`
}

func (o IndicatorDataObjectDetailIndicatorType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorDataObjectDetailIndicatorType struct{}"
	}

	return strings.Join([]string{"IndicatorDataObjectDetailIndicatorType", string(data)}, " ")
}
