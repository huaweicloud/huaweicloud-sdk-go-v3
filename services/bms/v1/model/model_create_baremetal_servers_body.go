package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建裸金属服务器接口请求结构体
type CreateBaremetalServersBody struct {
	Server *CreateServers `json:"server"`
}

func (o CreateBaremetalServersBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBaremetalServersBody struct{}"
	}

	return strings.Join([]string{"CreateBaremetalServersBody", string(data)}, " ")
}
