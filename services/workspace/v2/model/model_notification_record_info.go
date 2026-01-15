package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationRecordInfo 通知拦截记录信息详情
type NotificationRecordInfo struct {

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 发送SMN类型:EMAIL-邮件、SMS-短信
	Type *string `json:"type,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 桌面名
	DesktopName *string `json:"desktop_name,omitempty"`

	// 桌面池名称
	DesktopPoolName *string `json:"desktop_pool_name,omitempty"`

	// 发送结果
	Result *string `json:"result,omitempty"`

	// 报错信息
	ErrorMessage *string `json:"error_message,omitempty"`

	// 操作时间
	OperateTime *string `json:"operate_time,omitempty"`
}

func (o NotificationRecordInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationRecordInfo struct{}"
	}

	return strings.Join([]string{"NotificationRecordInfo", string(data)}, " ")
}
