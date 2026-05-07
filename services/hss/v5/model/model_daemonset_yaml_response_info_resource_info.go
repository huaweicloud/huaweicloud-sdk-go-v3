package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DaemonsetYamlResponseInfoResourceInfo 资源配置信息
type DaemonsetYamlResponseInfoResourceInfo struct {

	// **参数解释**： 资源限制类型：默认规则or自定义or自适应 **取值范围**： - default：默认类型。 - customized：用户自定义类型。 - adaptive：自适应类型。
	Mode *string `json:"mode,omitempty"`

	// **参数解释**: cpu最大值 **取值范围**: 字符长度0-32位
	CpuLimit *string `json:"cpu_limit,omitempty"`

	// **参数解释**: 内存最大值 **取值范围**: 字符长度0-32位
	MemLimit *string `json:"mem_limit,omitempty"`
}

func (o DaemonsetYamlResponseInfoResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DaemonsetYamlResponseInfoResourceInfo struct{}"
	}

	return strings.Join([]string{"DaemonsetYamlResponseInfoResourceInfo", string(data)}, " ")
}
