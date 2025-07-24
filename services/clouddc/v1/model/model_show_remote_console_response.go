package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRemoteConsoleResponse Response Object
type ShowRemoteConsoleResponse struct {

	// 远程登录地址
	Url            *string `json:"url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRemoteConsoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRemoteConsoleResponse struct{}"
	}

	return strings.Join([]string{"ShowRemoteConsoleResponse", string(data)}, " ")
}
