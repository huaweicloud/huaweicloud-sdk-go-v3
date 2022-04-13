package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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
