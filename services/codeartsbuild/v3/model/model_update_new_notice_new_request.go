package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNewNoticeNewRequest Request Object
type UpdateNewNoticeNewRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	Body *UpdateNoticeRequestBody `json:"body,omitempty"`
}

func (o UpdateNewNoticeNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNewNoticeNewRequest struct{}"
	}

	return strings.Join([]string{"UpdateNewNoticeNewRequest", string(data)}, " ")
}
