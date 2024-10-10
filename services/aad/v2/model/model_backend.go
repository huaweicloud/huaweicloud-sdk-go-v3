package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Backend backend
type Backend struct {

	// 当前后端协议
	Protocol *string `json:"protocol,omitempty"`

	// 当前后端端口
	Port *int32 `json:"port,omitempty"`

	// 当前后端 Host 值
	Host *string `json:"host,omitempty"`
}

func (o Backend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Backend struct{}"
	}

	return strings.Join([]string{"Backend", string(data)}, " ")
}
