package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Node struct {

	// 节点名称，只能包含六种字符：英文字母、数字、中文、中划线、下划线和点号，且长度小于等于128个字符。同一个作业中节点名称不能重复。
	Name string `json:"name"`

	// 节点的类型
	Type NodeType `json:"type"`

	Location *Location `json:"location"`

	// 本本节点依赖的前面的节点名称列表
	PreNodeName *[]string `json:"preNodeName,omitempty"`

	// 节点执行条件，如果配置此参数，本节点是否执行由condition的字段expression所保存的EL表达式计算结果决定
	Conditions *[]Condition `json:"conditions,omitempty"`

	// 节点属性
	Properties []Property `json:"properties"`

	// 轮询节点执行结果时间间隔
	PollingInterval *int32 `json:"pollingInterval,omitempty"`

	// 节点是否超时重试
	ExecTimeOutRetry *string `json:"execTimeOutRetry,omitempty"`

	// 节点最大执行时间
	MaxExecutionTime *int32 `json:"maxExecutionTime,omitempty"`

	// 节点失败重试次数
	RetryTimes *int32 `json:"retryTimes,omitempty"`

	// 失败重试时间间隔
	RetryInterval *int32 `json:"retryInterval,omitempty"`

	// 作业失败策略
	FailPolicy *NodeFailPolicy `json:"failPolicy,omitempty"`

	EventTrigger *Event `json:"eventTrigger,omitempty"`

	CronTrigger *CronTrigger `json:"cronTrigger,omitempty"`
}

func (o Node) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Node struct{}"
	}

	return strings.Join([]string{"Node", string(data)}, " ")
}

type NodeType struct {
	value string
}

type NodeTypeEnum struct {
	HIVE_SQL           NodeType
	SPARK_SQL          NodeType
	DWSSQL             NodeType
	DLISQL             NodeType
	SHELL              NodeType
	CDM_JOB            NodeType
	DIS_TRANSFER_TASK  NodeType
	CS_JOB             NodeType
	CLOUD_TABLE_MANAGE NodeType
	OBS_MANAGER        NodeType
	RESTAPI            NodeType
	MACHINE_LEARNING   NodeType
	SMN                NodeType
	MRS_SPARK          NodeType
	MAP_REDUCE         NodeType
	DLI_SPARK          NodeType
	MRS_FLINK          NodeType
	MRS_FLINK_JOB      NodeType
	MRS_HETU_ENGINE    NodeType
	RDS_SQL            NodeType
	DATA_MIGRATION     NodeType
	ONECLICK_CDC       NodeType
	DUMMY              NodeType
}

func GetNodeTypeEnum() NodeTypeEnum {
	return NodeTypeEnum{
		HIVE_SQL: NodeType{
			value: "HiveSQL",
		},
		SPARK_SQL: NodeType{
			value: "SparkSQL",
		},
		DWSSQL: NodeType{
			value: "DWSSQL",
		},
		DLISQL: NodeType{
			value: "DLISQL",
		},
		SHELL: NodeType{
			value: "Shell",
		},
		CDM_JOB: NodeType{
			value: "CDMJob",
		},
		DIS_TRANSFER_TASK: NodeType{
			value: "DISTransferTask",
		},
		CS_JOB: NodeType{
			value: "CSJob",
		},
		CLOUD_TABLE_MANAGE: NodeType{
			value: "CloudTableManage",
		},
		OBS_MANAGER: NodeType{
			value: "OBSManager",
		},
		RESTAPI: NodeType{
			value: "RESTAPI",
		},
		MACHINE_LEARNING: NodeType{
			value: "MachineLearning",
		},
		SMN: NodeType{
			value: "SMN",
		},
		MRS_SPARK: NodeType{
			value: "MRSSpark",
		},
		MAP_REDUCE: NodeType{
			value: "MapReduce",
		},
		DLI_SPARK: NodeType{
			value: "DLISpark",
		},
		MRS_FLINK: NodeType{
			value: "MRSFlink",
		},
		MRS_FLINK_JOB: NodeType{
			value: "MRSFlinkJob",
		},
		MRS_HETU_ENGINE: NodeType{
			value: "MRSHetuEngine",
		},
		RDS_SQL: NodeType{
			value: "RDS SQL",
		},
		DATA_MIGRATION: NodeType{
			value: "DataMigration",
		},
		ONECLICK_CDC: NodeType{
			value: "OneclickCDC",
		},
		DUMMY: NodeType{
			value: "Dummy",
		},
	}
}

func (c NodeType) Value() string {
	return c.value
}

func (c NodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeType) UnmarshalJSON(b []byte) error {
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
	FAIL       NodeFailPolicy
	IGNORE     NodeFailPolicy
	SUSPEND    NodeFailPolicy
	FAIL_CHILD NodeFailPolicy
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
		FAIL_CHILD: NodeFailPolicy{
			value: "FAIL_CHILD",
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
