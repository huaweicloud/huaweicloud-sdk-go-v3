package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowIpcsResponse Response Object
type BatchShowIpcsResponse struct {

	// **参数说明**：总数
	Count *int64 `json:"count,omitempty"`

	// **参数说明**：IPC列表
	Ipcs           *[]IpcResponseDto `json:"ipcs,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o BatchShowIpcsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowIpcsResponse struct{}"
	}

	return strings.Join([]string{"BatchShowIpcsResponse", string(data)}, " ")
}
