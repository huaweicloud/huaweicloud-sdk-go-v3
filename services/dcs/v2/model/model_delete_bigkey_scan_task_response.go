package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteBigkeyScanTaskResponse Response Object
type DeleteBigkeyScanTaskResponse struct {

	// 大key分析记录ID
	Id *string `json:"id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 分析任务状态
	Status *DeleteBigkeyScanTaskResponseStatus `json:"status,omitempty"`

	// 分析方式
	ScanType *DeleteBigkeyScanTaskResponseScanType `json:"scan_type,omitempty"`

	// 分析任务创建时间,格式为：\"2020-06-15T02:21:18.669Z\"
	CreatedAt *string `json:"created_at,omitempty"`

	// 分析任务开始时间,格式为：\"2020-06-15T02:21:18.669Z\"
	StartedAt *string `json:"started_at,omitempty"`

	// 分析任务结束时间,格式为：\"2020-06-15T02:21:18.669Z\"
	FinishedAt *string `json:"finished_at,omitempty"`

	// 大key的数量
	Num *int32 `json:"num,omitempty"`

	// 大key记录
	Keys           *[]BigkeysBody `json:"keys,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o DeleteBigkeyScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBigkeyScanTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteBigkeyScanTaskResponse", string(data)}, " ")
}

type DeleteBigkeyScanTaskResponseStatus struct {
	value string
}

type DeleteBigkeyScanTaskResponseStatusEnum struct {
	WAITING DeleteBigkeyScanTaskResponseStatus
	RUNNING DeleteBigkeyScanTaskResponseStatus
	SUCCESS DeleteBigkeyScanTaskResponseStatus
	FAILED  DeleteBigkeyScanTaskResponseStatus
}

func GetDeleteBigkeyScanTaskResponseStatusEnum() DeleteBigkeyScanTaskResponseStatusEnum {
	return DeleteBigkeyScanTaskResponseStatusEnum{
		WAITING: DeleteBigkeyScanTaskResponseStatus{
			value: "waiting",
		},
		RUNNING: DeleteBigkeyScanTaskResponseStatus{
			value: "running",
		},
		SUCCESS: DeleteBigkeyScanTaskResponseStatus{
			value: "success",
		},
		FAILED: DeleteBigkeyScanTaskResponseStatus{
			value: "failed",
		},
	}
}

func (c DeleteBigkeyScanTaskResponseStatus) Value() string {
	return c.value
}

func (c DeleteBigkeyScanTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteBigkeyScanTaskResponseStatus) UnmarshalJSON(b []byte) error {
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

type DeleteBigkeyScanTaskResponseScanType struct {
	value string
}

type DeleteBigkeyScanTaskResponseScanTypeEnum struct {
	MANUAL DeleteBigkeyScanTaskResponseScanType
	AUTO   DeleteBigkeyScanTaskResponseScanType
}

func GetDeleteBigkeyScanTaskResponseScanTypeEnum() DeleteBigkeyScanTaskResponseScanTypeEnum {
	return DeleteBigkeyScanTaskResponseScanTypeEnum{
		MANUAL: DeleteBigkeyScanTaskResponseScanType{
			value: "manual",
		},
		AUTO: DeleteBigkeyScanTaskResponseScanType{
			value: "auto",
		},
	}
}

func (c DeleteBigkeyScanTaskResponseScanType) Value() string {
	return c.value
}

func (c DeleteBigkeyScanTaskResponseScanType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteBigkeyScanTaskResponseScanType) UnmarshalJSON(b []byte) error {
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
