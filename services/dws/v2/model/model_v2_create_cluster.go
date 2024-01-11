package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2CreateCluster v2创建集群请求
type V2CreateCluster struct {

	// 集群名称，要求唯一性，必须以字母开头并只包含字母、数字、中划线或下划线，长度为4~64个字符。
	Name string `json:"name"`

	// 集群规格名称。节点类型详情请参见数据仓库类型数据仓库类型。
	Flavor string `json:"flavor"`

	// 集群CN数量，取值范围为2~集群节点数，最大值为20，默认值为3。
	NumCn int32 `json:"num_cn"`

	// 集群节点数量，集群模式取值范围为3~256，实时数仓（单机模式）取值为1。
	NumNode int32 `json:"num_node"`

	// 管理员用户名称。用户命名要求如下： 只能由小写字母、数字或下划线组成。 必须由小写字母或下划线开头。 长度为1~63个字符。用户名不能为DWS数据库的关键字。
	DbName string `json:"db_name"`

	// 管理员用户密码。 12~32个字符 至少包含以下字符中的3种：大写字母、小写字母、数字和特殊字符（~!?,.:;-_(){}[]/<>@#%^&*+|\\=）。不能与用户名或倒序的用户名相同。
	DbPassword string `json:"db_password"`

	// 集群数据库端口，取值范围为8000~30000，默认值：8000。
	DbPort int32 `json:"db_port"`

	// 专属存储池ID
	DssPoolId *string `json:"dss_pool_id,omitempty"`

	// 可用区列表。集群可用区选择详情请参见地区和终端节点地区和终端节点。
	AvailabilityZones []string `json:"availability_zones"`

	// 标签列表
	Tags *[]Tags `json:"tags,omitempty"`

	// 指定虚拟私有云ID，用于集群网络配置。
	VpcId string `json:"vpc_id"`

	// 指定子网ID，用于集群网络配置。
	SubnetId string `json:"subnet_id"`

	// 指定安全组ID，用于集群网络配置。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	PublicIp *PublicIp `json:"public_ip,omitempty"`

	// 集群版本
	DatastoreVersion string `json:"datastore_version"`

	// 密钥ID
	MasterKeyId *string `json:"master_key_id,omitempty"`

	// 密钥名称
	MasterKeyName *string `json:"master_key_name,omitempty"`

	// 加密算法
	CryptAlgorithm *string `json:"crypt_algorithm,omitempty"`

	Volume *Volume `json:"volume,omitempty"`

	// 企业项目ID，对集群指定企业项目，如果未指定，则使用默认企业项目“default”的ID，即0。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o V2CreateCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2CreateCluster struct{}"
	}

	return strings.Join([]string{"V2CreateCluster", string(data)}, " ")
}
