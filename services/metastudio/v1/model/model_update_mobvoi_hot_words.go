package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMobvoiHotWords mobvoi类型热词
type UpdateMobvoiHotWords struct {

	// 热词ID。
	VocabularyId *string `json:"vocabulary_id,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`
}

func (o UpdateMobvoiHotWords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMobvoiHotWords struct{}"
	}

	return strings.Join([]string{"UpdateMobvoiHotWords", string(data)}, " ")
}
