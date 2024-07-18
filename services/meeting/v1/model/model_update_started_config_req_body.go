package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStartedConfigReqBody 修改会议配置请求。
type UpdateStartedConfigReqBody struct {

	// 锁定共享标志位。 * 0: 不锁定 * 1: 锁定
	LockSharing *int32 `json:"lockSharing,omitempty"`

	// 允许加入会议的范围。 - 0: 所有用户 - 2: 企业内用户 - 3: 被邀请用户
	CallInRestriction *int32 `json:"callInRestriction,omitempty"`

	// 是否允许自己解除静音，默认为允许 - 0: 不允许 - 1: 允许
	AllowUnmuteByOneself *int32 `json:"allowUnmuteByOneself,omitempty"`

	// 会议聊天权限 1.全员禁止 2.仅允许私聊 3.仅允许公开聊天 4.允许自由聊天
	ChatPermission *int32 `json:"chatPermission,omitempty"`

	// 网络研讨会观众允许呼入的范围 0：所有用户  2：企业内用户和被邀请用户
	AudienceCallInRestriction *int32 `json:"audienceCallInRestriction,omitempty"`

	// 客户端本地录制权限的范围，默认为仅主持人支持本地录制 - 0: 所有用户 - 1：全部人可录制 - 2：部分人可录制
	ClientRecMode *int32 `json:"clientRecMode,omitempty"`

	// 与会人自行开启摄像头 0:禁止 1:允许
	AllowOpenCamera *int32 `json:"allowOpenCamera,omitempty"`

	// 是否允许与会人改名 0:不允许 1:允许
	AllowRename *int32 `json:"allowRename,omitempty"`

	// 锁定会议 0：解锁 1：锁定
	IsLock *int32 `json:"isLock,omitempty"`

	// 抢共享权限设置 0:仅主持人/联席 1:所有人可抢共享
	FreeShare *int32 `json:"freeShare,omitempty"`
}

func (o UpdateStartedConfigReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStartedConfigReqBody struct{}"
	}

	return strings.Join([]string{"UpdateStartedConfigReqBody", string(data)}, " ")
}
