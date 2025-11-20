package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TunnelInfoResult struct {

	// - 参数解释：ESW所在VPC资源ID。 - 约束限制：   - 需要使用本租户下可操作的VPC资源的ID。   - 带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	VpcId string `json:"vpc_id"`

	// - 参数解释：ESW所在隧道子网ID。 - 约束限制：   - 需要使用本租户下可操作的子网资源的ID；此值即为子网详情中的“网络ID”参数值。   - 带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	VirsubnetId string `json:"virsubnet_id"`

	// - 参数解释：ESW实例的本端隧道IP。 - 约束限制：不能与已存在的子网IP冲突。 - 取值范围：标准的IPv4格式，例：192.168.1.1。 - 默认取值：不涉及。
	TunnelIp string `json:"tunnel_ip"`

	// - 参数解释：ESW实例的隧道端口。 - 约束限制：不涉及。 - 取值范围：4789。 - 默认取值：不涉及。
	TunnelPort int32 `json:"tunnel_port"`

	// - 参数解释：ESW实例的隧道协议类型。 - 约束限制：不涉及。 - 取值范围：vxlan。 - 默认取值：不涉及。
	TunnelType string `json:"tunnel_type"`
}

func (o TunnelInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TunnelInfoResult struct{}"
	}

	return strings.Join([]string{"TunnelInfoResult", string(data)}, " ")
}
