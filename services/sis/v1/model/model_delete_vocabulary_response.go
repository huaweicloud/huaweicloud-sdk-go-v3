package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVocabularyResponse Response Object
type DeleteVocabularyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVocabularyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVocabularyResponse struct{}"
	}

	return strings.Join([]string{"DeleteVocabularyResponse", string(data)}, " ")
}
