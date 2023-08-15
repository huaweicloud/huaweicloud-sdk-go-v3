package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TcpSocketDto struct {

	// 端口
	Port int32 `json:"port"`
}

func (o TcpSocketDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TcpSocketDto struct{}"
	}

	return strings.Join([]string{"TcpSocketDto", string(data)}, " ")
}
