package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeEipSpec
type NodeEipSpec struct {

	// **参数解释**： 弹性IP类型，取值请参见申请EIP接口中publicip.type说明。 [链接请参见[申请EIP](https://support.huaweicloud.com/api-eip/eip_api_0001.html)。](tag:hws) [链接请参见[申请EIP](https://support.huaweicloud.com/intl/zh-cn/api-eip/eip_api_0001.html)。](tag:hws_hk) **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Iptype string `json:"iptype"`

	Bandwidth *NodeBandwidth `json:"bandwidth,omitempty"`
}

func (o NodeEipSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeEipSpec struct{}"
	}

	return strings.Join([]string{"NodeEipSpec", string(data)}, " ")
}
