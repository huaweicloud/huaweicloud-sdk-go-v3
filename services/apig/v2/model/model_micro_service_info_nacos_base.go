package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MicroServiceInfoNacosBase Nacos微服务详细信息。
type MicroServiceInfoNacosBase struct {

	// 命名空间ID，当选择默认命名空间public时，此项为空。由字母、数字、连接符('-')、下划线('_')组成且64个字符之内。
	Namespace *string `json:"namespace,omitempty"`

	// 集群名称，默认为DEFAULT。由字母、数字、连接符('-')、下划线('_')组成且64个字符之内。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 分组名称，默认为DEFAULT_GROUP。由字母、数字、连接符('-')、下划线('_')、点号('.')、冒号(':')组成且128个字符之内。
	GroupName *string `json:"group_name,omitempty"`

	// 微服务名称。不包含中文和@@，不得以@开头，512个字符以内。
	ServiceName string `json:"service_name"`

	// nacos服务端配置信息。
	ServerConfig []NacosServerConfig `json:"server_config"`

	UserInfo *NacosUserInfo `json:"user_info"`
}

func (o MicroServiceInfoNacosBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroServiceInfoNacosBase struct{}"
	}

	return strings.Join([]string{"MicroServiceInfoNacosBase", string(data)}, " ")
}
