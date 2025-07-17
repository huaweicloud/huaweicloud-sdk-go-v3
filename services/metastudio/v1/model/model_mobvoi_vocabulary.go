package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MobvoiVocabulary mobvoi类型热词表
type MobvoiVocabulary struct {

	// 词表文本
	Content string `json:"content"`
}

func (o MobvoiVocabulary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MobvoiVocabulary struct{}"
	}

	return strings.Join([]string{"MobvoiVocabulary", string(data)}, " ")
}
