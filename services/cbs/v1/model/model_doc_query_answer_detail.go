package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type DocQueryAnswerDetail struct {

	// 答案。
	Answer string `json:"answer" xml:"answer"`

	// 文档ID。
	DocId string `json:"doc_id" xml:"doc_id"`

	// 答案结束下标。
	EndIndex int32 `json:"end_index" xml:"end_index"`

	// 段落评分。
	ParagraphScore float64 `json:"paragraph_score" xml:"paragraph_score"`

	// 段落文字。
	ParagraphText string `json:"paragraph_text" xml:"paragraph_text"`

	// 文档问答阅读理解评分。
	PhraseScore float64 `json:"phrase_score" xml:"phrase_score"`

	// 答案开始下标。
	StartIndex int32 `json:"start_index" xml:"start_index"`

	// 文档问答总评分。
	TotalScore float64 `json:"total_score" xml:"total_score"`

	// 段落在文档中的编号。
	ParagraphNumber int32 `json:"paragraph_number" xml:"paragraph_number"`
}

func (o DocQueryAnswerDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocQueryAnswerDetail struct{}"
	}

	return strings.Join([]string{"DocQueryAnswerDetail", string(data)}, " ")
}
