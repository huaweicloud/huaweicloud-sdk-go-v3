package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAutoExpireScanTaskResponse Response Object
type CreateAutoExpireScanTaskResponse struct {

	// 过期key扫描记录ID
	Id *string `json:"id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 扫描任务状态
	Status *CreateAutoExpireScanTaskResponseStatus `json:"status,omitempty"`

	// 扫描方式
	ScanType *CreateAutoExpireScanTaskResponseScanType `json:"scan_type,omitempty"`

	// 扫描任务创建时间,格式为：\"2020-06-15T02:21:18.669Z\"
	CreatedAt *string `json:"created_at,omitempty"`

	// 扫描任务开始时间,格式为：\"2020-06-15T02:21:18.669Z\"（创建扫描任务时此值为null，不返回）
	StartedAt *string `json:"started_at,omitempty"`

	// 扫描任务结束时间,格式为：\"2020-06-15T02:21:18.669Z\"（创建扫描任务时此值为null，不返回）
	FinishedAt     *string `json:"finished_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAutoExpireScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAutoExpireScanTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAutoExpireScanTaskResponse", string(data)}, " ")
}

type CreateAutoExpireScanTaskResponseStatus struct {
	value string
}

type CreateAutoExpireScanTaskResponseStatusEnum struct {
	WAITING CreateAutoExpireScanTaskResponseStatus
	RUNNING CreateAutoExpireScanTaskResponseStatus
	SUCCESS CreateAutoExpireScanTaskResponseStatus
	FAILED  CreateAutoExpireScanTaskResponseStatus
}

func GetCreateAutoExpireScanTaskResponseStatusEnum() CreateAutoExpireScanTaskResponseStatusEnum {
	return CreateAutoExpireScanTaskResponseStatusEnum{
		WAITING: CreateAutoExpireScanTaskResponseStatus{
			value: "waiting",
		},
		RUNNING: CreateAutoExpireScanTaskResponseStatus{
			value: "running",
		},
		SUCCESS: CreateAutoExpireScanTaskResponseStatus{
			value: "success",
		},
		FAILED: CreateAutoExpireScanTaskResponseStatus{
			value: "failed",
		},
	}
}

func (c CreateAutoExpireScanTaskResponseStatus) Value() string {
	return c.value
}

func (c CreateAutoExpireScanTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAutoExpireScanTaskResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateAutoExpireScanTaskResponseScanType struct {
	value string
}

type CreateAutoExpireScanTaskResponseScanTypeEnum struct {
	MANUAL CreateAutoExpireScanTaskResponseScanType
	AUTO   CreateAutoExpireScanTaskResponseScanType
}

func GetCreateAutoExpireScanTaskResponseScanTypeEnum() CreateAutoExpireScanTaskResponseScanTypeEnum {
	return CreateAutoExpireScanTaskResponseScanTypeEnum{
		MANUAL: CreateAutoExpireScanTaskResponseScanType{
			value: "manual",
		},
		AUTO: CreateAutoExpireScanTaskResponseScanType{
			value: "auto",
		},
	}
}

func (c CreateAutoExpireScanTaskResponseScanType) Value() string {
	return c.value
}

func (c CreateAutoExpireScanTaskResponseScanType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAutoExpireScanTaskResponseScanType) UnmarshalJSON(b []byte) error {
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
