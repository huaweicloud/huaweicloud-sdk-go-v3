package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuestionAnswerResponse Response Object
type ListQuestionAnswerResponse struct {

	// 与第一条数据的偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 页面大小
	Limit *int32 `json:"limit,omitempty"`

	// 总数量
	Count *int32 `json:"count,omitempty"`

	// 技能信息
	Data *[]QuestionAnswerInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListQuestionAnswerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuestionAnswerResponse struct{}"
	}

	return strings.Join([]string{"ListQuestionAnswerResponse", string(data)}, " ")
}
