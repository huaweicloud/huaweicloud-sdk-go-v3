package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AsrVocabularyAssociation 热词表关联信息
type AsrVocabularyAssociation struct {

	// 热词记录ID
	HotWordsId *string `json:"hot_words_id,omitempty"`

	// 热词表ID
	VocabularyId *string `json:"vocabulary_id,omitempty"`

	// 应用ID。
	RobotId *string `json:"robot_id,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`
}

func (o AsrVocabularyAssociation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsrVocabularyAssociation struct{}"
	}

	return strings.Join([]string{"AsrVocabularyAssociation", string(data)}, " ")
}
