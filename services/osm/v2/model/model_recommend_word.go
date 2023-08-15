package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecommendWord struct {

	// 推荐词Id
	RecommendWordId *string `json:"recommend_word_id,omitempty"`

	// 推荐词
	RecommendWordName *string `json:"recommend_word_name,omitempty"`

	// 推荐词层级
	LevelValue *int32 `json:"level_value,omitempty"`

	// 推荐词排序，序号越小越靠前
	SortValue *int32 `json:"sort_value,omitempty"`

	// 主题Id
	ThemeId *string `json:"theme_id,omitempty"`

	// 主题名称
	ThemeName *string `json:"theme_name,omitempty"`

	AnswerInfo *AnswerInfo `json:"answer_info,omitempty"`
}

func (o RecommendWord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecommendWord struct{}"
	}

	return strings.Join([]string{"RecommendWord", string(data)}, " ")
}
