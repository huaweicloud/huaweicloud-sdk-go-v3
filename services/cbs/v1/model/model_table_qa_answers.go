package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type TableQaAnswers struct {

	// 表格问答答案。
	Answer string `json:"answer" xml:"answer"`

	// 评分。
	Score float64 `json:"score" xml:"score"`

	// 表格ID。
	TableId string `json:"table_id" xml:"table_id"`
}

func (o TableQaAnswers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableQaAnswers struct{}"
	}

	return strings.Join([]string{"TableQaAnswers", string(data)}, " ")
}
