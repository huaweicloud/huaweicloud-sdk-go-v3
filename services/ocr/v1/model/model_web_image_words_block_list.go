package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebImageWordsBlockList
type WebImageWordsBlockList struct {

	// 文字块识别结果。
	Words *string `json:"words,omitempty"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *float32 `json:"confidence,omitempty"`

	// 文字块的区域位置信息，列表形式，包含文字区域四个顶点的二维坐标（x,y）;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	Location *[][]int32 `json:"location,omitempty"`

	// 文字块所属字体类型，列表形式，表示与文字块的文字最接近的字体类型。
	FontList *[]string `json:"font_list,omitempty"`

	// 文字块所属字体类型的概率，列表形式，与font_list一一对应，表示文字块的文字属于某种字体类型的概率。
	FontScores *[]float32 `json:"font_scores,omitempty"`
}

func (o WebImageWordsBlockList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebImageWordsBlockList struct{}"
	}

	return strings.Join([]string{"WebImageWordsBlockList", string(data)}, " ")
}
