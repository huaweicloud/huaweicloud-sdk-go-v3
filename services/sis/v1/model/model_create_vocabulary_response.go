package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateVocabularyResponse struct {
	// 调用成功返回热词表ID，调用失败时无此字段。

	VocabularyId   *string `json:"vocabulary_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVocabularyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVocabularyResponse struct{}"
	}

	return strings.Join([]string{"CreateVocabularyResponse", string(data)}, " ")
}
