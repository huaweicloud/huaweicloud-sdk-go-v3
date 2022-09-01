package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例配置信息
type InstanceConfigResult struct {

	// 云服务器的规格ID。
	FlavorRef *string `json:"flavorRef,omitempty" xml:"flavorRef"`

	// 镜像ID，同image_id。
	ImageRef *string `json:"imageRef,omitempty" xml:"imageRef"`

	// 磁盘组信息。
	Disk *[]DiskResult `json:"disk,omitempty" xml:"disk"`

	// 登录云服务器的SSH密钥名称。
	KeyName *string `json:"key_name,omitempty" xml:"key_name"`

	// 登录云服务器的SSH密钥指纹。
	KeyFingerprint *string `json:"key_fingerprint,omitempty" xml:"key_fingerprint"`

	// 该参数为预留字段。
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name"`

	// 该参数为预留字段。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 登录云服务器的密码，非明文回显。
	AdminPass *string `json:"adminPass,omitempty" xml:"adminPass"`

	// 个人信息
	Personality *[]PersonalityResult `json:"personality,omitempty" xml:"personality"`

	PublicIp *PublicipResult `json:"public_ip,omitempty" xml:"public_ip"`

	// cloud-init用户数据，base64格式编码。
	UserData *string `json:"user_data,omitempty" xml:"user_data"`

	Metadata *VmMetaData `json:"metadata,omitempty" xml:"metadata"`

	// 安全组信息。
	SecurityGroups *[]SecurityGroups `json:"security_groups,omitempty" xml:"security_groups"`

	// 云服务器组ID。
	ServerGroupId *string `json:"server_group_id,omitempty" xml:"server_group_id"`

	// 在专属主机上创建弹性云服务器。
	Tenancy *string `json:"tenancy,omitempty" xml:"tenancy"`

	// 专属主机的ID。
	DedicatedHostId *string `json:"dedicated_host_id,omitempty" xml:"dedicated_host_id"`

	// 云服务器的计费模式，可以选择竞价计费或按需计费。
	MarketType *string `json:"market_type,omitempty" xml:"market_type"`

	// 使用伸缩配置创建云主机的时候，多规格使用的优先级策略。  PICK_FIRST（默认）：选择优先，虚拟机扩容时规格的选择按照flavorRef列表的顺序进行优先级排序。 COST_FIRST：成本优化，虚拟机扩容时规格的选择按照价格最优原则进行优先级排序。
	MultiFlavorPriorityPolicy *string `json:"multi_flavor_priority_policy,omitempty" xml:"multi_flavor_priority_policy"`
}

func (o InstanceConfigResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceConfigResult struct{}"
	}

	return strings.Join([]string{"InstanceConfigResult", string(data)}, " ")
}
