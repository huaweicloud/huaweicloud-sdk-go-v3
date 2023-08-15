package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWindowsBaremetalServerPwdResponse Response Object
type ShowWindowsBaremetalServerPwdResponse struct {

	// 加密后的密码
	Password       *string `json:"password,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowWindowsBaremetalServerPwdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWindowsBaremetalServerPwdResponse struct{}"
	}

	return strings.Join([]string{"ShowWindowsBaremetalServerPwdResponse", string(data)}, " ")
}
