package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomPropsResult 自定义属性任务的返回结果
type CustomPropsResult struct {

	// 自定义属性ID（API侧）
	Id string `json:"id"`

	PropDefinition *PropDefinition `json:"prop_definition"`

	// 自定义属性建模的评估指标集合
	Metrics []CustomPropsModelMetric `json:"metrics"`
}

func (o CustomPropsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomPropsResult struct{}"
	}

	return strings.Join([]string{"CustomPropsResult", string(data)}, " ")
}
