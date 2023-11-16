package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNoticeRequest Request Object
type UpdateNoticeRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	Body *UpdateNoticeRequestBody `json:"body,omitempty"`
}

func (o UpdateNoticeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNoticeRequest struct{}"
	}

	return strings.Join([]string{"UpdateNoticeRequest", string(data)}, " ")
}
