package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateVocabularyResponse struct {
	// 调用成功返回热词表ID，调用失败时无此字段。

	VocabularyId   *string `json:"vocabulary_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateVocabularyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVocabularyResponse struct{}"
	}

	return strings.Join([]string{"UpdateVocabularyResponse", string(data)}, " ")
}
