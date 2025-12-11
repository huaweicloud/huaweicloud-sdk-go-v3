package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerWordsBlockList struct {

	// 文字块识别结果。
	Words *string `json:"words,omitempty"`

	// 文字块的区域位置信息，列表形式，包含文字区域四个顶点的二维坐标（x,y）;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	Location *[][]int32 `json:"location,omitempty"`

	// 文字块识别结果的置信度。
	Confidence *float32 `json:"confidence,omitempty"`

	// 当入参character_mode为True时返回该字段，表示当前文字块对应的单字符识别列表，输出顺序从左到右，先上后下。
	CharList *[]SmartDocumentRecognizerCharList `json:"char_list,omitempty"`
}

func (o SmartDocumentRecognizerWordsBlockList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerWordsBlockList struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerWordsBlockList", string(data)}, " ")
}
