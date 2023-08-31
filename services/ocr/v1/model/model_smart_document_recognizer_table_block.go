package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerTableBlock struct {

	// 当前表格的位置信息，列表形式，分别表示文字块4个顶点的x, y坐标；坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	Location *[][]int32 `json:"location,omitempty"`

	// 表格中所包含的单元格数量。
	WordsBlockCount *int32 `json:"words_block_count,omitempty"`

	// 单元格识别结果列表。
	WordsBlockList *[]SmartDocumentRecognizerTableWordsBlock `json:"words_block_list,omitempty"`

	// 表格识别结果的base64编码，仅当return_excel为True时返回该字段。对返回的excel编码可用base64.b64decode解码并保存为.xlsx文件。
	Excel *string `json:"excel,omitempty"`
}

func (o SmartDocumentRecognizerTableBlock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerTableBlock struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerTableBlock", string(data)}, " ")
}
