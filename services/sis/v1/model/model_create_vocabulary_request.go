package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVocabularyRequest Request Object
type CreateVocabularyRequest struct {
	Body *PostCreateVocabReq `json:"body,omitempty"`
}

func (o CreateVocabularyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVocabularyRequest struct{}"
	}

	return strings.Join([]string{"CreateVocabularyRequest", string(data)}, " ")
}
