package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecommendWordsResponse Response Object
type ListRecommendWordsResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 推荐词列表
	RecommendWords *[]RecommendWord `json:"recommend_words,omitempty"`

	// 推荐词总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRecommendWordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecommendWordsResponse struct{}"
	}

	return strings.Join([]string{"ListRecommendWordsResponse", string(data)}, " ")
}
