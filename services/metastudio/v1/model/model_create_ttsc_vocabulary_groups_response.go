package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTtscVocabularyGroupsResponse Response Object
type CreateTtscVocabularyGroupsResponse struct {

	// 配置项id。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTtscVocabularyGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTtscVocabularyGroupsResponse struct{}"
	}

	return strings.Join([]string{"CreateTtscVocabularyGroupsResponse", string(data)}, " ")
}
