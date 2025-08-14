package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAsrVocabularyAssociationResponse Response Object
type ShowAsrVocabularyAssociationResponse struct {

	// 页面起始页,从0开始
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 总数量
	Count *int32 `json:"count,omitempty"`

	// 热词表关联信息
	Data *[]AsrVocabularyAssociation `json:"data,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAsrVocabularyAssociationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAsrVocabularyAssociationResponse struct{}"
	}

	return strings.Join([]string{"ShowAsrVocabularyAssociationResponse", string(data)}, " ")
}
