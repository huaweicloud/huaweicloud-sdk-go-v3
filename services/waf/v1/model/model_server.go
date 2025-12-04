package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Server struct {

	// 类型
	Type *string `json:"type,omitempty"`

	// 地址
	Address *string `json:"address,omitempty"`
}

func (o Server) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Server struct{}"
	}

	return strings.Join([]string{"Server", string(data)}, " ")
}
