package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// SwitchoverTestRecord 倒换测试记录信息
type SwitchoverTestRecord struct {

	// 倒换测试记录的唯一标识
	Id *string `json:"id,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 倒换测试的资源对象ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 倒换测试的资源对象类型
	ResourceType *SwitchoverTestRecordResourceType `json:"resource_type,omitempty"`

	// shutdown, undo_shutdown表示倒换测试操作类型
	Operation *SwitchoverTestRecordOperation `json:"operation,omitempty"`

	// 倒换测试操作的开始时间
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 倒换测试操作的结束时间
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 倒换测试状态记录 STARTING: 初始状态 INPROGRESS: 配置下发中 COMPLETE: 配置下发完成 ERROR: 配置下发失败
	OperateStatus *SwitchoverTestRecordOperateStatus `json:"operate_status,omitempty"`
}

func (o SwitchoverTestRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchoverTestRecord struct{}"
	}

	return strings.Join([]string{"SwitchoverTestRecord", string(data)}, " ")
}

type SwitchoverTestRecordResourceType struct {
	value string
}

type SwitchoverTestRecordResourceTypeEnum struct {
	VIRTUAL_INTERFACE SwitchoverTestRecordResourceType
}

func GetSwitchoverTestRecordResourceTypeEnum() SwitchoverTestRecordResourceTypeEnum {
	return SwitchoverTestRecordResourceTypeEnum{
		VIRTUAL_INTERFACE: SwitchoverTestRecordResourceType{
			value: "virtual_interface",
		},
	}
}

func (c SwitchoverTestRecordResourceType) Value() string {
	return c.value
}

func (c SwitchoverTestRecordResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchoverTestRecordResourceType) UnmarshalJSON(b []byte) error {
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

type SwitchoverTestRecordOperation struct {
	value string
}

type SwitchoverTestRecordOperationEnum struct {
	SHUTDOWN      SwitchoverTestRecordOperation
	UNDO_SHUTDOWN SwitchoverTestRecordOperation
}

func GetSwitchoverTestRecordOperationEnum() SwitchoverTestRecordOperationEnum {
	return SwitchoverTestRecordOperationEnum{
		SHUTDOWN: SwitchoverTestRecordOperation{
			value: "shutdown",
		},
		UNDO_SHUTDOWN: SwitchoverTestRecordOperation{
			value: "undo_shutdown",
		},
	}
}

func (c SwitchoverTestRecordOperation) Value() string {
	return c.value
}

func (c SwitchoverTestRecordOperation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchoverTestRecordOperation) UnmarshalJSON(b []byte) error {
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

type SwitchoverTestRecordOperateStatus struct {
	value string
}

type SwitchoverTestRecordOperateStatusEnum struct {
	STARTING   SwitchoverTestRecordOperateStatus
	INPROGRESS SwitchoverTestRecordOperateStatus
	COMPLETE   SwitchoverTestRecordOperateStatus
	ERROR      SwitchoverTestRecordOperateStatus
}

func GetSwitchoverTestRecordOperateStatusEnum() SwitchoverTestRecordOperateStatusEnum {
	return SwitchoverTestRecordOperateStatusEnum{
		STARTING: SwitchoverTestRecordOperateStatus{
			value: "STARTING",
		},
		INPROGRESS: SwitchoverTestRecordOperateStatus{
			value: "INPROGRESS",
		},
		COMPLETE: SwitchoverTestRecordOperateStatus{
			value: "COMPLETE",
		},
		ERROR: SwitchoverTestRecordOperateStatus{
			value: "ERROR",
		},
	}
}

func (c SwitchoverTestRecordOperateStatus) Value() string {
	return c.value
}

func (c SwitchoverTestRecordOperateStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchoverTestRecordOperateStatus) UnmarshalJSON(b []byte) error {
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
