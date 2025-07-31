package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTtscVocabularyGroupsResponse Response Object
type UpdateTtscVocabularyGroupsResponse struct {

	// 配置项id。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTtscVocabularyGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTtscVocabularyGroupsResponse struct{}"
	}

	return strings.Join([]string{"UpdateTtscVocabularyGroupsResponse", string(data)}, " ")
}
