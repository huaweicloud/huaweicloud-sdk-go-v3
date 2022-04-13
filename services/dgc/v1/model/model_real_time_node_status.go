package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type RealTimeNodeStatus struct {
	Name *string `json:"name,omitempty"`

	Status *RealTimeNodeStatusStatus `json:"status,omitempty"`

	LogPath *string `json:"logPath,omitempty"`

	NodeType *RealTimeNodeStatusNodeType `json:"nodeType,omitempty"`
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
	STARTTING RealTimeNodeStatusStatus
	NORMAL    RealTimeNodeStatusStatus
	EXCEPTION RealTimeNodeStatusStatus
	STOPPING  RealTimeNodeStatusStatus
	STOPPED   RealTimeNodeStatusStatus
}

func GetRealTimeNodeStatusStatusEnum() RealTimeNodeStatusStatusEnum {
	return RealTimeNodeStatusStatusEnum{
		STARTTING: RealTimeNodeStatusStatus{
			value: "STARTTING",
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
	}
}

func (c RealTimeNodeStatusStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RealTimeNodeStatusStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type RealTimeNodeStatusNodeType struct {
	value string
}

type RealTimeNodeStatusNodeTypeEnum struct {
	HIVE_SQL           RealTimeNodeStatusNodeType
	SPARK_SQL          RealTimeNodeStatusNodeType
	DWSSQL             RealTimeNodeStatusNodeType
	DLISQL             RealTimeNodeStatusNodeType
	SHELL              RealTimeNodeStatusNodeType
	CDM_JOB            RealTimeNodeStatusNodeType
	DIS_TRANSFER_TASK  RealTimeNodeStatusNodeType
	CS_JOB             RealTimeNodeStatusNodeType
	CLOUD_TABLE_MANAGE RealTimeNodeStatusNodeType
	OBS_MANAGER        RealTimeNodeStatusNodeType
	RESTAPI            RealTimeNodeStatusNodeType
	MACHINE_LEARNING   RealTimeNodeStatusNodeType
	SMN                RealTimeNodeStatusNodeType
	MRS_SPARK          RealTimeNodeStatusNodeType
	MAP_REDUCE         RealTimeNodeStatusNodeType
	DLI_SPARK          RealTimeNodeStatusNodeType
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
		CLOUD_TABLE_MANAGE: RealTimeNodeStatusNodeType{
			value: "CloudTableManage",
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

func (c RealTimeNodeStatusNodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RealTimeNodeStatusNodeType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
