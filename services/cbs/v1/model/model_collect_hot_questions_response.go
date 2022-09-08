package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CollectHotQuestionsResponse struct {

	// 指定时间范围内，热点问题列表。
	Questions      *[]HotQuestionCount `json:"questions,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CollectHotQuestionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectHotQuestionsResponse struct{}"
	}

	return strings.Join([]string{"CollectHotQuestionsResponse", string(data)}, " ")
}
