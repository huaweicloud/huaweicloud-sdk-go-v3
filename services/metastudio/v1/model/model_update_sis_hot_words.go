package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSisHotWords sis类型热词
type UpdateSisHotWords struct {

	// 热词ID。
	VocabularyId *string `json:"vocabulary_id,omitempty"`

	// SIS服务所在区域projectId
	SisProjectId *string `json:"sis_project_id,omitempty"`

	// 对接SIS服务的区域。 > 0：北京四；3：上海一；
	Region *int32 `json:"region,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`
}

func (o UpdateSisHotWords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSisHotWords struct{}"
	}

	return strings.Join([]string{"UpdateSisHotWords", string(data)}, " ")
}
