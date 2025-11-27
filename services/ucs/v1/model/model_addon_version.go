package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddonVersion 插件版本相关信息
type AddonVersion struct {

	// 插件包版本id
	Id *string `json:"id,omitempty"`

	// 插件版本信息
	Version *string `json:"version,omitempty"`

	// 输入
	Input map[string]interface{} `json:"input,omitempty"`

	// 是否为稳定版本
	Stable *bool `json:"stable,omitempty"`

	// 供界面使用的翻译信息
	Translate map[string]interface{} `json:"translate,omitempty"`

	// 支持的集群类型和和支持的集群版本信息
	SupportVersions *[]SupportVersion `json:"supportVersions,omitempty"`

	// 记录创建时间
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// 记录更新时间
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`
}

func (o AddonVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonVersion struct{}"
	}

	return strings.Join([]string{"AddonVersion", string(data)}, " ")
}
