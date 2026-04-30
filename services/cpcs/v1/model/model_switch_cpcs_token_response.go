package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchCpcsTokenResponse Response Object
type SwitchCpcsTokenResponse struct {
	Token *SwitchTokenResponseToken `json:"token,omitempty"`

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
