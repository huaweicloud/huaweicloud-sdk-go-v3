package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDesktopRemoteAssistanceInfoResponse Response Object
type ShowDesktopRemoteAssistanceInfoResponse struct {

	// 协同空间ID
	ShareSpaceId *string `json:"share_space_id,omitempty"`

	// 协同空间邀请码(大写英文+数字,共8位)
	InvitationCode *string `json:"invitation_code,omitempty"`

	// 协同空间名称
	ShareSpaceName *string `json:"share_space_name,omitempty"`

	// 协同空间密码
	ShareSpacePasswd *string `json:"share_space_passwd,omitempty"`

	// 专线分享链接
	PrivateShareLink *string `json:"private_share_link,omitempty"`

	// 互联网分享链接
	InternetShareLink *string `json:"internet_share_link,omitempty"`

	// 创建时间 UTC时间，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'。
	CreateTime *string `json:"create_time,omitempty"`

	// 协同空间状态 - OPEN 协同空间已创建 - CLOSE 协同空间已关闭 - WAIT_USER_CONFIRM 等待用户确认进入远程协助 - WAIT_USER_ACCESS 等待用户进入远程协助
	Status *string `json:"status,omitempty"`

	// 失败原因
	FailedReason   *string `json:"failed_reason,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDesktopRemoteAssistanceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopRemoteAssistanceInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowDesktopRemoteAssistanceInfoResponse", string(data)}, " ")
}
