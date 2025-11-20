package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateConnectionOption struct {

	// - 参数解释：下联面网口主备IP；ESW实例在本端二层子网中占用的主备接口IP。 - 约束限制：字符串列表，只能设置两个字符串，且每个字符串内容应该是标准IPv4格式；IP必须在二层子网网段范围内。 - 取值范围：不涉及。 - 默认取值：不涉及。
	FixedIps *[]string `json:"fixed_ips,omitempty"`

	// - 参数解释：二层连接的名称。 - 约束限制：   - 长度范围为1~64个字符。   - 名称由中文、英文字母、数字、下划线（_）、中划线（-）、点（.）组成。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Name string `json:"name"`

	// - 参数解释：远端隧道信息。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	RemoteInfos []RemoteInfosOption `json:"remote_infos"`

	// - 参数解释：二层连接关联的二层子网ID。 - 约束限制：   - 需要使用本租户下可操作的子网资源的ID；此值即为子网详情中的“网络ID”参数值。   - 带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	VirsubnetId string `json:"virsubnet_id"`

	// - 参数解释：ESW所在VPC资源ID。 - 约束限制：   - 需要使用本租户下可操作的VPC资源的ID。   - 带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	VpcId string `json:"vpc_id"`
}

func (o CreateConnectionOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionOption struct{}"
	}

	return strings.Join([]string{"CreateConnectionOption", string(data)}, " ")
}
