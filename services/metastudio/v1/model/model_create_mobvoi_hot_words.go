package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMobvoiHotWords mobvoi类型热词
type CreateMobvoiHotWords struct {

	// 热词ID。
	VocabularyId string `json:"vocabulary_id"`

	Language *LanguageEnum `json:"language"`
}

func (o CreateMobvoiHotWords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMobvoiHotWords struct{}"
	}

	return strings.Join([]string{"CreateMobvoiHotWords", string(data)}, " ")
}
