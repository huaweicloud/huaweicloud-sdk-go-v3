package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2CreateCluster **参数解释**： v2创建集群请求。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type V2CreateCluster struct {

	// **参数解释**： 集群名称，要求唯一性。 **约束限制**： 不涉及。 **取值范围**： 要求唯一性，必须以字母开头并只包含字母、数字、中划线或下划线，长度为4~64个字符。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 集群规格编码。 **约束限制**： 不涉及。 **取值范围**： 请参见集群规格接口返回的规格编码。 **默认取值**： 不涉及。
	Flavor string `json:"flavor"`

	// **参数解释**： 集群CN数量。 **约束限制**： 不涉及。 **取值范围**： 取值范围为2~集群节点数，最大值为20，默认值为3。 **默认取值**： 不涉及。
	NumCn *int32 `json:"num_cn,omitempty"`

	// **参数解释**： 集群节点数量。 **约束限制**： 不涉及。 **取值范围**： 集群模式取值范围为3~256，实时数仓（单机模式）取值为1。 **默认取值**： 不涉及。
	NumNode int32 `json:"num_node"`

	// **参数解释**： 管理员用户名称。 **约束限制**： 不涉及。 **取值范围**： 只能由小写字母、数字或下划线组成。 必须由小写字母或下划线开头。 长度为1~63个字符。 用户名不能为DWS数据库的关键字。 **默认取值**： 不涉及。
	DbName string `json:"db_name"`

	// **参数解释**： 管理员用户密码。 **约束限制**： 不涉及。 **取值范围**： 12~32个字符； 至少包含以下字符中的3种：大写字母、小写字母、数字和特殊字符（~!?,.:;-_(){}[]/<>@#%^&*+|\\=）。 不能与用户名或倒序的用户名相同。 **默认取值**： 不涉及。
	DbPassword string `json:"db_password"`

	// **参数解释**： 集群数据库端口。 **约束限制**： 不涉及。 **取值范围**： 8000~30000 **默认取值**： 8000
	DbPort int32 `json:"db_port"`

	// **参数解释**： 专属存储池ID，一般不需要填写。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	DssPoolId *string `json:"dss_pool_id,omitempty"`

	// **参数解释**： 可用区信息，创建3AZ集群时需传入3个不同可用区。 **约束限制**： 不涉及。 **取值范围**： 获取方法请参见[查询可用区列表接口](ListAvailabilityZones.xml)。 **默认取值**： 不涉及。
	AvailabilityZones []string `json:"availability_zones"`

	// **参数解释**： 标签信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Tags *[]Tags `json:"tags,omitempty"`

	// **参数解释**： 指定虚拟私有云ID，用于集群网络配置。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	VpcId string `json:"vpc_id"`

	// **参数解释**： 指定子网ID，用于集群网络配置。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SubnetId string `json:"subnet_id"`

	// **参数解释**： 指定安全组ID，用于集群网络配置。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	PublicIp *PublicIp `json:"public_ip,omitempty"`

	// **参数解释**： 集群版本。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DatastoreVersion string `json:"datastore_version"`

	// **参数解释**： KMS密钥ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MasterKeyId *string `json:"master_key_id,omitempty"`

	// **参数解释**： KMS密钥名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MasterKeyName *string `json:"master_key_name,omitempty"`

	// **参数解释**： KMS加密算法。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CryptAlgorithm *string `json:"crypt_algorithm,omitempty"`

	Volume *Volume `json:"volume"`

	// **参数解释**： 企业项目ID，对集群指定企业项目。如果未指定，则使用默认企业项目“default”的ID，即0。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 指定网络协议类型，表明是否支持IPv6，默认不使用IPv6。使用ipv6时必须所选择的子网也支持ipv6。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`
}

func (o V2CreateCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2CreateCluster struct{}"
	}

	return strings.Join([]string{"V2CreateCluster", string(data)}, " ")
}
