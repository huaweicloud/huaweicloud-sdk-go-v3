package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerStatus 服务器运行状态响应对象
type ServerStatus struct {
	Total *ServerState `json:"total"`

	Ok *ServerState `json:"ok"`

	Warning *ServerState `json:"warning"`

	Unknown *ServerState `json:"unknown,omitempty"`

	Critical *ServerState `json:"critical"`

	Health *ServerState `json:"health"`

	Unhealth *ServerState `json:"unhealth,omitempty"`

	Isolation *ServerState `json:"isolation"`
}

func (o ServerStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerStatus struct{}"
	}

	return strings.Join([]string{"ServerStatus", string(data)}, " ")
}
