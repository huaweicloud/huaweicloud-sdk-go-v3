package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTtscVocabularyConfigsResponse Response Object
type ListTtscVocabularyConfigsResponse struct {

	// 总记录数。
	Count *int32 `json:"count,omitempty"`

	// 自定义读法。
	Data           *[]VocabularyConfig `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListTtscVocabularyConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTtscVocabularyConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListTtscVocabularyConfigsResponse", string(data)}, " ")
}
