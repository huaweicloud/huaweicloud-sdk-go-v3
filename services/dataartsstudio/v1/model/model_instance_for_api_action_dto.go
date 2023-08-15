package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InstanceForApiActionDto struct {

	// 集群编号
	InstanceId *string `json:"instance_id,omitempty"`

	// 集群类型
	InstanceType *InstanceForApiActionDtoInstanceType `json:"instance_type,omitempty"`

	// 集群名称
	Name *string `json:"name,omitempty"`

	// api操作
	Action *InstanceForApiActionDtoAction `json:"action,omitempty"`

	// 校验结果
	Result *bool `json:"result,omitempty"`

	// 校验失败的原因
	Cause *InstanceForApiActionDtoCause `json:"cause,omitempty"`

	// api状态
	ApiStatus *InstanceForApiActionDtoApiStatus `json:"api_status,omitempty"`

	// api调试状态
	ApiDebug *InstanceForApiActionDtoApiDebug `json:"api_debug,omitempty"`
}

func (o InstanceForApiActionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceForApiActionDto struct{}"
	}

	return strings.Join([]string{"InstanceForApiActionDto", string(data)}, " ")
}

type InstanceForApiActionDtoInstanceType struct {
	value string
}

type InstanceForApiActionDtoInstanceTypeEnum struct {
	DLM       InstanceForApiActionDtoInstanceType
	APIG      InstanceForApiActionDtoInstanceType
	APIGW     InstanceForApiActionDtoInstanceType
	ROMA_APIC InstanceForApiActionDtoInstanceType
}

func GetInstanceForApiActionDtoInstanceTypeEnum() InstanceForApiActionDtoInstanceTypeEnum {
	return InstanceForApiActionDtoInstanceTypeEnum{
		DLM: InstanceForApiActionDtoInstanceType{
			value: "DLM",
		},
		APIG: InstanceForApiActionDtoInstanceType{
			value: "APIG",
		},
		APIGW: InstanceForApiActionDtoInstanceType{
			value: "APIGW",
		},
		ROMA_APIC: InstanceForApiActionDtoInstanceType{
			value: "ROMA_APIC",
		},
	}
}

func (c InstanceForApiActionDtoInstanceType) Value() string {
	return c.value
}

func (c InstanceForApiActionDtoInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceForApiActionDtoInstanceType) UnmarshalJSON(b []byte) error {
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

type InstanceForApiActionDtoAction struct {
	value string
}

type InstanceForApiActionDtoActionEnum struct {
	PUBLISH   InstanceForApiActionDtoAction
	UNPUBLISH InstanceForApiActionDtoAction
	STOP      InstanceForApiActionDtoAction
	RECOVER   InstanceForApiActionDtoAction
	WHITELIST InstanceForApiActionDtoAction
	AUTHORIZE InstanceForApiActionDtoAction
}

func GetInstanceForApiActionDtoActionEnum() InstanceForApiActionDtoActionEnum {
	return InstanceForApiActionDtoActionEnum{
		PUBLISH: InstanceForApiActionDtoAction{
			value: "PUBLISH",
		},
		UNPUBLISH: InstanceForApiActionDtoAction{
			value: "UNPUBLISH",
		},
		STOP: InstanceForApiActionDtoAction{
			value: "STOP",
		},
		RECOVER: InstanceForApiActionDtoAction{
			value: "RECOVER",
		},
		WHITELIST: InstanceForApiActionDtoAction{
			value: "WHITELIST",
		},
		AUTHORIZE: InstanceForApiActionDtoAction{
			value: "AUTHORIZE",
		},
	}
}

func (c InstanceForApiActionDtoAction) Value() string {
	return c.value
}

func (c InstanceForApiActionDtoAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceForApiActionDtoAction) UnmarshalJSON(b []byte) error {
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

type InstanceForApiActionDtoCause struct {
	value string
}

type InstanceForApiActionDtoCauseEnum struct {
	API_STATUS_ERROR InstanceForApiActionDtoCause
	API_DEBUG_ERROR  InstanceForApiActionDtoCause
	TYPE_MISMATCH    InstanceForApiActionDtoCause
}

func GetInstanceForApiActionDtoCauseEnum() InstanceForApiActionDtoCauseEnum {
	return InstanceForApiActionDtoCauseEnum{
		API_STATUS_ERROR: InstanceForApiActionDtoCause{
			value: "API_STATUS_ERROR",
		},
		API_DEBUG_ERROR: InstanceForApiActionDtoCause{
			value: "API_DEBUG_ERROR",
		},
		TYPE_MISMATCH: InstanceForApiActionDtoCause{
			value: "TYPE_MISMATCH",
		},
	}
}

func (c InstanceForApiActionDtoCause) Value() string {
	return c.value
}

func (c InstanceForApiActionDtoCause) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceForApiActionDtoCause) UnmarshalJSON(b []byte) error {
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

type InstanceForApiActionDtoApiStatus struct {
	value string
}

type InstanceForApiActionDtoApiStatusEnum struct {
	API_STATUS_CREATED             InstanceForApiActionDtoApiStatus
	API_STATUS_PUBLISH_WAIT_REVIEW InstanceForApiActionDtoApiStatus
	API_STATUS_PUBLISH_REJECT      InstanceForApiActionDtoApiStatus
	API_STATUS_PUBLISHED           InstanceForApiActionDtoApiStatus
	API_STATUS_WAITING_STOP        InstanceForApiActionDtoApiStatus
	API_STATUS_STOPPED             InstanceForApiActionDtoApiStatus
	API_STATUS_RECOVER_WAIT_REVIEW InstanceForApiActionDtoApiStatus
	API_STATUS_WAITING_OFFLINE     InstanceForApiActionDtoApiStatus
	API_STATUS_OFFLINE             InstanceForApiActionDtoApiStatus
}

func GetInstanceForApiActionDtoApiStatusEnum() InstanceForApiActionDtoApiStatusEnum {
	return InstanceForApiActionDtoApiStatusEnum{
		API_STATUS_CREATED: InstanceForApiActionDtoApiStatus{
			value: "API_STATUS_CREATED",
		},
		API_STATUS_PUBLISH_WAIT_REVIEW: InstanceForApiActionDtoApiStatus{
			value: "API_STATUS_PUBLISH_WAIT_REVIEW",
		},
		API_STATUS_PUBLISH_REJECT: InstanceForApiActionDtoApiStatus{
			value: "API_STATUS_PUBLISH_REJECT",
		},
		API_STATUS_PUBLISHED: InstanceForApiActionDtoApiStatus{
			value: "API_STATUS_PUBLISHED",
		},
		API_STATUS_WAITING_STOP: InstanceForApiActionDtoApiStatus{
			value: "API_STATUS_WAITING_STOP",
		},
		API_STATUS_STOPPED: InstanceForApiActionDtoApiStatus{
			value: "API_STATUS_STOPPED",
		},
		API_STATUS_RECOVER_WAIT_REVIEW: InstanceForApiActionDtoApiStatus{
			value: "API_STATUS_RECOVER_WAIT_REVIEW",
		},
		API_STATUS_WAITING_OFFLINE: InstanceForApiActionDtoApiStatus{
			value: "API_STATUS_WAITING_OFFLINE",
		},
		API_STATUS_OFFLINE: InstanceForApiActionDtoApiStatus{
			value: "API_STATUS_OFFLINE",
		},
	}
}

func (c InstanceForApiActionDtoApiStatus) Value() string {
	return c.value
}

func (c InstanceForApiActionDtoApiStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceForApiActionDtoApiStatus) UnmarshalJSON(b []byte) error {
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

type InstanceForApiActionDtoApiDebug struct {
	value string
}

type InstanceForApiActionDtoApiDebugEnum struct {
	API_DEBUG_WAITING InstanceForApiActionDtoApiDebug
	API_DEBUG_FAILED  InstanceForApiActionDtoApiDebug
	API_DEBUG_SUCCESS InstanceForApiActionDtoApiDebug
}

func GetInstanceForApiActionDtoApiDebugEnum() InstanceForApiActionDtoApiDebugEnum {
	return InstanceForApiActionDtoApiDebugEnum{
		API_DEBUG_WAITING: InstanceForApiActionDtoApiDebug{
			value: "API_DEBUG_WAITING",
		},
		API_DEBUG_FAILED: InstanceForApiActionDtoApiDebug{
			value: "API_DEBUG_FAILED",
		},
		API_DEBUG_SUCCESS: InstanceForApiActionDtoApiDebug{
			value: "API_DEBUG_SUCCESS",
		},
	}
}

func (c InstanceForApiActionDtoApiDebug) Value() string {
	return c.value
}

func (c InstanceForApiActionDtoApiDebug) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceForApiActionDtoApiDebug) UnmarshalJSON(b []byte) error {
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
