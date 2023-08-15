package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVocabularyRequest Request Object
type ShowVocabularyRequest struct {

	// 热词表id。
	VocabularyId string `json:"vocabulary_id"`
}

func (o ShowVocabularyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVocabularyRequest struct{}"
	}

	return strings.Join([]string{"ShowVocabularyRequest", string(data)}, " ")
}
