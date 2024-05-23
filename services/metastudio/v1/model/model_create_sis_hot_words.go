package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSisHotWords sis类型热词
type CreateSisHotWords struct {

	// 热词ID。
	VocabularyId string `json:"vocabulary_id"`

	// SIS服务所在区域projectId
	SisProjectId string `json:"sis_project_id"`

	// 对接SIS服务的区域。 > 0：北京四；3：上海一；
	Region int32 `json:"region"`

	Language *LanguageEnum `json:"language"`
}

func (o CreateSisHotWords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSisHotWords struct{}"
	}

	return strings.Join([]string{"CreateSisHotWords", string(data)}, " ")
}
