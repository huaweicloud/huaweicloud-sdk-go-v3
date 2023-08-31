package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FeedbackInfo struct {

	// 编号
	Id *string `json:"id,omitempty"`

	// 项目Id
	ProjectId *string `json:"project_id,omitempty"`

	// 任务消息唯一Id
	MessageId *string `json:"message_id,omitempty"`

	// 反馈等级
	FeedbackGrade *string `json:"feedback_grade,omitempty"`

	// 反馈内容
	FeedbackContent *string `json:"feedback_content,omitempty"`

	// 创建时间
	GmtCreated *int64 `json:"gmt_created,omitempty"`

	// 修改时间
	GmtModified *int64 `json:"gmt_modified,omitempty"`
}

func (o FeedbackInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FeedbackInfo struct{}"
	}

	return strings.Join([]string{"FeedbackInfo", string(data)}, " ")
}
