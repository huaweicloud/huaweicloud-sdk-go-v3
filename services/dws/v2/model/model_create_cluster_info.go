package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterInfo **参数解释**： 创建集群对象。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type CreateClusterInfo struct {

	// **参数解释**： 节点规格名称 **取值范围**： 集群规格详情请参见[数据仓库规格](dws_01_00018.html)。
	NodeType string `json:"node_type"`

	// **参数解释**： 节点数量。 **约束限制**： 不涉及。 **取值范围**： 集群模式取值范围为3~256。 **默认取值**： 不涉及。
	NumberOfNode int32 `json:"number_of_node"`

	// **参数解释**： 指定子网ID，用于集群网络配置。 **约束限制**： 不涉及。 **取值范围**： 必须是虚拟私有云ID下的某个子网。 **默认取值**： 不涉及。
	SubnetId string `json:"subnet_id"`

	// **参数解释**： 指定安全组ID，用于集群网络配置。 **约束限制**： 不涉及。 **取值范围**： 参数非空时必须是有效的安全组ID。参数为空时将自动创建安全组。 **默认取值**： null
	SecurityGroupId string `json:"security_group_id"`

	// **参数解释**： 指定虚拟私有云ID，用于集群网络配置。 **约束限制**： 不涉及。 **取值范围**： 必须是有效的虚拟私有云ID。 **默认取值**： 不涉及。
	VpcId string `json:"vpc_id"`

	// **参数解释**： 配置集群可用区。 **约束限制**： 不涉及。 **取值范围**： 必须是当前局点下状态有效且当前用户可见的可用区编码。 **默认取值**： 查询可用区时第一个可用的可用区编码。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释**： 集群数据库端口。 **约束限制**： 不涉及。 **取值范围**： 8000~30000 **默认取值**： 8000
	Port *int32 `json:"port,omitempty"`

	// **参数解释**： 集群名称。 **约束限制**： 要求唯一性，必须以字母开头并只包含字母、数字、中划线或下划线，长度为4~64个字符。 **取值范围**： 4~64个字符。 **默认取值**： 8000
	Name string `json:"name"`

	// **参数解释**： DWS集群管理员用户名。 **约束限制**： - 只能由小写字母、数字或下划线组成。 - 必须由小写字母或下划线开头。 - 长度为1~63个字符。 - 用户名不能为DWS数据库的关键字。    **取值范围**：   1~63个字符； **默认取值**： dbadmin
	UserName string `json:"user_name"`

	// **参数解释**： DWS集群管理员密码。 **约束限制**： 不涉及。 **取值范围**： 12~32个字符； 至少包含以下字符的3种：大写字母、小写字母、数字和特殊字符(~!?,.:;_(){}[]/<>@#%^&*+|\\\\=-)； 不能与用户名或倒序的用户名相同； **默认取值**： 不涉及。
	UserPwd string `json:"user_pwd"`

	PublicIp *PublicIp `json:"public_ip,omitempty"`

	// **参数解释**： CN部署量。 **约束限制**： 不涉及。 **取值范围**： 2~集群节点数，最大值为20。 **默认取值**： 默认值为3。
	NumberOfCn *int32 `json:"number_of_cn,omitempty"`

	// **参数解释**： 企业项目ID，对集群指定企业项目。如果未指定，则使用默认企业项目“default”的ID，即0。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *Tags `json:"tags,omitempty"`
}

func (o CreateClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterInfo struct{}"
	}

	return strings.Join([]string{"CreateClusterInfo", string(data)}, " ")
}
