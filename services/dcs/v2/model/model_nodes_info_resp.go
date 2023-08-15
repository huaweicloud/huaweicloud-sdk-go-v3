package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NodesInfoResp struct {

	// 逻辑节点ID
	LogicalNodeId *string `json:"logical_node_id,omitempty"`

	// 节点名称
	Name *string `json:"name,omitempty"`

	// 节点状态，所有值如下: - Creating：创建中。 - Active：运行中。 - Inactive：故障。 - Deleting：删除中。 - AddSharding：添加分片中。
	Status *NodesInfoRespStatus `json:"status,omitempty"`

	// 可用区code
	AzCode *string `json:"az_code,omitempty"`

	// 节点角色，所有值如下: - redis-server：Redis server节点。 - redis-proxy：proxy节点。
	NodeRole *NodesInfoRespNodeRole `json:"node_role,omitempty"`

	// 节点主从角色: - master：主 - slave：从 - proxy: proxy实例节点角色为\"proxy\"
	NodeType *NodesInfoRespNodeType `json:"node_type,omitempty"`

	// 节点的IP
	NodeIp *string `json:"node_ip,omitempty"`

	// 节点的port
	NodePort *string `json:"node_port,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 节点权重
	PriorityWeight *int32 `json:"priority_weight,omitempty"`

	// 节点的IP是否可直接访问
	IsAccess *bool `json:"is_access,omitempty"`

	// 分片ID
	GroupId *string `json:"group_id,omitempty"`

	// 分片名称
	GroupName *string `json:"group_name,omitempty"`

	// 是否从只读域名中摘除IP
	IsRemoveIp *bool `json:"is_remove_ip,omitempty"`

	// 副本id
	ReplicationId *string `json:"replication_id,omitempty"`

	// 副本对应的监控指标维度信息。可用于调用云监控服务的查询监控数据指标相关接口 - 副本的监控维度为多维度，返回数组中包含两个维度信息。从云监控查询监控数据时，要按多维度传递指标维度参数，才能查询到监控指标值 - 第一个维度为副本父维度信息 维度名称为dcs_instance_id，维度值对应副本所在的实例ID - 第二个维度，维度名称为dcs_cluster_redis_node,维度值为副本的监控对象ID，与副本ID和节点ID不同。
	Dimensions *[]InstanceReplicationDimensionsInfo `json:"dimensions,omitempty"`
}

func (o NodesInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodesInfoResp struct{}"
	}

	return strings.Join([]string{"NodesInfoResp", string(data)}, " ")
}

type NodesInfoRespStatus struct {
	value string
}

type NodesInfoRespStatusEnum struct {
	CREATING     NodesInfoRespStatus
	ACTIVE       NodesInfoRespStatus
	INACTIVE     NodesInfoRespStatus
	DELETING     NodesInfoRespStatus
	ADD_SHARDING NodesInfoRespStatus
}

func GetNodesInfoRespStatusEnum() NodesInfoRespStatusEnum {
	return NodesInfoRespStatusEnum{
		CREATING: NodesInfoRespStatus{
			value: "Creating",
		},
		ACTIVE: NodesInfoRespStatus{
			value: "Active",
		},
		INACTIVE: NodesInfoRespStatus{
			value: "Inactive",
		},
		DELETING: NodesInfoRespStatus{
			value: "Deleting",
		},
		ADD_SHARDING: NodesInfoRespStatus{
			value: "AddSharding",
		},
	}
}

func (c NodesInfoRespStatus) Value() string {
	return c.value
}

func (c NodesInfoRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodesInfoRespStatus) UnmarshalJSON(b []byte) error {
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

type NodesInfoRespNodeRole struct {
	value string
}

type NodesInfoRespNodeRoleEnum struct {
	REDIS_SERVER NodesInfoRespNodeRole
	REDIS_PROXY  NodesInfoRespNodeRole
}

func GetNodesInfoRespNodeRoleEnum() NodesInfoRespNodeRoleEnum {
	return NodesInfoRespNodeRoleEnum{
		REDIS_SERVER: NodesInfoRespNodeRole{
			value: "redis-server",
		},
		REDIS_PROXY: NodesInfoRespNodeRole{
			value: "redis-proxy",
		},
	}
}

func (c NodesInfoRespNodeRole) Value() string {
	return c.value
}

func (c NodesInfoRespNodeRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodesInfoRespNodeRole) UnmarshalJSON(b []byte) error {
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

type NodesInfoRespNodeType struct {
	value string
}

type NodesInfoRespNodeTypeEnum struct {
	MASTER NodesInfoRespNodeType
	SLAVE  NodesInfoRespNodeType
	PROXY  NodesInfoRespNodeType
}

func GetNodesInfoRespNodeTypeEnum() NodesInfoRespNodeTypeEnum {
	return NodesInfoRespNodeTypeEnum{
		MASTER: NodesInfoRespNodeType{
			value: "master",
		},
		SLAVE: NodesInfoRespNodeType{
			value: "slave",
		},
		PROXY: NodesInfoRespNodeType{
			value: "proxy",
		},
	}
}

func (c NodesInfoRespNodeType) Value() string {
	return c.value
}

func (c NodesInfoRespNodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodesInfoRespNodeType) UnmarshalJSON(b []byte) error {
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
