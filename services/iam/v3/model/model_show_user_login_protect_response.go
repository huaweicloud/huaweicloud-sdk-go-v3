package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowUserLoginProtectResponse struct {
	LoginProtect   *LoginProtectResult `json:"login_protect,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowUserLoginProtectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserLoginProtectResponse struct{}"
	}

	return strings.Join([]string{"ShowUserLoginProtectResponse", string(data)}, " ")
}
