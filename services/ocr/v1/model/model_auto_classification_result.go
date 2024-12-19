package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoClassificationResult
type AutoClassificationResult struct {
	Status *AutoClassificationResultStatus `json:"status"`

	// 对应票证具体结构化识别的结果。
	Content *interface{} `json:"content"`

	// 对应票证的类别。
	Type string `json:"type"`

	// 票证的区域位置信息，列表形式，包含票证区域四个顶点的二维坐标（x,y）;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	Location [][]int32 `json:"location"`

	// 对应票证中是否含有印章。可选值包括： -  true：该票证中含有印章。 -  false：该票证中不含有印章。
	SealMark *bool `json:"seal_mark,omitempty"`
}

func (o AutoClassificationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoClassificationResult struct{}"
	}

	return strings.Join([]string{"AutoClassificationResult", string(data)}, " ")
}
