package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// server字段数据结构说明
type ChangeBaremetalNameServer struct {

	// 修改后的裸金属服务器名称
	Name string `json:"name"`
}

func (o ChangeBaremetalNameServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeBaremetalNameServer struct{}"
	}

	return strings.Join([]string{"ChangeBaremetalNameServer", string(data)}, " ")
}
