package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyDesktopAttributesReqDesktop 桌面属性。
type ModifyDesktopAttributesReqDesktop struct {

	// 桌面名。
	ComputerName *string `json:"computer_name,omitempty"`

	// 是否开启快照的操作类型,\"0\":关闭 \"1\":开启。
	SelfBackupManagement *string `json:"self_backup_management,omitempty"`
}

func (o ModifyDesktopAttributesReqDesktop) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDesktopAttributesReqDesktop struct{}"
	}

	return strings.Join([]string{"ModifyDesktopAttributesReqDesktop", string(data)}, " ")
}
