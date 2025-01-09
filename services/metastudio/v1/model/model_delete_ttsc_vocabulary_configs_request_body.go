package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteTtscVocabularyConfigsRequestBody struct {

	// 自定义读法id
	Id *[]string `json:"id,omitempty"`
}

func (o DeleteTtscVocabularyConfigsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTtscVocabularyConfigsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteTtscVocabularyConfigsRequestBody", string(data)}, " ")
}
