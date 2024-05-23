package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHotWordsResponse Response Object
type UpdateHotWordsResponse struct {

	// 热词记录ID。
	HotWordsId *string `json:"hot_words_id,omitempty"`

	// 应用ID。
	RobotId *string `json:"robot_id,omitempty"`

	HotWordsType *HotWordsTypeEnum `json:"hot_words_type,omitempty"`

	// 热词ID(sis中配置)。
	VocabularyId *string `json:"vocabulary_id,omitempty"`

	// SIS服务所在区域projectId
	SisProjectId *string `json:"sis_project_id,omitempty"`

	// 对接SIS服务的区域。 > 0：北京四；3：上海一；
	Region *int32 `json:"region,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateHotWordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHotWordsResponse struct{}"
	}

	return strings.Join([]string{"UpdateHotWordsResponse", string(data)}, " ")
}
