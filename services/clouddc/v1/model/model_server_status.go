package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerStatus 服务器运行状态响应对象
type ServerStatus struct {
	Total *ServerState `json:"total"`

	// 服务器运行OK状态所有对象
	Ok *interface{} `json:"ok"`

	// 服务器运行Warning状态所有对象
	Warning *interface{} `json:"warning"`

	// 服务器运行Critical状态所有对象
	Critical *interface{} `json:"critical"`

	// 服务器运行Health状态所有对象
	Health *interface{} `json:"health"`

	// 服务器运行UnHealth状态所有对象
	Unhealth *interface{} `json:"unhealth,omitempty"`

	// 服务器运行Isolation状态所有对象
	Isolation *interface{} `json:"isolation"`
}

func (o ServerStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerStatus struct{}"
	}

	return strings.Join([]string{"ServerStatus", string(data)}, " ")
}
