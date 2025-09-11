package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAgentInvocationsRequest Request Object
type ListAgentInvocationsRequest struct {

	// **参数解释**: 主机id **约束限制**: 不涉及 **取值范围**: 1到64个字符的字符串，且只包含字母、数字和连字符 **默认取值**: 不涉及
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 主机类型，仅支持ECS弹性云服务器和BMS裸金属服务器 **约束限制**: 不涉及 **取值范围**: - ECS: 弹性云服务器 - BMS：裸金属服务器 **默认取值**: 不涉及
	InstanceType *ListAgentInvocationsRequestInstanceType `json:"instance_type,omitempty"`

	// **参数解释**: 任务id **约束限制**: 不涉及 **取值范围**: 以字母或数字开头，后续可包含字母、数字、下划线或连字符的字符串，长度至少为 1 **默认取值**: 不涉及
	InvocationId *string `json:"invocation_id,omitempty"`

	// **参数解释**: 任务类型, 仅包含：INSTALL安装, UPDATE升级, ROLLBACK回退，RETRY重试，SET_REMOTE_INSTALLER设置远程安装主机，REMOTE_INSTALL执行远程安装。 **约束限制**: 不涉及。 **取值范围**: - INSTALL：安装 - UPDATE：升级 - ROLLBACK：回退 - RETRY：重试 - SET_REMOTE_INSTALLER：设置远程安装主机 - REMOTE_INSTALL：执行远程安装 **默认取值**: 不涉及
	InvocationType *ListAgentInvocationsRequestInvocationType `json:"invocation_type,omitempty"`

	// **参数解释**: 任务对象, 支持telescope监控 **约束限制**: 不涉及。 **取值范围**: - telescope: 主机监控插件telescope **默认取值**: telescope。
	InvocationTarget *ListAgentInvocationsRequestInvocationTarget `json:"invocation_target,omitempty"`

	// **参数解释**: 分页偏移量 **约束限制**: 不涉及 **取值范围**: 数字范围为[0,9999999999999] **默认取值**: 0
	Offset *int64 `json:"offset,omitempty"`

	// **参数解释**: 分页大小。 **约束限制**: 不涉及。 **取值范围**: 数字范围为[1,100] **默认取值**: 100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAgentInvocationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentInvocationsRequest struct{}"
	}

	return strings.Join([]string{"ListAgentInvocationsRequest", string(data)}, " ")
}

type ListAgentInvocationsRequestInstanceType struct {
	value string
}

type ListAgentInvocationsRequestInstanceTypeEnum struct {
	ECS ListAgentInvocationsRequestInstanceType
	BMS ListAgentInvocationsRequestInstanceType
}

func GetListAgentInvocationsRequestInstanceTypeEnum() ListAgentInvocationsRequestInstanceTypeEnum {
	return ListAgentInvocationsRequestInstanceTypeEnum{
		ECS: ListAgentInvocationsRequestInstanceType{
			value: "ECS",
		},
		BMS: ListAgentInvocationsRequestInstanceType{
			value: "BMS",
		},
	}
}

func (c ListAgentInvocationsRequestInstanceType) Value() string {
	return c.value
}

func (c ListAgentInvocationsRequestInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgentInvocationsRequestInstanceType) UnmarshalJSON(b []byte) error {
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

type ListAgentInvocationsRequestInvocationType struct {
	value string
}

type ListAgentInvocationsRequestInvocationTypeEnum struct {
	INSTALL              ListAgentInvocationsRequestInvocationType
	UPDATE               ListAgentInvocationsRequestInvocationType
	ROLLBACK             ListAgentInvocationsRequestInvocationType
	RETRY                ListAgentInvocationsRequestInvocationType
	SET_REMOTE_INSTALLER ListAgentInvocationsRequestInvocationType
	REMOTE_INSTALL       ListAgentInvocationsRequestInvocationType
}

func GetListAgentInvocationsRequestInvocationTypeEnum() ListAgentInvocationsRequestInvocationTypeEnum {
	return ListAgentInvocationsRequestInvocationTypeEnum{
		INSTALL: ListAgentInvocationsRequestInvocationType{
			value: "INSTALL",
		},
		UPDATE: ListAgentInvocationsRequestInvocationType{
			value: "UPDATE",
		},
		ROLLBACK: ListAgentInvocationsRequestInvocationType{
			value: "ROLLBACK",
		},
		RETRY: ListAgentInvocationsRequestInvocationType{
			value: "RETRY",
		},
		SET_REMOTE_INSTALLER: ListAgentInvocationsRequestInvocationType{
			value: "SET_REMOTE_INSTALLER",
		},
		REMOTE_INSTALL: ListAgentInvocationsRequestInvocationType{
			value: "REMOTE_INSTALL",
		},
	}
}

func (c ListAgentInvocationsRequestInvocationType) Value() string {
	return c.value
}

func (c ListAgentInvocationsRequestInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgentInvocationsRequestInvocationType) UnmarshalJSON(b []byte) error {
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

type ListAgentInvocationsRequestInvocationTarget struct {
	value string
}

type ListAgentInvocationsRequestInvocationTargetEnum struct {
	TELESCOPE ListAgentInvocationsRequestInvocationTarget
}

func GetListAgentInvocationsRequestInvocationTargetEnum() ListAgentInvocationsRequestInvocationTargetEnum {
	return ListAgentInvocationsRequestInvocationTargetEnum{
		TELESCOPE: ListAgentInvocationsRequestInvocationTarget{
			value: "telescope",
		},
	}
}

func (c ListAgentInvocationsRequestInvocationTarget) Value() string {
	return c.value
}

func (c ListAgentInvocationsRequestInvocationTarget) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgentInvocationsRequestInvocationTarget) UnmarshalJSON(b []byte) error {
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
