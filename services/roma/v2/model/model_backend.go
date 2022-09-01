package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 后端服务定义
type Backend struct {

	// 后端主机地址列表
	Ip *[]string `json:"ip,omitempty" xml:"ip"`

	// 后端服务端口，不存在时使用监听器端口
	Port *int32 `json:"port,omitempty" xml:"port"`

	HealthCheck *HealthCheck `json:"health_check,omitempty" xml:"health_check"`
}

func (o Backend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Backend struct{}"
	}

	return strings.Join([]string{"Backend", string(data)}, " ")
}
