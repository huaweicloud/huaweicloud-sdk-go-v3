package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodesRequest Request Object
type ListNodesRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`

	// 集群状态兼容Error参数，用于API平滑切换。 兼容场景下，errorStatus为空则屏蔽Error状态为Deleting状态。
	ErrorStatus *string `json:"errorStatus,omitempty"`

	// **参数解释**： 设置每页显示的数据条数。 **约束限制**： 不涉及 **取值范围**： 1到2000之间（含1和2000）的整数。 **默认取值**： 2000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 通过资源uid进行分页查询,默认为查询第一页数据。marker={{uid}}表示查询该uid后的资源列表的信息(查询结果不包含该uid的资源)。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 无
	Marker *string `json:"marker,omitempty"`
}

func (o ListNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodesRequest struct{}"
	}

	return strings.Join([]string{"ListNodesRequest", string(data)}, " ")
}
