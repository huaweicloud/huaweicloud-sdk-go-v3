package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// absolute数据结构说明
type Absolute struct {

	// 裸金属服务器最大申请数量
	MaxTotalInstances *int32 `json:"maxTotalInstances,omitempty" xml:"maxTotalInstances"`

	// CPU核数最大申请数量
	MaxTotalCores *int32 `json:"maxTotalCores,omitempty" xml:"maxTotalCores"`

	// 内存最大申请容量（单位：MB）
	MaxTotalRAMSize *int32 `json:"maxTotalRAMSize,omitempty" xml:"maxTotalRAMSize"`

	// 可以申请的SSH密钥对最大数量
	MaxTotalKeypairs *int32 `json:"maxTotalKeypairs,omitempty" xml:"maxTotalKeypairs"`

	// 可输入元数据的最大长度
	MaxServerMeta *int32 `json:"maxServerMeta,omitempty" xml:"maxServerMeta"`

	// 可注入文件的最大个数
	MaxPersonality *int32 `json:"maxPersonality,omitempty" xml:"maxPersonality"`

	// 注入文件内容的最大长度（单位：Byte）
	MaxPersonalitySize *int32 `json:"maxPersonalitySize,omitempty" xml:"maxPersonalitySize"`

	// 服务器组的最大个数
	MaxServerGroups *int32 `json:"maxServerGroups,omitempty" xml:"maxServerGroups"`

	// 服务器组中的最大裸金属服务器数。
	MaxServerGroupMembers *int32 `json:"maxServerGroupMembers,omitempty" xml:"maxServerGroupMembers"`

	// 已使用的服务器组个数
	TotalServerGroupsUsed *int32 `json:"totalServerGroupsUsed,omitempty" xml:"totalServerGroupsUsed"`

	// 安全组最大使用个数。 说明：具体配额限制请以VPC配额限制为准。
	MaxSecurityGroups *int32 `json:"maxSecurityGroups,omitempty" xml:"maxSecurityGroups"`

	// 安全组中安全组规则最大的配置个数。 说明：具体配额限制请以VPC配额限制为准。
	MaxSecurityGroupRules *int32 `json:"maxSecurityGroupRules,omitempty" xml:"maxSecurityGroupRules"`

	// 最大的浮动IP使用个数
	MaxTotalFloatingIps *int32 `json:"maxTotalFloatingIps,omitempty" xml:"maxTotalFloatingIps"`

	// 镜像元数据最大的长度
	MaxImageMeta *int32 `json:"maxImageMeta,omitempty" xml:"maxImageMeta"`

	// 当前裸金属服务器使用个数
	TotalInstancesUsed *int32 `json:"totalInstancesUsed,omitempty" xml:"totalInstancesUsed"`

	// 当前已使用CPU核数
	TotalCoresUsed *int32 `json:"totalCoresUsed,omitempty" xml:"totalCoresUsed"`

	// 当前内存使用容量（单位：MB）
	TotalRAMUsed *int32 `json:"totalRAMUsed,omitempty" xml:"totalRAMUsed"`

	// 当前安全组使用个数
	TotalSecurityGroupsUsed *int32 `json:"totalSecurityGroupsUsed,omitempty" xml:"totalSecurityGroupsUsed"`

	// 当前浮动IP使用个数
	TotalFloatingIpsUsed *int32 `json:"totalFloatingIpsUsed,omitempty" xml:"totalFloatingIpsUsed"`
}

func (o Absolute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Absolute struct{}"
	}

	return strings.Join([]string{"Absolute", string(data)}, " ")
}
