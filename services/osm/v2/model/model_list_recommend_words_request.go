package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRecommendWordsRequest struct {

	// 调用智能客服服务标志。
	XServiceKey *string `json:"x-service-key,omitempty"`

	// 站点标记，0-中国站  1-国际站
	XSite *string `json:"x-site,omitempty"`

	// 区域语言简写，en-us  zh-cn
	XLanguage *string `json:"x-language,omitempty"`

	// 推荐词Id
	RecommendWordId *string `json:"recommend_word_id,omitempty"`

	// 推荐词层级
	LevelValue *int32 `json:"level_value,omitempty"`

	// 主题名称
	ThemeName *string `json:"theme_name,omitempty"`
}

func (o ListRecommendWordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecommendWordsRequest struct{}"
	}

	return strings.Join([]string{"ListRecommendWordsRequest", string(data)}, " ")
}
