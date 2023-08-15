package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Node struct {
	Name *string `json:"name,omitempty"`

	// 节点的类型
	NodeType *NodeNodeType `json:"nodeType,omitempty"`

	Location *Location `json:"location,omitempty"`

	// 本节点依赖的前一个节点名称
	PreNodeNames *string `json:"preNodeNames,omitempty"`

	// 节点执行条件
	Condition *[]Condition `json:"condition,omitempty"`

	// 节点的属性
	NodeProperties *string `json:"nodeProperties,omitempty"`

	// 轮询节点执行结果时间间隔
	PollingInterval *int32 `json:"pollingInterval,omitempty"`

	// 节点最大执行时间
	MaxExecutionTime *int32 `json:"maxExecutionTime,omitempty"`

	// 节点失败重试次数
	RetryTimes *int32 `json:"retryTimes,omitempty"`

	// 失败重试时间间隔
	RetryInterval *int32 `json:"retryInterval,omitempty"`

	// 作业失败策略
	FailPolicy *NodeFailPolicy `json:"failPolicy,omitempty"`

	EventTrigger *Event `json:"eventTrigger,omitempty"`

	CronTrigger *Cron `json:"cronTrigger,omitempty"`
}

func (o Node) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Node struct{}"
	}

	return strings.Join([]string{"Node", string(data)}, " ")
}

type NodeNodeType struct {
	value string
}

type NodeNodeTypeEnum struct {
	HIVE_SQL           NodeNodeType
	SPARK_SQL          NodeNodeType
	DWSSQL             NodeNodeType
	DLISQL             NodeNodeType
	SHELL              NodeNodeType
	CDM_JOB            NodeNodeType
	DIS_TRANSFER_TASK  NodeNodeType
	CS_JOB             NodeNodeType
	CLOUD_TABLE_MANAGE NodeNodeType
	OBS_MANAGER        NodeNodeType
	RESTAPI            NodeNodeType
	MACHINE_LEARNING   NodeNodeType
	SMN                NodeNodeType
	MRS_SPARK          NodeNodeType
	MAP_REDUCE         NodeNodeType
	DLI_SPARK          NodeNodeType
}

func GetNodeNodeTypeEnum() NodeNodeTypeEnum {
	return NodeNodeTypeEnum{
		HIVE_SQL: NodeNodeType{
			value: "HiveSQL",
		},
		SPARK_SQL: NodeNodeType{
			value: "SparkSQL",
		},
		DWSSQL: NodeNodeType{
			value: "DWSSQL",
		},
		DLISQL: NodeNodeType{
			value: "DLISQL",
		},
		SHELL: NodeNodeType{
			value: "Shell",
		},
		CDM_JOB: NodeNodeType{
			value: "CDMJob",
		},
		DIS_TRANSFER_TASK: NodeNodeType{
			value: "DISTransferTask",
		},
		CS_JOB: NodeNodeType{
			value: "CSJob",
		},
		CLOUD_TABLE_MANAGE: NodeNodeType{
			value: "CloudTableManage",
		},
		OBS_MANAGER: NodeNodeType{
			value: "OBSManager",
		},
		RESTAPI: NodeNodeType{
			value: "RESTAPI",
		},
		MACHINE_LEARNING: NodeNodeType{
			value: "MachineLearning",
		},
		SMN: NodeNodeType{
			value: "SMN",
		},
		MRS_SPARK: NodeNodeType{
			value: "MRSSpark",
		},
		MAP_REDUCE: NodeNodeType{
			value: "MapReduce",
		},
		DLI_SPARK: NodeNodeType{
			value: "DLISpark",
		},
	}
}

func (c NodeNodeType) Value() string {
	return c.value
}

func (c NodeNodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeNodeType) UnmarshalJSON(b []byte) error {
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

type NodeFailPolicy struct {
	value string
}

type NodeFailPolicyEnum struct {
	FAIL    NodeFailPolicy
	IGNORE  NodeFailPolicy
	SUSPEND NodeFailPolicy
}

func GetNodeFailPolicyEnum() NodeFailPolicyEnum {
	return NodeFailPolicyEnum{
		FAIL: NodeFailPolicy{
			value: "FAIL",
		},
		IGNORE: NodeFailPolicy{
			value: "IGNORE",
		},
		SUSPEND: NodeFailPolicy{
			value: "SUSPEND",
		},
	}
}

func (c NodeFailPolicy) Value() string {
	return c.value
}

func (c NodeFailPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeFailPolicy) UnmarshalJSON(b []byte) error {
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
