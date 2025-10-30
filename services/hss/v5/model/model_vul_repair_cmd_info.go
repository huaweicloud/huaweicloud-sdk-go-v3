package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulRepairCmdInfo **参数解释**: 漏洞修复命令信息 **取值范围**: 不涉及
type VulRepairCmdInfo struct {

	// **参数解释**: 主机ID **取值范围**: 字符长度1-128
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 主机名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器公网IP **取值范围**: 字符长度0-128
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 服务器私网IP **取值范围**: 字符长度0-128
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 主机的资产重要性 **取值范围**: - important：重要资产 - common：一般资产 - test：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**： 漏洞名称 **取值范围**： 字符长度0-256位
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**: 漏洞ID **取值范围**: 字符长度0-64位
	VulId *string `json:"vul_id,omitempty"`

	// **参数解释**: 修复命令行 **取值范围**: 字符范围1-256位
	RepairCmd *string `json:"repair_cmd,omitempty"`

	// **参数解释**: 关联镜像名称 **取值范围**: 字符范围1-256位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 关联镜像ID **取值范围**: 字符范围1-256位
	ImageId *string `json:"image_id,omitempty"`
}

func (o VulRepairCmdInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulRepairCmdInfo struct{}"
	}

	return strings.Join([]string{"VulRepairCmdInfo", string(data)}, " ")
}
