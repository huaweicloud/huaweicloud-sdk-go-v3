package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EniNetwork ENI网络配置，创建集群指定使用云原生网络2.0网络模式时必填subnets和eniSubnetId其中一个字段(eniSubnetCIDR可选，若填写了会校验是否合法)，1.19.10及新版本集群使用subnets字段，1.19.8及老版本若使用subnets字段，则取值subnets数组中的第一个子网ID作为容器地址使用的子网ID。
type EniNetwork struct {

	// **参数解释：** ENI所在子网的IPv4子网ID。 **约束限制：** 暂不支持IPv6。该字段将会被废弃，推荐使用新字段subnets。 **取值范围：** 不涉及 **默认取值：** 不涉及  获取方法如下：  - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找IPv4子网ID。 - 方法2：通过虚拟私有云服务的查询子网列表接口查询。   [链接请参见[查询子网列表](https://support.huaweicloud.com/api-vpc/vpc_subnet01_0003.html)。](tag:hws)   [链接请参见[查询子网列表](https://support.huaweicloud.com/intl/zh-cn/api-vpc/vpc_subnet01_0003.html)。](tag:hws_hk)
	EniSubnetId *string `json:"eniSubnetId,omitempty"`

	// ENI子网CIDR(废弃中)
	EniSubnetCIDR *string `json:"eniSubnetCIDR,omitempty"`

	// IPv4子网ID列表
	Subnets []NetworkSubnet `json:"subnets"`
}

func (o EniNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EniNetwork struct{}"
	}

	return strings.Join([]string{"EniNetwork", string(data)}, " ")
}
