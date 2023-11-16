package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableNoticeRequest Request Object
type DisableNoticeRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	// 通知的类型,分为消息，邮件和钉钉
	NoticeType string `json:"notice_type"`
}

func (o DisableNoticeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableNoticeRequest struct{}"
	}

	return strings.Join([]string{"DisableNoticeRequest", string(data)}, " ")
}
