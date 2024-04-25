package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerFormulaBlock struct {

	// 数学公式识别结果，以latex字符串表示。
	Formula *string `json:"formula,omitempty"`

	// 数学公式位置信息，列表形式，分别表示4个顶点的x, y坐标；坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	Location *[][]int32 `json:"location,omitempty"`
}

func (o SmartDocumentRecognizerFormulaBlock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerFormulaBlock struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerFormulaBlock", string(data)}, " ")
}
