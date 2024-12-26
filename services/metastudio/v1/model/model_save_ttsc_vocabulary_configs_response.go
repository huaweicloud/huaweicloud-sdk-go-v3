package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveTtscVocabularyConfigsResponse Response Object
type SaveTtscVocabularyConfigsResponse struct {

	// 配置项id。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SaveTtscVocabularyConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTtscVocabularyConfigsResponse struct{}"
	}

	return strings.Join([]string{"SaveTtscVocabularyConfigsResponse", string(data)}, " ")
}
