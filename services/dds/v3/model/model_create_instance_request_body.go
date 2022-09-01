package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例信息。
type CreateInstanceRequestBody struct {

	// 实例名称。用于表示实例的名称，用于表示实例的名称，允许和已有名称重复。 取值范围：长度为4~64位，必须以字母开头（A~Z或a~z），区分大小写，可以包含字母、数字（0~9）、中划线（-）或者下划线（_），不能包含其他特殊字符。
	Name string `json:"name" xml:"name"`

	Datastore *Datastore `json:"datastore" xml:"datastore"`

	// - 区域ID - 取值：非空。
	Region string `json:"region" xml:"region"`

	// 可用区ID。非专属云用户可以选择多个AZ，创建跨AZ的集群。专属云用户暂不支持创建跨AZ的集群。取值：非空，请参见[地区和终端节点](https://developer.huaweicloud.com/endpoint)。
	AvailabilityZone string `json:"availability_zone" xml:"availability_zone"`

	// 虚拟私有云ID。获取方法请参见《虚拟私有云API参考》中“VPC”的内容。 取值：非空，字符长度校验，严格UUID正则校验。
	VpcId string `json:"vpc_id" xml:"vpc_id"`

	// 子网ID。获取方法请参见《虚拟私有云API参考》中“子网”的内容。
	SubnetId string `json:"subnet_id" xml:"subnet_id"`

	// 指定实例所属的安全组ID。 获取方法请参见《虚拟私有云API参考》中“安全组”的内容。
	SecurityGroupId string `json:"security_group_id" xml:"security_group_id"`

	// 数据库访问端口。 取值范围：2100~9500，以及27017、27018、27019。 不传该参数时，创建实例的访问端口默认为8635。
	Port *string `json:"port,omitempty" xml:"port"`

	// 数据库密码。 取值范围：长度为8~32位，必须是大写字母（A~Z）、小写字母（a~z）、数字（0~9）、特殊字符~!@#%^*-_=+?的组合。 建议您输入高强度密码，以提高安全性，防止出现密码被暴力破解等安全风险。
	Password *string `json:"password,omitempty" xml:"password"`

	// 磁盘加密时的密钥ID，严格UUID正则校验。 不传该参数时，表示不进行磁盘加密。
	DiskEncryptionId *string `json:"disk_encryption_id,omitempty" xml:"disk_encryption_id"`

	// 实例类型。支持集群、副本集、以及单节点。 取值   - Sharding   - ReplicaSet   - Single
	Mode string `json:"mode" xml:"mode"`

	// 参数组配置信息。
	Configurations *[]CreateInstanceConfigurationsOption `json:"configurations,omitempty" xml:"configurations"`

	// 实例规格详情。
	Flavor []CreateInstanceFlavorOption `json:"flavor" xml:"flavor"`

	BackupStrategy *BackupStrategy `json:"backup_strategy,omitempty" xml:"backup_strategy"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// SSL开关选项。 取值： - 取“0”，表示DDS实例默认不启用SSL连接。 - 取“1”，表示DDS实例默认启用SSL连接。 - 不传该参数时，默认启用SSL连接。
	SslOption *string `json:"ssl_option,omitempty" xml:"ssl_option"`

	// Dec用户专属存储ID，默认为空。仅Dec用户支持该参数。
	DssPoolId *string `json:"dss_pool_id,omitempty" xml:"dss_pool_id"`

	// 创建新实例设置云服务器组关联的策略名称列表，仅专属云创建实例时有效。 取值    - 取“anti-affinity”，表示DDS实例开启反亲和部署，反亲和部署是出于高可用性考虑，将您的Primary、Secondary和Hidden节点分别创建在不同的物理机上。当前仅支持该值，不传该值默认不开启反亲和部署。
	ServerGroupPolicies *[]string `json:"server_group_policies,omitempty" xml:"server_group_policies"`

	// 标签列表。单个实例总标签数上限20个。
	Tags *[]TagWithKeyValue `json:"tags,omitempty" xml:"tags"`

	ChargeInfo *ChargeInfoOption `json:"charge_info,omitempty" xml:"charge_info"`
}

func (o CreateInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequestBody", string(data)}, " ")
}
