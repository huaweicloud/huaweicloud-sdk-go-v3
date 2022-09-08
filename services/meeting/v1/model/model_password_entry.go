package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 密码信息
type PasswordEntry struct {

	// 会议角色。 - chair: 会议主持人。 - general: 普通与会者。
	ConferenceRole *string `json:"conferenceRole,omitempty"`

	// 会议中角色的密码（明文）。
	Password *string `json:"password,omitempty"`
}

func (o PasswordEntry) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PasswordEntry struct{}"
	}

	return strings.Join([]string{"PasswordEntry", string(data)}, " ")
}
