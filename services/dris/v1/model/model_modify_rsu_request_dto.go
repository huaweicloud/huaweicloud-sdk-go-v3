package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyRsuRequestDto struct {

	// **参数说明**：RSU的名字。  **取值范围**：长度不低于1不超过128，只允许中文、字母、数字、下划线（_）、连接符（-）的组合。
	Name *string `json:"name,omitempty"`

	// **参数说明**：RSU的描述。  **取值范围**：只允许中文、字母、数字、下划线（_）、中文分号（；）、中文冒号（：）、中文问号（？）、中文感叹号（！）中文逗号（，）、中文句号（。）、英文分号（;）、英文冒号（:）、英文逗号（,）、英文句号（.）、英文问号（?）、英文感叹号（!）、顿号（、）、连接符（-）的组合。
	Description *string `json:"description,omitempty"`

	// **参数说明**：RSU的IP。满足IP的格式，例如127.0.0.1。
	Ip *string `json:"ip,omitempty"`

	// **参数说明**：安装位置编码，由用户自定义。  **取值范围**：长度不低于1不超过128，只允许字母、数字、下划线（_）的组合。
	PositionDescription *string `json:"position_description,omitempty"`

	// **参数说明**：RSU可关联的Edge的数量，修改值需大于等于当前已连接设备。
	RelatedEdgeNum *int32 `json:"related_edge_num,omitempty"`

	// **参数说明**：在地图中，rsu所在区域对应的路口ID，也即区域ID拼接路口ID，格式为：region-node_id。其中路网最基本的构成即节点和节点之间连接的路段。节点可以是路口，也可以是一条 路的端点。一个节点的ID在同一个区域内是唯一的。
	IntersectionId *string `json:"intersection_id,omitempty"`
}

func (o ModifyRsuRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyRsuRequestDto struct{}"
	}

	return strings.Join([]string{"ModifyRsuRequestDto", string(data)}, " ")
}
