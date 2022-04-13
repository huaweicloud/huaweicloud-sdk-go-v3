package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 流计算输出映射
type StreamOutput struct {
	// 输出参数名称，必须是接收数据类型为资产数据的实时分析作业中已定义的

	Name string `json:"name"`
	// 输出属性名，必须是本模型分析任务类别的属性的属性名

	OutputProperty string `json:"output_property"`
}

func (o StreamOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StreamOutput struct{}"
	}

	return strings.Join([]string{"StreamOutput", string(data)}, " ")
}
