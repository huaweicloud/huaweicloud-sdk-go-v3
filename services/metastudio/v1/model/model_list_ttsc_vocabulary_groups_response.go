package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTtscVocabularyGroupsResponse Response Object
type ListTtscVocabularyGroupsResponse struct {

	// 总记录数。
	Count *int32 `json:"count,omitempty"`

	// 自定义词表分组列表。
	Data           *[]TtscGroupAssetsConfig `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListTtscVocabularyGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTtscVocabularyGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListTtscVocabularyGroupsResponse", string(data)}, " ")
}
