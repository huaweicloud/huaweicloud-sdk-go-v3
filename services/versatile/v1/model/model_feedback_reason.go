package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FeedbackReason 用户反馈原因
type FeedbackReason struct {

	// 原因标签列表
	Tags []string `json:"tags"`

	// 具体原因内容
	Content *string `json:"content,omitempty"`
}

func (o FeedbackReason) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FeedbackReason struct{}"
	}

	return strings.Join([]string{"FeedbackReason", string(data)}, " ")
}
