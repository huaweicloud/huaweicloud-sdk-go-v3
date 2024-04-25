package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerLayoutBlock struct {

	// 文字块的区域位置信息，列表形式，包含文字区域四个顶点的二维坐标（x,y）;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	Location *[][]int32 `json:"location,omitempty"`

	// 文档区域类别，包含text、title、sub_title、image、image_caption、form、table、table_caption、header、footer、page_number、reference、formula、stamp、directory共15个类别。
	Type *string `json:"type,omitempty"`

	// 文档区域文字内容。对于表格与图像，不返回其中的文字内容。
	Text *string `json:"text,omitempty"`

	// 文字识别结果索引列表，表示ocr_result的words_block_list中哪些文本框位于该文档区域内。
	WordsIds *[]int32 `json:"words_ids,omitempty"`

	// 仅当type为\"table\"且入参table为True时返回该字段，表示当前逻辑表格区域对应table_result中哪一项识别结果。
	TableId *int32 `json:"table_id,omitempty"`

	// 仅当type为\"form\"且入参form为True时返回该字段，表示当前有线表单区域对应form_result中哪一项识别结果。
	FormId *int32 `json:"form_id,omitempty"`
}

func (o SmartDocumentRecognizerLayoutBlock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerLayoutBlock struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerLayoutBlock", string(data)}, " ")
}
