package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAsrVocabularyReq 创建热词表请求。
type CreateAsrVocabularyReq struct {
	VocabularyType *AsrVocabularyTypeEnum `json:"vocabulary_type"`

	// 奇妙问热词表名
	Name string `json:"name"`

	Language *LanguageEnum `json:"language"`

	MobvoiVocabulary *MobvoiVocabulary `json:"mobvoi_vocabulary,omitempty"`
}

func (o CreateAsrVocabularyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAsrVocabularyReq struct{}"
	}

	return strings.Join([]string{"CreateAsrVocabularyReq", string(data)}, " ")
}
