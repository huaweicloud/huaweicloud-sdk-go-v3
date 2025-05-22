package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterCheckBody **参数解释**： 集群校验对象。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ClusterCheckBody struct {

	// **参数解释**： 企业项目ID，对集群指定企业项目。如果未指定，则使用默认企业项目“default”的ID，即0。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 集群规格名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Flavor string `json:"flavor"`

	// **参数解释**： 可用区列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AvailabilityZones []string `json:"availability_zones"`

	// **参数解释**： 实例节点个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NumNode int32 `json:"num_node"`

	// **参数解释**： 安全组ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// **参数解释**： 集群版本。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DatastoreVersion string `json:"datastore_version"`

	// **参数解释**： 集群虚拟私有云ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	VpcId string `json:"vpc_id"`

	// **参数解释**： 集群子网ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SubnetId string `json:"subnet_id"`

	PublicIp *OpenPublicIp `json:"public_ip,omitempty"`

	// **参数解释**： 跨规格恢复信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	CrossSpecRestore *string `json:"cross_spec_restore,omitempty"`

	Volume *Volume `json:"volume,omitempty"`

	// **参数解释**： 旧主机名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	OldClusterHostname *string `json:"old_cluster_hostname,omitempty"`

	RestorePoint *RestorePoint `json:"restore_point,omitempty"`

	// **参数解释**： 标签列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	TagList *[]Tag `json:"tag_list,omitempty"`

	// **参数解释**： 存储池ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	DssPoolId *string `json:"dss_pool_id,omitempty"`

	// **参数解释**： 数据库端口。 **约束限制**： 不涉及。 **取值范围**： 8000~30000 **默认取值**： 8000
	DbPort *string `json:"db_port,omitempty"`

	// **参数解释**： 管理员密码。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DbPassword *string `json:"db_password,omitempty"`

	// **参数解释**： 管理员用户。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： dbadmin
	DbName *string `json:"db_name,omitempty"`

	// **参数解释**： cn节点数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NumCn *int32 `json:"num_cn,omitempty"`

	// **参数解释**： 集群名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o ClusterCheckBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterCheckBody struct{}"
	}

	return strings.Join([]string{"ClusterCheckBody", string(data)}, " ")
}
