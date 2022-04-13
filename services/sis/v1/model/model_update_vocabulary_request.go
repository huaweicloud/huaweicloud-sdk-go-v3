package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateVocabularyRequest struct {
	// 被更新的热词表id。

	VocabularyId string `json:"vocabulary_id"`

	Body *PutUpdateVocabReq `json:"body,omitempty"`
}

func (o UpdateVocabularyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVocabularyRequest struct{}"
	}

	return strings.Join([]string{"UpdateVocabularyRequest", string(data)}, " ")
}
