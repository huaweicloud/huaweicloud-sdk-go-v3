package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRemoteAssistanceResponse Response Object
type CreateRemoteAssistanceResponse struct {

	// 协同空间ID
	ShareSpaceId *string `json:"share_space_id,omitempty"`

	// 协同空间状态 - OPEN 协同空间已创建 - CLOSE 协同空间已关闭 - WAIT_USER_CONFIRM 等待用户确认进入远程协助 - WAIT_USER_ACCESS 等待用户进入远程协助
	Status *string `json:"status,omitempty"`

	// 桌面的desktopId
	DesktopId *string `json:"desktop_id,omitempty"`

	// 发起方类型 - ADMIN_INITIATE 管理员发起 - ENDUSER_INITIATE 终端用户发起
	InitiatorType  *string `json:"initiator_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRemoteAssistanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRemoteAssistanceResponse struct{}"
	}

	return strings.Join([]string{"CreateRemoteAssistanceResponse", string(data)}, " ")
}
