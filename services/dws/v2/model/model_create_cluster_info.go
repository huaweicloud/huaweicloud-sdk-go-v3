package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群对象
type CreateClusterInfo struct {
	// 配置集群可用区

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	PublicIp *PublicIp `json:"public_ip,omitempty"`
	// 集群节点数量，取值范围为3~32。

	NumberOfNode int32 `json:"number_of_node"`
	// 指定虚拟私有云ID，用于集群网络配置。

	VpcId string `json:"vpc_id"`
	// DWS集群管理员用户名。

	UserName string `json:"user_name"`
	// 指定安全组ID，用于集群网络配置。

	SecurityGroupId string `json:"security_group_id"`
	// CN部署量，取值范围为2~集群节点数-1，最大值为5，默认值为2。

	NumberOfCn *int32 `json:"number_of_cn,omitempty"`
	// DWS集群管理员密码。

	UserPwd string `json:"user_pwd"`
	// 企业项目ID，对集群指定企业项目，如果未指定，则使用默认企业项目“default”的ID，即0。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 节点类型

	NodeType string `json:"node_type"`
	// 集群服务端口，取值范围为8000~30000，默认值：8000。

	Port *int32 `json:"port,omitempty"`
	// 集群名称，要求唯一性，必须以字母开头并只包含字母、数字、中划线或下划线，长度为4~64个字符。

	Name string `json:"name"`
	// 指定子网ID，用于集群网络配置。

	SubnetId string `json:"subnet_id"`
}

func (o CreateClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterInfo struct{}"
	}

	return strings.Join([]string{"CreateClusterInfo", string(data)}, " ")
}
