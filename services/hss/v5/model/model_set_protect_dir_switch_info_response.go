package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetProtectDirSwitchInfoResponse Response Object
type SetProtectDirSwitchInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetProtectDirSwitchInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetProtectDirSwitchInfoResponse struct{}"
	}

	return strings.Join([]string{"SetProtectDirSwitchInfoResponse", string(data)}, " ")
}
