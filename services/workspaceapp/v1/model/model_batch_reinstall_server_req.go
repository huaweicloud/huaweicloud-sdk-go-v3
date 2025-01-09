package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchReinstallServerReq 批量重建服务器的请求体。
type BatchReinstallServerReq struct {

	// 应用服务器id集合。
	ServerIds *[]string `json:"server_ids,omitempty"`

	// 是否自动升级hda版本。
	UpdateAccessAgent *bool `json:"update_access_agent,omitempty"`
}

func (o BatchReinstallServerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchReinstallServerReq struct{}"
	}

	return strings.Join([]string{"BatchReinstallServerReq", string(data)}, " ")
}
