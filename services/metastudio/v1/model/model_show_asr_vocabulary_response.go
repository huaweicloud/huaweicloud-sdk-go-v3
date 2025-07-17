package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAsrVocabularyResponse Response Object
type ShowAsrVocabularyResponse struct {

	// 热词表ID。
	AsrVocabularyId *string `json:"asr_vocabulary_id,omitempty"`

	VocabularyType *AsrVocabularyTypeEnum `json:"vocabulary_type,omitempty"`

	// 奇妙问热词表名
	Name *string `json:"name,omitempty"`

	// 词表内容
	Content *string `json:"content,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAsrVocabularyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAsrVocabularyResponse struct{}"
	}

	return strings.Join([]string{"ShowAsrVocabularyResponse", string(data)}, " ")
}
