package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SimpleKeyScanRecord struct {

	// 扫描ID
	Id *string `json:"id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 状态
	Status *SimpleKeyScanRecordStatus `json:"status,omitempty"`

	// 扫描类型
	ScanType *string `json:"scan_type,omitempty"`

	// 数量
	Num *int32 `json:"num,omitempty"`

	// 创建时间, 格式为: 2023-06-13T13:46:14.771Z
	CreatedAt *string `json:"created_at,omitempty"`

	// 开始时间, 格式为: 2023-06-13T13:46:14.771Z
	StartedAt *string `json:"started_at,omitempty"`

	// 完成时间, 格式为：2020-06-15T02:21:18.669Z
	FinishedAt *string `json:"finished_at,omitempty"`
}

func (o SimpleKeyScanRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleKeyScanRecord struct{}"
	}

	return strings.Join([]string{"SimpleKeyScanRecord", string(data)}, " ")
}

type SimpleKeyScanRecordStatus struct {
	value string
}

type SimpleKeyScanRecordStatusEnum struct {
	WAITING SimpleKeyScanRecordStatus
	RUNNING SimpleKeyScanRecordStatus
	SUCCESS SimpleKeyScanRecordStatus
	FAILED  SimpleKeyScanRecordStatus
}

func GetSimpleKeyScanRecordStatusEnum() SimpleKeyScanRecordStatusEnum {
	return SimpleKeyScanRecordStatusEnum{
		WAITING: SimpleKeyScanRecordStatus{
			value: "waiting",
		},
		RUNNING: SimpleKeyScanRecordStatus{
			value: "running",
		},
		SUCCESS: SimpleKeyScanRecordStatus{
			value: "success",
		},
		FAILED: SimpleKeyScanRecordStatus{
			value: "failed",
		},
	}
}

func (c SimpleKeyScanRecordStatus) Value() string {
	return c.value
}

func (c SimpleKeyScanRecordStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SimpleKeyScanRecordStatus) UnmarshalJSON(b []byte) error {
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
