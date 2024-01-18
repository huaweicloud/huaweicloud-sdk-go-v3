package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpnResourceTag 在VPN服务的各个创建资源接口中使用，只面向租户
type VpnResourceTag struct {

	// 标签的key
	Key string `json:"key"`

	// 标签的value
	Value string `json:"value"`
}

func (o VpnResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpnResourceTag struct{}"
	}

	return strings.Join([]string{"VpnResourceTag", string(data)}, " ")
}
