package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateFeedbackReq struct {

	// 任务类型，例如 ecs诊断任务 1，rds诊断任务 2
	Type int32 `json:"type"`

	// 反馈内容
	Content *string `json:"content,omitempty"`

	// 是否有帮助
	IsHelpful bool `json:"is_helpful"`
}

func (o CreateFeedbackReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFeedbackReq struct{}"
	}

	return strings.Join([]string{"CreateFeedbackReq", string(data)}, " ")
}
