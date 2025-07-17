package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAsrVocabularyReq 修改热词表请求。
type UpdateAsrVocabularyReq struct {
	VocabularyType *AsrVocabularyTypeEnum `json:"vocabulary_type"`

	// 奇妙问热词表名
	Name *string `json:"name,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`

	MobvoiVocabulary *MobvoiVocabulary `json:"mobvoi_vocabulary,omitempty"`
}

func (o UpdateAsrVocabularyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAsrVocabularyReq struct{}"
	}

	return strings.Join([]string{"UpdateAsrVocabularyReq", string(data)}, " ")
}
