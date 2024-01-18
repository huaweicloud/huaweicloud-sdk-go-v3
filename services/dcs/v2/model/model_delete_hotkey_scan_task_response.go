package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteHotkeyScanTaskResponse Response Object
type DeleteHotkeyScanTaskResponse struct {

	// 热key分析记录ID
	Id *string `json:"id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 分析任务状态
	Status *DeleteHotkeyScanTaskResponseStatus `json:"status,omitempty"`

	// 分析方式
	ScanType *DeleteHotkeyScanTaskResponseScanType `json:"scan_type,omitempty"`

	// 分析任务创建时间,格式为：\"2020-06-15T02:21:18.669Z\"
	CreatedAt *string `json:"created_at,omitempty"`

	// 分析任务开始时间,格式为：\"2020-06-15T02:21:18.669Z\"
	StartedAt *string `json:"started_at,omitempty"`

	// 分析任务结束时间,格式为：\"2020-06-15T02:21:18.669Z\"
	FinishedAt *string `json:"finished_at,omitempty"`

	// 热key的数量
	Num *int32 `json:"num,omitempty"`

	// 热key记录
	Keys           *[]HotkeysBody `json:"keys,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o DeleteHotkeyScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHotkeyScanTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteHotkeyScanTaskResponse", string(data)}, " ")
}

type DeleteHotkeyScanTaskResponseStatus struct {
	value string
}

type DeleteHotkeyScanTaskResponseStatusEnum struct {
	WAITING DeleteHotkeyScanTaskResponseStatus
	RUNNING DeleteHotkeyScanTaskResponseStatus
	SUCCESS DeleteHotkeyScanTaskResponseStatus
	FAILED  DeleteHotkeyScanTaskResponseStatus
}

func GetDeleteHotkeyScanTaskResponseStatusEnum() DeleteHotkeyScanTaskResponseStatusEnum {
	return DeleteHotkeyScanTaskResponseStatusEnum{
		WAITING: DeleteHotkeyScanTaskResponseStatus{
			value: "waiting",
		},
		RUNNING: DeleteHotkeyScanTaskResponseStatus{
			value: "running",
		},
		SUCCESS: DeleteHotkeyScanTaskResponseStatus{
			value: "success",
		},
		FAILED: DeleteHotkeyScanTaskResponseStatus{
			value: "failed",
		},
	}
}

func (c DeleteHotkeyScanTaskResponseStatus) Value() string {
	return c.value
}

func (c DeleteHotkeyScanTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteHotkeyScanTaskResponseStatus) UnmarshalJSON(b []byte) error {
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

type DeleteHotkeyScanTaskResponseScanType struct {
	value string
}

type DeleteHotkeyScanTaskResponseScanTypeEnum struct {
	MANUAL DeleteHotkeyScanTaskResponseScanType
	AUTO   DeleteHotkeyScanTaskResponseScanType
}

func GetDeleteHotkeyScanTaskResponseScanTypeEnum() DeleteHotkeyScanTaskResponseScanTypeEnum {
	return DeleteHotkeyScanTaskResponseScanTypeEnum{
		MANUAL: DeleteHotkeyScanTaskResponseScanType{
			value: "manual",
		},
		AUTO: DeleteHotkeyScanTaskResponseScanType{
			value: "auto",
		},
	}
}

func (c DeleteHotkeyScanTaskResponseScanType) Value() string {
	return c.value
}

func (c DeleteHotkeyScanTaskResponseScanType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteHotkeyScanTaskResponseScanType) UnmarshalJSON(b []byte) error {
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
