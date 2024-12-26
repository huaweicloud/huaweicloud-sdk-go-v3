package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTtscVocabularyConfigsResponse Response Object
type CreateTtscVocabularyConfigsResponse struct {

	// 配置项id。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTtscVocabularyConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTtscVocabularyConfigsResponse struct{}"
	}

	return strings.Join([]string{"CreateTtscVocabularyConfigsResponse", string(data)}, " ")
}
