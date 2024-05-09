package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StarRocksReplicationInfo StarRocks同步任务信息。
type StarRocksReplicationInfo struct {

	// GaussDB(for MySQL)数据库。
	SourceDatabase *string `json:"source_database,omitempty"`

	// 目标数据库。
	TargetDatabase *string `json:"target_database,omitempty"`

	// 同步任务名。
	TaskName *string `json:"task_name,omitempty"`

	// 当前状态。Yes:正常;No:异常。
	Status *StarRocksReplicationInfoStatus `json:"status,omitempty"`

	// 同步阶段。wait:等待同步;incremental:增量同步;full:全量同步;cancelled:删除;paused:暂停同步。
	Stage *StarRocksReplicationInfoStage `json:"stage,omitempty"`

	// 进度百分比。
	Percentage *string `json:"percentage,omitempty"`

	// 追赶阶段。wait:等待同步;incremental:增量同步;full:全量同步;cancelled:删除;paused:暂停同步。
	CatchupStage *StarRocksReplicationInfoCatchupStage `json:"catchup_stage,omitempty"`

	// 追赶进度百分比。
	CatchupPercentage *string `json:"catchup_percentage,omitempty"`
}

func (o StarRocksReplicationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksReplicationInfo struct{}"
	}

	return strings.Join([]string{"StarRocksReplicationInfo", string(data)}, " ")
}

type StarRocksReplicationInfoStatus struct {
	value string
}

type StarRocksReplicationInfoStatusEnum struct {
	YES StarRocksReplicationInfoStatus
	NO  StarRocksReplicationInfoStatus
}

func GetStarRocksReplicationInfoStatusEnum() StarRocksReplicationInfoStatusEnum {
	return StarRocksReplicationInfoStatusEnum{
		YES: StarRocksReplicationInfoStatus{
			value: "Yes",
		},
		NO: StarRocksReplicationInfoStatus{
			value: "No",
		},
	}
}

func (c StarRocksReplicationInfoStatus) Value() string {
	return c.value
}

func (c StarRocksReplicationInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StarRocksReplicationInfoStatus) UnmarshalJSON(b []byte) error {
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

type StarRocksReplicationInfoStage struct {
	value string
}

type StarRocksReplicationInfoStageEnum struct {
	WAIT        StarRocksReplicationInfoStage
	INCREMENTAL StarRocksReplicationInfoStage
	FULL        StarRocksReplicationInfoStage
	CANCELLED   StarRocksReplicationInfoStage
	PAUSED      StarRocksReplicationInfoStage
}

func GetStarRocksReplicationInfoStageEnum() StarRocksReplicationInfoStageEnum {
	return StarRocksReplicationInfoStageEnum{
		WAIT: StarRocksReplicationInfoStage{
			value: "wait",
		},
		INCREMENTAL: StarRocksReplicationInfoStage{
			value: "incremental",
		},
		FULL: StarRocksReplicationInfoStage{
			value: "full",
		},
		CANCELLED: StarRocksReplicationInfoStage{
			value: "cancelled",
		},
		PAUSED: StarRocksReplicationInfoStage{
			value: "paused",
		},
	}
}

func (c StarRocksReplicationInfoStage) Value() string {
	return c.value
}

func (c StarRocksReplicationInfoStage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StarRocksReplicationInfoStage) UnmarshalJSON(b []byte) error {
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

type StarRocksReplicationInfoCatchupStage struct {
	value string
}

type StarRocksReplicationInfoCatchupStageEnum struct {
	WAIT        StarRocksReplicationInfoCatchupStage
	INCREMENTAL StarRocksReplicationInfoCatchupStage
	FULL        StarRocksReplicationInfoCatchupStage
	CANCELLED   StarRocksReplicationInfoCatchupStage
	PAUSED      StarRocksReplicationInfoCatchupStage
}

func GetStarRocksReplicationInfoCatchupStageEnum() StarRocksReplicationInfoCatchupStageEnum {
	return StarRocksReplicationInfoCatchupStageEnum{
		WAIT: StarRocksReplicationInfoCatchupStage{
			value: "wait",
		},
		INCREMENTAL: StarRocksReplicationInfoCatchupStage{
			value: "incremental",
		},
		FULL: StarRocksReplicationInfoCatchupStage{
			value: "full",
		},
		CANCELLED: StarRocksReplicationInfoCatchupStage{
			value: "cancelled",
		},
		PAUSED: StarRocksReplicationInfoCatchupStage{
			value: "paused",
		},
	}
}

func (c StarRocksReplicationInfoCatchupStage) Value() string {
	return c.value
}

func (c StarRocksReplicationInfoCatchupStage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StarRocksReplicationInfoCatchupStage) UnmarshalJSON(b []byte) error {
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
