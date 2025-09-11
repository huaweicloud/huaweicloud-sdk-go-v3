package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InvocationInfo struct {

	// **参数解释**: 任务id **取值范围**: 以字母或数字开头，后续可以包含字母、数字、下划线（_）或连字符（-），长度至少为1，且不允许出现其他字符
	InvocationId *string `json:"invocation_id,omitempty"`

	// **参数解释**: 主机id **取值范围**: 1到64个字符的字符串，且只包含字母、数字和连字符
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 主机名称 **取值范围**: 字符串长度范围为[1,128]
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**: 主机类型，仅支持ECS弹性云服务器和BMS裸金属服务器 **取值范围**: - ECS: 弹性云服务器 - BMS：裸金属服务器
	InstanceType *InvocationInfoInstanceType `json:"instance_type,omitempty"`

	// **参数解释**: 内网ip列表 **取值范围**: 返回数组长度为[0,10]，数组内元素格式为：长度为1到15个字符的字符串，其中每个字符可以是数字（0-9）或任意其他单个字符
	IntranetIps *[]string `json:"intranet_ips,omitempty"`

	// **参数解释**: 弹性公网ip列表 **取值范围**: 返回数组长度为[0,10]，数组内元素格式为：长度为1到15个字符的字符串，其中每个字符可以是数字（0-9）或任意其他单个字符
	ElasticIps *[]string `json:"elastic_ips,omitempty"`

	// **参数解释**: 任务类型 **取值范围**: - INSTALL：安装 - UPDATE：升级 - ROLLBACK：回滚 - RETRY：重试
	InvocationType *InvocationInfoInvocationType `json:"invocation_type,omitempty"`

	// **参数解释**: 任务状态 **取值范围**: - PENDING：待执行 - RUNNING：运行中 - TIMEOUT：超时 - FAILED：失败 - SUCCEEDED：成功 - CANCELED：取消 - ROLLBACKED：已回退
	InvocationStatus *InvocationInfoInvocationStatus `json:"invocation_status,omitempty"`

	// **参数解释**: 任务对象, 支持telescope监控 **取值范围**: - telescope: 主机监控插件telescope
	InvocationTarget *InvocationInfoInvocationTarget `json:"invocation_target,omitempty"`

	// **参数解释**: 任务创建时间 **取值范围**: 数字范围为[1111111111111,9999999999999]
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**: 任务更新时间 **取值范围**: 数字范围为[1111111111111,9999999999999]
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**: 当前版本 **取值范围**: 字符串长度范围为[1,64]
	CurrentVersion *string `json:"current_version,omitempty"`

	// **参数解释**: 目标版本 **取值范围**: 字符串长度范围为[1,64]
	TargetVersion *string `json:"target_version,omitempty"`

	// **参数解释**: 任务执行结果信息 **取值范围**: 字符串长度范围为[1,5000]
	ResultMsg *string `json:"result_msg,omitempty"`
}

func (o InvocationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvocationInfo struct{}"
	}

	return strings.Join([]string{"InvocationInfo", string(data)}, " ")
}

type InvocationInfoInstanceType struct {
	value string
}

type InvocationInfoInstanceTypeEnum struct {
	ECS InvocationInfoInstanceType
	BMS InvocationInfoInstanceType
}

func GetInvocationInfoInstanceTypeEnum() InvocationInfoInstanceTypeEnum {
	return InvocationInfoInstanceTypeEnum{
		ECS: InvocationInfoInstanceType{
			value: "ECS",
		},
		BMS: InvocationInfoInstanceType{
			value: "BMS",
		},
	}
}

func (c InvocationInfoInstanceType) Value() string {
	return c.value
}

func (c InvocationInfoInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InvocationInfoInstanceType) UnmarshalJSON(b []byte) error {
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

type InvocationInfoInvocationType struct {
	value string
}

type InvocationInfoInvocationTypeEnum struct {
	INSTALL  InvocationInfoInvocationType
	UPDATE   InvocationInfoInvocationType
	ROLLBACK InvocationInfoInvocationType
	RETRY    InvocationInfoInvocationType
}

func GetInvocationInfoInvocationTypeEnum() InvocationInfoInvocationTypeEnum {
	return InvocationInfoInvocationTypeEnum{
		INSTALL: InvocationInfoInvocationType{
			value: "INSTALL",
		},
		UPDATE: InvocationInfoInvocationType{
			value: "UPDATE",
		},
		ROLLBACK: InvocationInfoInvocationType{
			value: "ROLLBACK",
		},
		RETRY: InvocationInfoInvocationType{
			value: "RETRY",
		},
	}
}

func (c InvocationInfoInvocationType) Value() string {
	return c.value
}

func (c InvocationInfoInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InvocationInfoInvocationType) UnmarshalJSON(b []byte) error {
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

type InvocationInfoInvocationStatus struct {
	value string
}

type InvocationInfoInvocationStatusEnum struct {
	PENDING    InvocationInfoInvocationStatus
	RUNNING    InvocationInfoInvocationStatus
	TIMEOUT    InvocationInfoInvocationStatus
	FAILED     InvocationInfoInvocationStatus
	SUCCEEDED  InvocationInfoInvocationStatus
	CANCELED   InvocationInfoInvocationStatus
	ROLLBACKED InvocationInfoInvocationStatus
}

func GetInvocationInfoInvocationStatusEnum() InvocationInfoInvocationStatusEnum {
	return InvocationInfoInvocationStatusEnum{
		PENDING: InvocationInfoInvocationStatus{
			value: "PENDING",
		},
		RUNNING: InvocationInfoInvocationStatus{
			value: "RUNNING",
		},
		TIMEOUT: InvocationInfoInvocationStatus{
			value: "TIMEOUT",
		},
		FAILED: InvocationInfoInvocationStatus{
			value: "FAILED",
		},
		SUCCEEDED: InvocationInfoInvocationStatus{
			value: "SUCCEEDED",
		},
		CANCELED: InvocationInfoInvocationStatus{
			value: "CANCELED",
		},
		ROLLBACKED: InvocationInfoInvocationStatus{
			value: "ROLLBACKED",
		},
	}
}

func (c InvocationInfoInvocationStatus) Value() string {
	return c.value
}

func (c InvocationInfoInvocationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InvocationInfoInvocationStatus) UnmarshalJSON(b []byte) error {
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

type InvocationInfoInvocationTarget struct {
	value string
}

type InvocationInfoInvocationTargetEnum struct {
	TELESCOPE InvocationInfoInvocationTarget
}

func GetInvocationInfoInvocationTargetEnum() InvocationInfoInvocationTargetEnum {
	return InvocationInfoInvocationTargetEnum{
		TELESCOPE: InvocationInfoInvocationTarget{
			value: "telescope",
		},
	}
}

func (c InvocationInfoInvocationTarget) Value() string {
	return c.value
}

func (c InvocationInfoInvocationTarget) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InvocationInfoInvocationTarget) UnmarshalJSON(b []byte) error {
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
