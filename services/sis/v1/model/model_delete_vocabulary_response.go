package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVocabularyResponse Response Object
type DeleteVocabularyResponse struct {

	// 删除的热词表的名称。
	VocabularyId *string `json:"vocabulary_id,omitempty"`

	// 删除成功的消息。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteVocabularyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVocabularyResponse struct{}"
	}

	return strings.Join([]string{"DeleteVocabularyResponse", string(data)}, " ")
}
