package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderLog 服务单日志
type OrderLog struct {

	// 阶段
	Stage *string `json:"stage,omitempty"`

	// 处理人
	Handler *string `json:"handler,omitempty"`

	// 处理时间
	Time *string `json:"time,omitempty"`

	// 处理动作
	Action *string `json:"action,omitempty"`

	// 处理说明
	Description *string `json:"description,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 附件
	Attachments *[]UploadFileInfo `json:"attachments,omitempty"`
}

func (o OrderLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderLog struct{}"
	}

	return strings.Join([]string{"OrderLog", string(data)}, " ")
}
