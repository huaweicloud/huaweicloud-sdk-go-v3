package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FeedbackOption struct {

	// 反馈选项id
	FeedbackOptionId *string `json:"feedback_option_id,omitempty"`

	// 反馈选项名称
	FeedbackOptionName *string `json:"feedback_option_name,omitempty"`
}

func (o FeedbackOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FeedbackOption struct{}"
	}

	return strings.Join([]string{"FeedbackOption", string(data)}, " ")
}
