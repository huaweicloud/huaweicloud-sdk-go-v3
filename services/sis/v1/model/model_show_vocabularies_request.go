package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVocabulariesRequest struct {

	// 页码偏移量，表示从此页码偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	Body *ShowVocabulariesParams `json:"body,omitempty"`
}

func (o ShowVocabulariesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVocabulariesRequest struct{}"
	}

	return strings.Join([]string{"ShowVocabulariesRequest", string(data)}, " ")
}
