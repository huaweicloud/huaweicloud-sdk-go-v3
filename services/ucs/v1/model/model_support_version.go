package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupportVersion 支持的集群类型（BareMetal，VirtualMachine，windows等）和支持的集群版本（支持正则表达式，如\".*\"匹配所有集群版本 ）信息等
type SupportVersion struct {

	// 支持的集群类型
	Category *string `json:"category,omitempty"`

	// 支持的集群类型（BareMetal，VirtualMachine，windows等）
	ClusterType *string `json:"clusterType,omitempty"`

	// 支持的集群版本，支持正则表达式，如\".*\"匹配所有集群版本
	ClusterVersion *[]string `json:"clusterVersion,omitempty"`
}

func (o SupportVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportVersion struct{}"
	}

	return strings.Join([]string{"SupportVersion", string(data)}, " ")
}
