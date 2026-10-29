package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RealTimeNodeStatus struct {

	// 节点名称。
	Name *string `json:"name,omitempty"`

	// 节点状态： - STARTING：启动中 - NORMAL：正常 - EXCEPTION：异常 - STOPPING： 停止中 - STOPPED：停止 - PAUSE: 暂停 - ABNORMAL: 运行异常/失败 - DISABLE: 节点禁用 - OVERLOAD: 繁忙 - INIT: 初始化
	Status *RealTimeNodeStatusStatus `json:"status,omitempty"`

	// 节点运行日志路径
	LogPath *string `json:"logPath,omitempty"`

	// 节点类型
	NodeType *RealTimeNodeStatusNodeType `json:"nodeType,omitempty"`

	// 上一次实例的执行状态。
	LastInstanceStatus *string `json:"lastInstanceStatus,omitempty"`

	// 节点运行时内部存放的部分数据。
	RunningData *string `json:"runningData,omitempty"`

	// 扩展计数器。
	ExtendCounter *string `json:"extendCounter,omitempty"`
}

func (o RealTimeNodeStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RealTimeNodeStatus struct{}"
	}

	return strings.Join([]string{"RealTimeNodeStatus", string(data)}, " ")
}

type RealTimeNodeStatusStatus struct {
	value string
}

type RealTimeNodeStatusStatusEnum struct {
	STARTING  RealTimeNodeStatusStatus
	NORMAL    RealTimeNodeStatusStatus
	EXCEPTION RealTimeNodeStatusStatus
	STOPPING  RealTimeNodeStatusStatus
	STOPPED   RealTimeNodeStatusStatus
	PAUSE     RealTimeNodeStatusStatus
	ABNORMAL  RealTimeNodeStatusStatus
	DISABLE   RealTimeNodeStatusStatus
	OVERLOAD  RealTimeNodeStatusStatus
	INIT      RealTimeNodeStatusStatus
}

func GetRealTimeNodeStatusStatusEnum() RealTimeNodeStatusStatusEnum {
	return RealTimeNodeStatusStatusEnum{
		STARTING: RealTimeNodeStatusStatus{
			value: "STARTING",
		},
		NORMAL: RealTimeNodeStatusStatus{
			value: "NORMAL",
		},
		EXCEPTION: RealTimeNodeStatusStatus{
			value: "EXCEPTION",
		},
		STOPPING: RealTimeNodeStatusStatus{
			value: "STOPPING",
		},
		STOPPED: RealTimeNodeStatusStatus{
			value: "STOPPED",
		},
		PAUSE: RealTimeNodeStatusStatus{
			value: "PAUSE",
		},
		ABNORMAL: RealTimeNodeStatusStatus{
			value: "ABNORMAL",
		},
		DISABLE: RealTimeNodeStatusStatus{
			value: "DISABLE",
		},
		OVERLOAD: RealTimeNodeStatusStatus{
			value: "OVERLOAD",
		},
		INIT: RealTimeNodeStatusStatus{
			value: "INIT",
		},
	}
}

func (c RealTimeNodeStatusStatus) Value() string {
	return c.value
}

func (c RealTimeNodeStatusStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RealTimeNodeStatusStatus) UnmarshalJSON(b []byte) error {
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

type RealTimeNodeStatusNodeType struct {
	value string
}

type RealTimeNodeStatusNodeTypeEnum struct {
	HIVE_SQL            RealTimeNodeStatusNodeType
	SPARK_SQL           RealTimeNodeStatusNodeType
	DWSSQL              RealTimeNodeStatusNodeType
	DLISQL              RealTimeNodeStatusNodeType
	SHELL               RealTimeNodeStatusNodeType
	CDM_JOB             RealTimeNodeStatusNodeType
	DIS_TRANSFER_TASK   RealTimeNodeStatusNodeType
	CS_JOB              RealTimeNodeStatusNodeType
	CLOUD_TABLE_MANAGER RealTimeNodeStatusNodeType
	OBS_MANAGER         RealTimeNodeStatusNodeType
	RESTAPI             RealTimeNodeStatusNodeType
	MACHINE_LEARNING    RealTimeNodeStatusNodeType
	SMN                 RealTimeNodeStatusNodeType
	MRS_SPARK           RealTimeNodeStatusNodeType
	MAP_REDUCE          RealTimeNodeStatusNodeType
	DLI_SPARK           RealTimeNodeStatusNodeType
}

func GetRealTimeNodeStatusNodeTypeEnum() RealTimeNodeStatusNodeTypeEnum {
	return RealTimeNodeStatusNodeTypeEnum{
		HIVE_SQL: RealTimeNodeStatusNodeType{
			value: "HiveSQL",
		},
		SPARK_SQL: RealTimeNodeStatusNodeType{
			value: "SparkSQL",
		},
		DWSSQL: RealTimeNodeStatusNodeType{
			value: "DWSSQL",
		},
		DLISQL: RealTimeNodeStatusNodeType{
			value: "DLISQL",
		},
		SHELL: RealTimeNodeStatusNodeType{
			value: "Shell",
		},
		CDM_JOB: RealTimeNodeStatusNodeType{
			value: "CDMJob",
		},
		DIS_TRANSFER_TASK: RealTimeNodeStatusNodeType{
			value: "DISTransferTask",
		},
		CS_JOB: RealTimeNodeStatusNodeType{
			value: "CSJob",
		},
		CLOUD_TABLE_MANAGER: RealTimeNodeStatusNodeType{
			value: "CloudTableManager",
		},
		OBS_MANAGER: RealTimeNodeStatusNodeType{
			value: "OBSManager",
		},
		RESTAPI: RealTimeNodeStatusNodeType{
			value: "RESTAPI",
		},
		MACHINE_LEARNING: RealTimeNodeStatusNodeType{
			value: "MachineLearning",
		},
		SMN: RealTimeNodeStatusNodeType{
			value: "SMN",
		},
		MRS_SPARK: RealTimeNodeStatusNodeType{
			value: "MRSSpark",
		},
		MAP_REDUCE: RealTimeNodeStatusNodeType{
			value: "MapReduce",
		},
		DLI_SPARK: RealTimeNodeStatusNodeType{
			value: "DLISpark",
		},
	}
}

func (c RealTimeNodeStatusNodeType) Value() string {
	return c.value
}

func (c RealTimeNodeStatusNodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RealTimeNodeStatusNodeType) UnmarshalJSON(b []byte) error {
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
