package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobNodeVpcInfo 任务实例VPC信息体。
type JobNodeVpcInfo struct {

	// 任务实例所在虚拟私有云ID。
	VpcId string `json:"vpc_id"`

	// 任务实例所在子网ID。
	SubnetId string `json:"subnet_id"`

	// 指定创建任务实例IP地址，多个IP端口之间请用“,”英文逗号分隔，目前仅支持设置IPv4地址，获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找子网的网段，选择未被占用的IP 。 - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询私有IP列表，选择“device_owner”为空的私有IP。
	CustomNodeIp *string `json:"custom_node_ip,omitempty"`

	// 任务实例所在的安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty"`
}

func (o JobNodeVpcInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobNodeVpcInfo struct{}"
	}

	return strings.Join([]string{"JobNodeVpcInfo", string(data)}, " ")
}
