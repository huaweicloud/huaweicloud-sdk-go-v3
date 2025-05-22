package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TopologyInfo 拓扑详情
type TopologyInfo struct {

	// **参数解释**： 节点ID。 **取值范围**： 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**： 节点名称。 **取值范围**： 不涉及。
	NodeName *string `json:"node_name,omitempty"`

	// **参数解释**： 节点IP。 **取值范围**： 不涉及。
	Ip *string `json:"ip,omitempty"`

	// **参数解释**： 节点端口。 **取值范围**： 不涉及。
	Port *string `json:"port,omitempty"`

	// **参数解释**： 节点主从角色。 **取值范围**： master：主节点 slave：从节点 proxy：proxy节点
	NodeType *TopologyInfoNodeType `json:"node_type,omitempty"`

	// **参数解释**： 总内存，单位：MB。 **取值范围**： 不涉及。
	MaxMemory *string `json:"max_memory,omitempty"`

	// **参数解释**： 已使用的内存，单位：MB。 **取值范围**： 不涉及。
	UsedMemory *string `json:"used_memory,omitempty"`

	// **参数解释**： 每秒查询率。 **取值范围**： 不涉及。
	Qps *string `json:"qps,omitempty"`

	Bandwidth *BandWidth `json:"bandwidth,omitempty"`

	// **参数解释**： 该实例的DB数量。 **取值范围**： 不涉及。
	DbNum *string `json:"db_num,omitempty"`

	Dbs *KeySpace `json:"dbs,omitempty"`

	// **参数解释**： 关联IP。 **取值范围**： 不涉及。
	RelationIp *string `json:"relation_ip,omitempty"`

	// **参数解释**： 关联端口。 **取值范围**： 不涉及。
	RelationPort *string `json:"relation_port,omitempty"`

	// **参数解释**： 分片ID。 **取值范围**： 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 节点状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	Dims *DimsInfo `json:"dims,omitempty"`
}

func (o TopologyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopologyInfo struct{}"
	}

	return strings.Join([]string{"TopologyInfo", string(data)}, " ")
}

type TopologyInfoNodeType struct {
	value string
}

type TopologyInfoNodeTypeEnum struct {
	MASTER TopologyInfoNodeType
	SLAVE  TopologyInfoNodeType
	PROXY  TopologyInfoNodeType
}

func GetTopologyInfoNodeTypeEnum() TopologyInfoNodeTypeEnum {
	return TopologyInfoNodeTypeEnum{
		MASTER: TopologyInfoNodeType{
			value: "master",
		},
		SLAVE: TopologyInfoNodeType{
			value: "slave",
		},
		PROXY: TopologyInfoNodeType{
			value: "proxy",
		},
	}
}

func (c TopologyInfoNodeType) Value() string {
	return c.value
}

func (c TopologyInfoNodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TopologyInfoNodeType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
