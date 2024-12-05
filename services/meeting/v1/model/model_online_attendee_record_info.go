package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OnlineAttendeeRecordInfo 在线与会者信息
type OnlineAttendeeRecordInfo struct {

	// 与会者标识
	ParticipantId *string `json:"participant_id,omitempty"`

	// 与会者名称
	Name *string `json:"name,omitempty"`

	// 呼叫号码
	CallNumber *string `json:"call_number,omitempty"`

	// 会议中的角色，枚举值如下： 1：会议主席 0：普通与会者
	Role *int32 `json:"role,omitempty"`

	// 开放性场景标识第三方账号信息
	ThirdAccount *string `json:"third_account,omitempty"`

	// 用户账号
	Account *string `json:"account,omitempty"`

	// 用户UUID
	UserId *string `json:"user_id,omitempty"`
}

func (o OnlineAttendeeRecordInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OnlineAttendeeRecordInfo struct{}"
	}

	return strings.Join([]string{"OnlineAttendeeRecordInfo", string(data)}, " ")
}
