package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Node 节点对象
type Node struct {

	// 节点名称。只能包含六种字符：英文字母、数字、中文、中划线、下划线和点号。同一个作业中节点名称不能重复。
	Name string `json:"name"`

	// 节点类型： - HiveSQL：执行Hive SQL脚本 - SparkSQL：执行Spark SQL脚本 - DWSSQL：执行DWS SQL脚本 - DLISQL：执行DLI SQL脚本 - RDSSQL：执行RDS SQL脚本 - Shell：执行Shell脚本 - Python：执行Python脚本 - DISTransferTask：创建DIS转储任务 - CDMJob：执行CDM作业 - OBSManager：执行OBS相关操作 - Dummy：虚拟节点 - RESTAPI：执行Rest API调用 - DLISpark：执行DLI Spark作业 - SMN：执行SMN通知 - MRSSpark：执行MRS Spark作业 - MapReduce：执行MapReduce作业 - MRSFlink：执行MRS服务的FLlink作业 - MRSFlinkJob：执行MRS服务的FlinkJob作业 - MRSHetuEngine: 执行MRS服务的HetuEngine作业
	Type NodeType `json:"type"`

	Location *Location `json:"location"`

	// 本节点依赖的前面的节点名称列表。
	PreNodeNames *[]string `json:"pre_node_names,omitempty"`

	// 节点执行条件，如果配置此参数，本节点是否执行由condition的字段expression所保存的EL表达式计算结果决定
	Conditions *[]Condition `json:"conditions,omitempty"`

	// 节点的属性。
	Properties []Property `json:"properties"`

	// 轮询节点执行结果时间间隔。单位：秒。
	PollingInterval *int32 `json:"polling_interval,omitempty"`

	// 节点是否超时重试
	ExecTimeOutRetry *string `json:"exec_time_out_retry,omitempty"`

	// 节点最大执行时间，如果节点在最大执行时间内还未执行完成，会把节点置为失败状态。单位：分钟。
	MaxExecutionTime *int32 `json:"max_execution_time,omitempty"`

	// 节点失败重试次数。0代表不重试。
	RetryTimes *int32 `json:"retry_times,omitempty"`

	// 失败重试时间间隔。单位：秒。
	RetryInterval *int32 `json:"retry_interval,omitempty"`

	// 作业失败策略： - FAIL：终止当前作业执行计划 - IGNORE：继续执行下一个节点 - SUSPEND：挂起当前作业执行计划 - FAIL_CHILD： 终止后续节点执行计划
	FailPolicy *NodeFailPolicy `json:"fail_policy,omitempty"`

	EventTrigger *Event `json:"event_trigger,omitempty"`

	CronTrigger *Cron `json:"cron_trigger,omitempty"`
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
	HIVE_SQL          NodeType
	SPARK_SQL         NodeType
	DWSSQL            NodeType
	DLISQL            NodeType
	RDSSQL            NodeType
	SHELL             NodeType
	PYTHON            NodeType
	DIS_TRANSFER_TASK NodeType
	CDM_JOB           NodeType
	OBS_MANAGER       NodeType
	DUMMY             NodeType
	RESTAPI           NodeType
	DLI_SPARK         NodeType
	SMN               NodeType
	MRS_SPARK         NodeType
	MAP_REDUCE        NodeType
	MRS_FLINK         NodeType
	MRS_FLINK_JOB     NodeType
	MRS_HETU_ENGINE   NodeType
	DATA_MIGRATION    NodeType
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
		RDSSQL: NodeType{
			value: "RDSSQL",
		},
		SHELL: NodeType{
			value: "Shell",
		},
		PYTHON: NodeType{
			value: "Python",
		},
		DIS_TRANSFER_TASK: NodeType{
			value: "DISTransferTask",
		},
		CDM_JOB: NodeType{
			value: "CDMJob",
		},
		OBS_MANAGER: NodeType{
			value: "OBSManager",
		},
		DUMMY: NodeType{
			value: "Dummy",
		},
		RESTAPI: NodeType{
			value: "RESTAPI",
		},
		DLI_SPARK: NodeType{
			value: "DLISpark",
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
		MRS_FLINK: NodeType{
			value: "MRSFlink",
		},
		MRS_FLINK_JOB: NodeType{
			value: "MRSFlinkJob",
		},
		MRS_HETU_ENGINE: NodeType{
			value: "MRSHetuEngine",
		},
		DATA_MIGRATION: NodeType{
			value: "DataMigration",
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
