package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例节点信息。
type ListInstancesNodeResult struct {

	// 节点ID。
	Id string `json:"id" xml:"id"`

	// 节点名称。
	Name string `json:"name" xml:"name"`

	// 节点状态。
	Status string `json:"status" xml:"status"`

	// 节点角色。 该参数仅对GaussDB(for Mongo)引擎的副本集实例有效。
	Role string `json:"role" xml:"role"`

	// 节点所在的子网的ID。
	SubnetId string `json:"subnet_id" xml:"subnet_id"`

	// 节点内网IP。在弹性云服务器创建成功后参数值存在，否则，值为\"\"。
	PrivateIp string `json:"private_ip" xml:"private_ip"`

	// 绑定的公网IP。该参数仅针对绑定了公网IP的节点有效。
	PublicIp string `json:"public_ip" xml:"public_ip"`

	// 资源规格编码。关于实例的规格信息，请参见查询所有实例规格信息中响应参数“flavors.spec_code”的值。
	SpecCode string `json:"spec_code" xml:"spec_code"`

	// 可用区。
	AvailabilityZone string `json:"availability_zone" xml:"availability_zone"`

	// 是否支持节点缩容。 - true，表示该节点支持节点缩容。 - false，表示该节点不支持节点缩容。
	SupportReduce bool `json:"support_reduce" xml:"support_reduce"`
}

func (o ListInstancesNodeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesNodeResult struct{}"
	}

	return strings.Join([]string{"ListInstancesNodeResult", string(data)}, " ")
}
