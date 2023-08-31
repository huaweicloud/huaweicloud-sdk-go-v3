package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerTableWordsBlock struct {

	// 单元格内的文字识别结果。
	Words *string `json:"words,omitempty"`

	// 文字块占用的行信息，编号从0开始，列表形式，数据类型为Integer。
	Rows *[]int32 `json:"rows,omitempty"`

	// 文字块占用的列信息，编号从0开始，列表形式，数据类型为Integer。
	Columns *[]int32 `json:"columns,omitempty"`
}

func (o SmartDocumentRecognizerTableWordsBlock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerTableWordsBlock struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerTableWordsBlock", string(data)}, " ")
}
