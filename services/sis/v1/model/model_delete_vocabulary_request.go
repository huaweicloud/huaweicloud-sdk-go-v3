package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteVocabularyRequest struct {

	// 热词表id。
	VocabularyId string `json:"vocabulary_id"`
}

func (o DeleteVocabularyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVocabularyRequest struct{}"
	}

	return strings.Join([]string{"DeleteVocabularyRequest", string(data)}, " ")
}
