package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 流计算
type StreamModel struct {

	// 输入参数，最多支持10个；流计算的输入参数名需要在接收数据类型为资产数据的实时分析作业中定义，模型中必须与其保持一致
	Inputs []InputModel `json:"inputs" xml:"inputs"`

	// 实时分析作业ID
	JobId string `json:"job_id" xml:"job_id"`

	// 输出属性，最多支持10个
	Outputs []StreamOutput `json:"outputs" xml:"outputs"`
}

func (o StreamModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StreamModel struct{}"
	}

	return strings.Join([]string{"StreamModel", string(data)}, " ")
}
