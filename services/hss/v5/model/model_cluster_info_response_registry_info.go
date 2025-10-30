package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterInfoResponseRegistryInfo 镜像仓信息
type ClusterInfoResponseRegistryInfo struct {

	// **参数解释** 镜像仓库类型 **取值范围** 字符长度1-256位
	RegistryType *string `json:"registry_type,omitempty"`

	// **参数解释** 镜像仓地址 **取值范围** 字符长度1-256位
	RegistryAddr *string `json:"registry_addr,omitempty"`

	// **参数解释** 镜像仓用户名 **取值范围** 字符长度1-256位
	RegistryUsername *string `json:"registry_username,omitempty"`

	// **参数解释** 组织 **取值范围** 字符长度1-256位
	Namespace *string `json:"namespace,omitempty"`
}

func (o ClusterInfoResponseRegistryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInfoResponseRegistryInfo struct{}"
	}

	return strings.Join([]string{"ClusterInfoResponseRegistryInfo", string(data)}, " ")
}
