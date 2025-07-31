package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SaveTtscVocabularyGroupsRequestBody struct {

	// 分组名称
	GroupName *string `json:"group_name,omitempty"`
}

func (o SaveTtscVocabularyGroupsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTtscVocabularyGroupsRequestBody struct{}"
	}

	return strings.Join([]string{"SaveTtscVocabularyGroupsRequestBody", string(data)}, " ")
}
