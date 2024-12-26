package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTtscVocabularyConfigsResponse Response Object
type DeleteTtscVocabularyConfigsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTtscVocabularyConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTtscVocabularyConfigsResponse struct{}"
	}

	return strings.Join([]string{"DeleteTtscVocabularyConfigsResponse", string(data)}, " ")
}
