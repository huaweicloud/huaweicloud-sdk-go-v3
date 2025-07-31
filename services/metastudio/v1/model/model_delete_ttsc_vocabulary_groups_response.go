package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTtscVocabularyGroupsResponse Response Object
type DeleteTtscVocabularyGroupsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTtscVocabularyGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTtscVocabularyGroupsResponse struct{}"
	}

	return strings.Join([]string{"DeleteTtscVocabularyGroupsResponse", string(data)}, " ")
}
