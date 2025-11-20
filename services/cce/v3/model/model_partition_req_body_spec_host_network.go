package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PartitionReqBodySpecHostNetwork **参数解释**： 分区子网 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type PartitionReqBodySpecHostNetwork struct {

	// **参数解释**： 子网ID 获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找IPv4子网ID。 - 方法2：通过虚拟私有云服务的查询子网列表接口查询。   [链接请参见[查询子网列表](https://support.huaweicloud.com/api-vpc/vpc_subnet01_0003.html)。](tag:hws)   [链接请参见[查询子网列表](https://support.huaweicloud.com/intl/zh-cn/api-vpc/vpc_subnet01_0003.html)。](tag:hws_hk) **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SubnetID *string `json:"subnetID,omitempty"`
}

func (o PartitionReqBodySpecHostNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionReqBodySpecHostNetwork struct{}"
	}

	return strings.Join([]string{"PartitionReqBodySpecHostNetwork", string(data)}, " ")
}
