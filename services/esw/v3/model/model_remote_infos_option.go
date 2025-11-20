package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoteInfosOption struct {

	// - 参数解释：二层连接的隧道号，对应VXLAN网络标识VNI。 - 约束限制：需与对端VXLAN设置的VNI保持一致。 - 取值范围：1-16777216。 - 默认取值：不涉及。
	SegmentationId int32 `json:"segmentation_id"`

	// - 参数解释：ESW实例的远端隧道IP。 - 约束限制：需与对端VXLAN设置的VTEP IP保持一致。 - 取值范围：标准的IPv4格式，例：192.168.1.1。 - 默认取值：不涉及。
	TunnelIp string `json:"tunnel_ip"`

	// - 参数解释：二层连接的远端隧道端口。 - 约束限制：不涉及。 - 取值范围：4789。 - 默认取值：不涉及。
	TunnelPort *int32 `json:"tunnel_port,omitempty"`
}

func (o RemoteInfosOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoteInfosOption struct{}"
	}

	return strings.Join([]string{"RemoteInfosOption", string(data)}, " ")
}
