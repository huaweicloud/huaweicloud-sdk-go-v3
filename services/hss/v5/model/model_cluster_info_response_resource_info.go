package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterInfoResponseResourceInfo 资源配置信息
type ClusterInfoResponseResourceInfo struct {

	// **参数解释** 资源限制类型：默认规则or自定义or自适应 **取值范围** 取值0-32
	Mode *string `json:"mode,omitempty"`

	// **参数解释** cpu最大值 **取值范围** 字符长度0-32位
	CpuLimit *string `json:"cpu_limit,omitempty"`

	// **参数解释** 内存最大值 **取值范围** 字符长度0-32位
	MemLimit *string `json:"mem_limit,omitempty"`
}

func (o ClusterInfoResponseResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInfoResponseResourceInfo struct{}"
	}

	return strings.Join([]string{"ClusterInfoResponseResourceInfo", string(data)}, " ")
}
