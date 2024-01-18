package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReinstallServerReq 重建服务器的请求体。
type ReinstallServerReq struct {

	// 是否自动升级hda版本。
	UpdateAccessAgent *bool `json:"update_access_agent,omitempty"`
}

func (o ReinstallServerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallServerReq struct{}"
	}

	return strings.Join([]string{"ReinstallServerReq", string(data)}, " ")
}
