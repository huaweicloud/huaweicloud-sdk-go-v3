package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateExportTaskReq struct {

	// OBS桶名
	BucketName string `json:"bucket_name"`

	// 开始时间(Unix timestamp),单位:毫秒。
	StartAt int64 `json:"start_at"`

	// 结束时间(Unix timestamp),单位:毫秒。
	EndAt int64 `json:"end_at"`

	// OBS文件目录
	FilePath *string `json:"file_path,omitempty"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`

	// 排序字段。 - collectTime 收集时间 - occurrenceTime 发生时间 - lastSec 事务持续时间 - waitLockStructCount 持有锁数量 - holdLockStructCount 等待锁数量
	Order *CreateExportTaskReqOrder `json:"order,omitempty"`

	// 排序规则。 - asc 升序 - desc 降序
	OrderBy *CreateExportTaskReqOrderBy `json:"order_by,omitempty"`

	// 持续时间下限
	LastSecMin *int64 `json:"last_sec_min,omitempty"`

	// 持续时间上限
	LastSecMax *int64 `json:"last_sec_max,omitempty"`
}

func (o CreateExportTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExportTaskReq struct{}"
	}

	return strings.Join([]string{"CreateExportTaskReq", string(data)}, " ")
}

type CreateExportTaskReqOrder struct {
	value string
}

type CreateExportTaskReqOrderEnum struct {
	COLLECT_TIME           CreateExportTaskReqOrder
	OCCURRENCE_TIME        CreateExportTaskReqOrder
	LAST_SEC               CreateExportTaskReqOrder
	WAIT_LOCK_STRUCT_COUNT CreateExportTaskReqOrder
	HOLD_LOCK_STRUCT_COUNT CreateExportTaskReqOrder
}

func GetCreateExportTaskReqOrderEnum() CreateExportTaskReqOrderEnum {
	return CreateExportTaskReqOrderEnum{
		COLLECT_TIME: CreateExportTaskReqOrder{
			value: "collectTime",
		},
		OCCURRENCE_TIME: CreateExportTaskReqOrder{
			value: "occurrenceTime",
		},
		LAST_SEC: CreateExportTaskReqOrder{
			value: "lastSec",
		},
		WAIT_LOCK_STRUCT_COUNT: CreateExportTaskReqOrder{
			value: "waitLockStructCount",
		},
		HOLD_LOCK_STRUCT_COUNT: CreateExportTaskReqOrder{
			value: "holdLockStructCount",
		},
	}
}

func (c CreateExportTaskReqOrder) Value() string {
	return c.value
}

func (c CreateExportTaskReqOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateExportTaskReqOrder) UnmarshalJSON(b []byte) error {
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

type CreateExportTaskReqOrderBy struct {
	value string
}

type CreateExportTaskReqOrderByEnum struct {
	ASC  CreateExportTaskReqOrderBy
	DESC CreateExportTaskReqOrderBy
}

func GetCreateExportTaskReqOrderByEnum() CreateExportTaskReqOrderByEnum {
	return CreateExportTaskReqOrderByEnum{
		ASC: CreateExportTaskReqOrderBy{
			value: "asc",
		},
		DESC: CreateExportTaskReqOrderBy{
			value: "desc",
		},
	}
}

func (c CreateExportTaskReqOrderBy) Value() string {
	return c.value
}

func (c CreateExportTaskReqOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateExportTaskReqOrderBy) UnmarshalJSON(b []byte) error {
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
