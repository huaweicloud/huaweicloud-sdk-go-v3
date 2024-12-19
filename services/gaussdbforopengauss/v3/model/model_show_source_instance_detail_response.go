package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSourceInstanceDetailResponse Response Object
type ShowSourceInstanceDetailResponse struct {

	// 实例部署形态。集中式Ha(主备)、分布式Independent(独立部署)。
	ClusterMode *ShowSourceInstanceDetailResponseClusterMode `json:"cluster_mode,omitempty"`

	// 实例模型，企业版enterprise，标准版standard，基础版basic。
	InstanceMode *ShowSourceInstanceDetailResponseInstanceMode `json:"instance_mode,omitempty"`

	// 磁盘大小，单位：GB。
	DataVolumeSize *string `json:"data_volume_size,omitempty"`

	// 解决方案模板类型。集中式Ha一般用triset，分布式Independent一般为空或者默认hws。  描述如下：  triset：高可用(1主2备)  hws：默认。
	Solution *ShowSourceInstanceDetailResponseSolution `json:"solution,omitempty"`

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

func (o ShowSourceInstanceDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSourceInstanceDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowSourceInstanceDetailResponse", string(data)}, " ")
}

type ShowSourceInstanceDetailResponseClusterMode struct {
	value string
}

type ShowSourceInstanceDetailResponseClusterModeEnum struct {
	HA          ShowSourceInstanceDetailResponseClusterMode
	INDEPENDENT ShowSourceInstanceDetailResponseClusterMode
}

func GetShowSourceInstanceDetailResponseClusterModeEnum() ShowSourceInstanceDetailResponseClusterModeEnum {
	return ShowSourceInstanceDetailResponseClusterModeEnum{
		HA: ShowSourceInstanceDetailResponseClusterMode{
			value: "Ha",
		},
		INDEPENDENT: ShowSourceInstanceDetailResponseClusterMode{
			value: "Independent",
		},
	}
}

func (c ShowSourceInstanceDetailResponseClusterMode) Value() string {
	return c.value
}

func (c ShowSourceInstanceDetailResponseClusterMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSourceInstanceDetailResponseClusterMode) UnmarshalJSON(b []byte) error {
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

type ShowSourceInstanceDetailResponseInstanceMode struct {
	value string
}

type ShowSourceInstanceDetailResponseInstanceModeEnum struct {
	BASIC      ShowSourceInstanceDetailResponseInstanceMode
	STANDARD   ShowSourceInstanceDetailResponseInstanceMode
	ENTERPRISE ShowSourceInstanceDetailResponseInstanceMode
}

func GetShowSourceInstanceDetailResponseInstanceModeEnum() ShowSourceInstanceDetailResponseInstanceModeEnum {
	return ShowSourceInstanceDetailResponseInstanceModeEnum{
		BASIC: ShowSourceInstanceDetailResponseInstanceMode{
			value: "basic",
		},
		STANDARD: ShowSourceInstanceDetailResponseInstanceMode{
			value: "standard",
		},
		ENTERPRISE: ShowSourceInstanceDetailResponseInstanceMode{
			value: "enterprise",
		},
	}
}

func (c ShowSourceInstanceDetailResponseInstanceMode) Value() string {
	return c.value
}

func (c ShowSourceInstanceDetailResponseInstanceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSourceInstanceDetailResponseInstanceMode) UnmarshalJSON(b []byte) error {
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

type ShowSourceInstanceDetailResponseSolution struct {
	value string
}

type ShowSourceInstanceDetailResponseSolutionEnum struct {
	TRISET ShowSourceInstanceDetailResponseSolution
	HWS    ShowSourceInstanceDetailResponseSolution
}

func GetShowSourceInstanceDetailResponseSolutionEnum() ShowSourceInstanceDetailResponseSolutionEnum {
	return ShowSourceInstanceDetailResponseSolutionEnum{
		TRISET: ShowSourceInstanceDetailResponseSolution{
			value: "triset",
		},
		HWS: ShowSourceInstanceDetailResponseSolution{
			value: "hws",
		},
	}
}

func (c ShowSourceInstanceDetailResponseSolution) Value() string {
	return c.value
}

func (c ShowSourceInstanceDetailResponseSolution) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSourceInstanceDetailResponseSolution) UnmarshalJSON(b []byte) error {
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
