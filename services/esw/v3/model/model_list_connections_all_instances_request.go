package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionsAllInstancesRequest Request Object
type ListConnectionsAllInstancesRequest struct {

	// - 参数解释：二层连接资源ID。二层连接创建成功后，会生成一个ID，是二层连接对应的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Id *string `json:"id,omitempty"`

	// - 参数解释：二层连接的名称。 - 约束限制：   - 长度范围为1~64个字符。   - 名称由中文、英文字母、数字、下划线（_）、中划线（-）、点（.）组成。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Name *string `json:"name,omitempty"`

	// - 参数解释：ESW资源ID。ESW创建成功后，会生成一个ESW ID，是ESW对应的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// - 参数解释：ESW所在VPC资源ID。 - 约束限制：   - 需要使用本租户下可操作的VPC资源的ID。   - 带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// - 参数解释：二层连接关联的二层子网ID。 - 约束限制：   - 需要使用本租户下可操作的子网资源的ID；此值即为子网详情中的“网络ID”参数值。   - 带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	VirsubnetId *string `json:"virsubnet_id,omitempty"`
}

func (o ListConnectionsAllInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsAllInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionsAllInstancesRequest", string(data)}, " ")
}
