package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVocabulariesResponse Response Object
type ShowVocabulariesResponse struct {

	// 热词表数。
	Count *int32 `json:"count,omitempty"`

	// 调用成功返回热词表列表，调用失败时无此字段。
	Result         *[]VocabInfo `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowVocabulariesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVocabulariesResponse struct{}"
	}

	return strings.Join([]string{"ShowVocabulariesResponse", string(data)}, " ")
}
