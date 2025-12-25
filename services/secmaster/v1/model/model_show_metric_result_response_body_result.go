package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetricResultResponseBodyResult 指标查询结果内容
type ShowMetricResultResponseBodyResult struct {

	// 指标查询结果表格标题
	Labels []string `json:"labels"`

	// 指标查询结果内容表格
	Datarows [][]interface{} `json:"datarows"`

	// 生效的列, 当有该参数时，使用指定列作为指标数据结果
	EffectiveColumn *string `json:"effective_column,omitempty"`
}

func (o ShowMetricResultResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricResultResponseBodyResult struct{}"
	}

	return strings.Join([]string{"ShowMetricResultResponseBodyResult", string(data)}, " ")
}
