package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchCpcsTokenResponse Response Object
type SwitchCpcsTokenResponse struct {

	// 角色列表
	Roles *[]string `json:"roles,omitempty"`

	Ak *SwitchTokenResponseAk `json:"ak,omitempty"`

	// 过期时间
	ExpiredAt *string `json:"expired_at,omitempty"`

	// 签发时间
	IssuedAt *string `json:"issued_at,omitempty"`

	User *SwitchTokenResponseUser `json:"user,omitempty"`

	XCPCSToken     *string `json:"X-CPCS-Token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchCpcsTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchCpcsTokenResponse struct{}"
	}

	return strings.Join([]string{"SwitchCpcsTokenResponse", string(data)}, " ")
}
