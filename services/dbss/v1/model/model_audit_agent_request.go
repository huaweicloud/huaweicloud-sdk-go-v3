package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AuditAgentRequest struct {

	// 数据库ID, 可在查询数据库列表接口的ID字段获取。
	DbId string `json:"db_id"`

	// 模式 - 0：创建agent - 1：选择已有agent
	Mode AuditAgentRequestMode `json:"mode"`

	// 选择已有agent时必输
	AgentId *string `json:"agent_id,omitempty"`

	// agent类型 - APP：应用端 - DB：数据库端
	AgentType AuditAgentRequestAgentType `json:"agent_type"`

	// agent OS类型: - LINUX64_X86 - LINUX64_ARM - WINDOWS64
	AgentOs AuditAgentRequestAgentOs `json:"agent_os"`

	// agent IP，安装节点类型为应用端时必输。
	AgentIp *string `json:"agent_ip,omitempty"`

	// agent审计网卡名称
	AgentNic *string `json:"agent_nic,omitempty"`

	// CPU阈值
	CpuThreshold *int32 `json:"cpu_threshold,omitempty"`

	// 内存阈值
	MemThreshold *int32 `json:"mem_threshold,omitempty"`
}

func (o AuditAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditAgentRequest struct{}"
	}

	return strings.Join([]string{"AuditAgentRequest", string(data)}, " ")
}

type AuditAgentRequestMode struct {
	value int32
}

type AuditAgentRequestModeEnum struct {
	E_0 AuditAgentRequestMode
	E_1 AuditAgentRequestMode
}

func GetAuditAgentRequestModeEnum() AuditAgentRequestModeEnum {
	return AuditAgentRequestModeEnum{
		E_0: AuditAgentRequestMode{
			value: 0,
		}, E_1: AuditAgentRequestMode{
			value: 1,
		},
	}
}

func (c AuditAgentRequestMode) Value() int32 {
	return c.value
}

func (c AuditAgentRequestMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuditAgentRequestMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type AuditAgentRequestAgentType struct {
	value string
}

type AuditAgentRequestAgentTypeEnum struct {
	APP AuditAgentRequestAgentType
	DB  AuditAgentRequestAgentType
}

func GetAuditAgentRequestAgentTypeEnum() AuditAgentRequestAgentTypeEnum {
	return AuditAgentRequestAgentTypeEnum{
		APP: AuditAgentRequestAgentType{
			value: "APP",
		},
		DB: AuditAgentRequestAgentType{
			value: "DB",
		},
	}
}

func (c AuditAgentRequestAgentType) Value() string {
	return c.value
}

func (c AuditAgentRequestAgentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuditAgentRequestAgentType) UnmarshalJSON(b []byte) error {
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

type AuditAgentRequestAgentOs struct {
	value string
}

type AuditAgentRequestAgentOsEnum struct {
	LINUX64_X86 AuditAgentRequestAgentOs
	LINUX64_ARM AuditAgentRequestAgentOs
	WINDOWS64   AuditAgentRequestAgentOs
}

func GetAuditAgentRequestAgentOsEnum() AuditAgentRequestAgentOsEnum {
	return AuditAgentRequestAgentOsEnum{
		LINUX64_X86: AuditAgentRequestAgentOs{
			value: "LINUX64_X86",
		},
		LINUX64_ARM: AuditAgentRequestAgentOs{
			value: "LINUX64_ARM",
		},
		WINDOWS64: AuditAgentRequestAgentOs{
			value: "WINDOWS64",
		},
	}
}

func (c AuditAgentRequestAgentOs) Value() string {
	return c.value
}

func (c AuditAgentRequestAgentOs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuditAgentRequestAgentOs) UnmarshalJSON(b []byte) error {
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
