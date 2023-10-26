package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisConnectionParam Redis连接参数
type RedisConnectionParam struct {

	// 密码
	Password *string `json:"password,omitempty"`

	// 实例id
	Id *string `json:"id,omitempty"`

	// 地址
	Addrs *string `json:"addrs,omitempty"`

	// 实例名称
	Name *string `json:"name,omitempty"`
}

func (o RedisConnectionParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisConnectionParam struct{}"
	}

	return strings.Join([]string{"RedisConnectionParam", string(data)}, " ")
}
