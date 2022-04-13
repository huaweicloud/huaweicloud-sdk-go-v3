package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 专属主机上创建的云服务器规格。
type RespInstanceCapacity struct {
	// 支持创建的云服务器规格。

	Flavor string `json:"flavor"`
}

func (o RespInstanceCapacity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespInstanceCapacity struct{}"
	}

	return strings.Join([]string{"RespInstanceCapacity", string(data)}, " ")
}
