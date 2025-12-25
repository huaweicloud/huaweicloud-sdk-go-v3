package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LayoutMetricFormat struct {

	// 数据格式
	Data *string `json:"data,omitempty"`

	// 显示格式
	Display *string `json:"display,omitempty"`

	// 显示参数
	DisplayParam map[string]string `json:"display_param,omitempty"`

	// 数据参数
	DataParam map[string]string `json:"data_param,omitempty"`
}

func (o LayoutMetricFormat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayoutMetricFormat struct{}"
	}

	return strings.Join([]string{"LayoutMetricFormat", string(data)}, " ")
}
