package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageVulInfo struct {

	// **参数解释**: 漏洞ID **取值范围**: 字符长度0-128
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 危险程度 **取值范围**: - immediate_repair：高危 - delay_repair：中危 - not_needed_repair：低危
	RepairNecessity *string `json:"repair_necessity,omitempty"`

	// **参数解释**: 漏洞描述 **取值范围** : 字符长度0-128
	Description *string `json:"description,omitempty"`

	// **参数解释**: 漏洞所在镜像层 **取值范围** : 字符长度0-128
	Position *string `json:"position,omitempty"`

	// **参数解释**: 漏洞的软件名称 **取值范围** : 字符长度0-128
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 应用软件的路径（只有应用漏洞有该字段） **取值范围** : 字符长度1-512
	AppPath *string `json:"app_path,omitempty"`

	// **参数解释**: 软件版本 **取值范围** : 字符长度0-128
	Version *string `json:"version,omitempty"`

	// **参数解释**: 解决方案 **取值范围** : 字符长度0-256
	Solution *string `json:"solution,omitempty"`

	// **参数解释**: 补丁地址 **取值范围** : 字符长度0-128
	Url *string `json:"url,omitempty"`
}

func (o ImageVulInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageVulInfo struct{}"
	}

	return strings.Join([]string{"ImageVulInfo", string(data)}, " ")
}
