package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群对象
type CreateClusterInfo struct {

	// 节点类型
	NodeType string `json:"node_type"`

	// 集群节点数量，取值范围为3~32。
	NumberOfNode int32 `json:"number_of_node"`

	// 指定子网ID，用于集群网络配置。
	SubnetId string `json:"subnet_id"`

	// 指定安全组ID，用于集群网络配置。
	SecurityGroupId string `json:"security_group_id"`

	// 指定虚拟私有云ID，用于集群网络配置。
	VpcId string `json:"vpc_id"`

	// 配置集群可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 集群服务端口，取值范围为8000~30000，默认值：8000。
	Port *int32 `json:"port,omitempty"`

	// 集群名称，要求唯一性，必须以字母开头并只包含字母、数字、中划线或下划线，长度为4~64个字符。
	Name string `json:"name"`

	// DWS集群管理员用户名。用户命名要求如下：  - 只能由小写字母、数字或下划线组成。 - 必须由小写字母或下划线开头。 - 长度为1~63个字符。 - 用户名不能为DWS数据库的关键字。
	UserName string `json:"user_name"`

	// DWS集群管理员密码。
	UserPwd string `json:"user_pwd"`

	PublicIp *PublicIp `json:"public_ip,omitempty"`

	// CN部署量，取值范围为2~集群节点数-1，最大值为20，默认值为3。
	NumberOfCn *int32 `json:"number_of_cn,omitempty"`

	// 企业项目ID，对集群指定企业项目，如果未指定，则使用默认企业项目“default”的ID，即0。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterInfo struct{}"
	}

	return strings.Join([]string{"CreateClusterInfo", string(data)}, " ")
}
