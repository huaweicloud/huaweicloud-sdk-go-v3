package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodePublicIp
type NodePublicIp struct {

	// **参数解释**： 已有的弹性IP的ID列表。数量不得大于待创建节点数 > 若已配置ids参数，则无需配置count和eip参数 **约束限制**： 该参数仅给节点使用
	Ids *[]string `json:"ids,omitempty"`

	// **参数解释**： 要动态创建的弹性IP个数。 > count参数与eip参数必须同时配置。 **约束限制**： 该参数仅给节点使用
	Count *int32 `json:"count,omitempty"`

	Eip *NodeEipSpec `json:"eip,omitempty"`

	// **参数解释**： 表示节点所属分区。分区可以选择中心云[或者[边缘小站](https://support.huaweicloud.com/usermanual-cloudpond/ies_02_0401.html)。](tag:hws)[或者[边缘小站](https://support.huaweicloud.com/intl/zh-cn/usermanual-cloudpond/ies_02_0401.html)。](tag:hws_hk) 仅开启了对分布式云支持的Turbo集群支持指定该字段。 弹性IP类型，取值请参见申请EIP接口中publicip.type说明。 [链接请参见[申请EIP](https://support.huaweicloud.com/api-eip/eip_api_0001.html)](tag:hws) [链接请参见[申请EIP](https://support.huaweicloud.com/intl/zh-cn/api-eip/eip_api_0001.html)](tag:hws_hk) **约束限制**： 该参数仅给节点池使用
	Iptype *string `json:"iptype,omitempty"`

	Bandwidth *NodeBandwidth `json:"bandwidth,omitempty"`
}

func (o NodePublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePublicIp struct{}"
	}

	return strings.Join([]string{"NodePublicIp", string(data)}, " ")
}
