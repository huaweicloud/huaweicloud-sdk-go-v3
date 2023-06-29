package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceSnapshotResponse Response Object
type ShowInstanceSnapshotResponse struct {

	// 实例部署形态。集中式Ha(主备)、分布式Independent(独立部署)。
	ClusterMode *ShowInstanceSnapshotResponseClusterMode `json:"cluster_mode,omitempty"`

	// 实例模型，企业版enterprise，标准版standard，基础版basic。
	InstanceMode *ShowInstanceSnapshotResponseInstanceMode `json:"instance_mode,omitempty"`

	// 磁盘大小，单位：GB。
	DataVolumeSize *string `json:"data_volume_size,omitempty"`

	// 解决方案模板类型。集中式Ha一般用triset，分布式Independent一般为空或者默认hws。  描述如下：  triset：高可用(1主2备)  hws：默认。
	Solution *ShowInstanceSnapshotResponseSolution `json:"solution,omitempty"`

	// 节点数量。
	NodeNum *int32 `json:"node_num,omitempty"`

	// 协调节点数量。
	CoordinatorNum *int32 `json:"coordinator_num,omitempty"`

	// 分片数量。
	ShardingNum *int32 `json:"sharding_num,omitempty"`

	// 副本数量。
	ReplicaNum *int32 `json:"replica_num,omitempty"`

	// 引擎版本。
	EngineVersion  *string `json:"engine_version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceSnapshotResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceSnapshotResponse", string(data)}, " ")
}

type ShowInstanceSnapshotResponseClusterMode struct {
	value string
}

type ShowInstanceSnapshotResponseClusterModeEnum struct {
	HA          ShowInstanceSnapshotResponseClusterMode
	INDEPENDENT ShowInstanceSnapshotResponseClusterMode
}

func GetShowInstanceSnapshotResponseClusterModeEnum() ShowInstanceSnapshotResponseClusterModeEnum {
	return ShowInstanceSnapshotResponseClusterModeEnum{
		HA: ShowInstanceSnapshotResponseClusterMode{
			value: "Ha",
		},
		INDEPENDENT: ShowInstanceSnapshotResponseClusterMode{
			value: "Independent",
		},
	}
}

func (c ShowInstanceSnapshotResponseClusterMode) Value() string {
	return c.value
}

func (c ShowInstanceSnapshotResponseClusterMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceSnapshotResponseClusterMode) UnmarshalJSON(b []byte) error {
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

type ShowInstanceSnapshotResponseInstanceMode struct {
	value string
}

type ShowInstanceSnapshotResponseInstanceModeEnum struct {
	BASIC      ShowInstanceSnapshotResponseInstanceMode
	STANDARD   ShowInstanceSnapshotResponseInstanceMode
	ENTERPRISE ShowInstanceSnapshotResponseInstanceMode
}

func GetShowInstanceSnapshotResponseInstanceModeEnum() ShowInstanceSnapshotResponseInstanceModeEnum {
	return ShowInstanceSnapshotResponseInstanceModeEnum{
		BASIC: ShowInstanceSnapshotResponseInstanceMode{
			value: "basic",
		},
		STANDARD: ShowInstanceSnapshotResponseInstanceMode{
			value: "standard",
		},
		ENTERPRISE: ShowInstanceSnapshotResponseInstanceMode{
			value: "enterprise",
		},
	}
}

func (c ShowInstanceSnapshotResponseInstanceMode) Value() string {
	return c.value
}

func (c ShowInstanceSnapshotResponseInstanceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceSnapshotResponseInstanceMode) UnmarshalJSON(b []byte) error {
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

type ShowInstanceSnapshotResponseSolution struct {
	value string
}

type ShowInstanceSnapshotResponseSolutionEnum struct {
	TRISET ShowInstanceSnapshotResponseSolution
	HWS    ShowInstanceSnapshotResponseSolution
}

func GetShowInstanceSnapshotResponseSolutionEnum() ShowInstanceSnapshotResponseSolutionEnum {
	return ShowInstanceSnapshotResponseSolutionEnum{
		TRISET: ShowInstanceSnapshotResponseSolution{
			value: "triset",
		},
		HWS: ShowInstanceSnapshotResponseSolution{
			value: "hws",
		},
	}
}

func (c ShowInstanceSnapshotResponseSolution) Value() string {
	return c.value
}

func (c ShowInstanceSnapshotResponseSolution) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceSnapshotResponseSolution) UnmarshalJSON(b []byte) error {
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
