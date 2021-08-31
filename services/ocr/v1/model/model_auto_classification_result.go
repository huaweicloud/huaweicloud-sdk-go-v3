package model

import (
	"encoding/json"

	"strings"
)

//
type AutoClassificationResult struct {
	// 指示各对应票证的识别状态。

	Status *interface{} `json:"status"`
	// 对应票证具体结构化识别的结果。

	Content *interface{} `json:"content"`
	// 对应票证的类别。

	Type string `json:"type"`
	// 文字块的区域位置信息，列表形式，包含文字区域四个顶点的二维坐标（x,y）;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。

	Location [][]int32 `json:"location"`
	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。

	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o AutoClassificationResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AutoClassificationResult struct{}"
	}

	return strings.Join([]string{"AutoClassificationResult", string(data)}, " ")
}
