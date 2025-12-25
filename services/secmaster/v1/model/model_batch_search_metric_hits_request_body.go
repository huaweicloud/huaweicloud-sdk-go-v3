package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSearchMetricHitsRequestBody 查询指标Hit结果请求体。
type BatchSearchMetricHitsRequestBody struct {

	// 待查询的指标Id列表, 可参照附录中指标信息说明获取已有指标信息。
	MetricIds []string `json:"metric_ids"`

	// 工作空间列表, 当指标支持获取多工作空间数据时填写。
	WorkspaceIds *[]string `json:"workspace_ids,omitempty"`

	// 待查询指标的参数列表，列表内每个元素为<String, String>的K-V形式，元素数量必须与metric_ids列表相同，具体填写方式请参照附录。
	Params *[]map[string]string `json:"params,omitempty"`

	// 交互式参数查询，当指标支持交互式参数时，填写<String, String>的K-V形式的参数列表，具体填写方式请参照附录。
	InteractiveParams *[]map[string]string `json:"interactive_params,omitempty"`

	// 指标卡片ID列表
	FieldIds *[]string `json:"field_ids,omitempty"`
}

func (o BatchSearchMetricHitsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSearchMetricHitsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchSearchMetricHitsRequestBody", string(data)}, " ")
}
