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

	// - 参数解释：分页查询每页返回的记录个数。 - 约束限制：limit需要和marker配合使用，详细规则请见marker的参数说明。 - 取值范围：0~2^31-1。 - 默认取值：2000。
	Limit *int32 `json:"limit,omitempty"`

	// - 参数解释：分页查询的起始资源ID，表示从指定资源的下一条记录开始查询。 - 约束限制：marker需要和limit配合使用。   - 若不传入marker和limit参数，查询结果返回第一页全部资源记录。   - 若不传入marker参数，limit为10，查询结果返回第1~10条资源记录。   - 若marker为第10条记录的资源ID，limit为10，查询结果返回第11~20条资源记录。   - 若marker为第10条记录的资源ID，不传入limit参数，查询结果返回第11~2000条（limit默认值2000）资源记录。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Marker *string `json:"marker,omitempty"`
}

func (o ListConnectionsAllInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsAllInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionsAllInstancesRequest", string(data)}, " ")
}
