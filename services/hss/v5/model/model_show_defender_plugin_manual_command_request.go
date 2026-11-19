package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDefenderPluginManualCommandRequest Request Object
type ShowDefenderPluginManualCommandRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 操作类型 **约束限制**： 必填 **取值范围**： - install：安装 - uninstall：卸载  **默认取值**： 不涉及
	Opt ShowDefenderPluginManualCommandRequestOpt `json:"opt"`

	// **参数解释**： 网络类型，用于指定返回的命令中使用的下载地址和服务端地址的网络类型 **约束限制**： 非必填 **取值范围**： - public：公网 - internal：内网  **默认取值**： internal
	Network *ShowDefenderPluginManualCommandRequestNetwork `json:"network,omitempty"`

	// **参数解释**： 智能体类型，用于区分OpenClaw和自研Claw **约束限制**： 必填 **取值范围**： - OpenClaw：OpenClaw - SelfDevelopedClaw：自研claw  **默认取值**： OpenClaw
	AgentType ShowDefenderPluginManualCommandRequestAgentType `json:"agent_type"`
}

func (o ShowDefenderPluginManualCommandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefenderPluginManualCommandRequest struct{}"
	}

	return strings.Join([]string{"ShowDefenderPluginManualCommandRequest", string(data)}, " ")
}

type ShowDefenderPluginManualCommandRequestOpt struct {
	value string
}

type ShowDefenderPluginManualCommandRequestOptEnum struct {
	INSTALL   ShowDefenderPluginManualCommandRequestOpt
	UNINSTALL ShowDefenderPluginManualCommandRequestOpt
}

func GetShowDefenderPluginManualCommandRequestOptEnum() ShowDefenderPluginManualCommandRequestOptEnum {
	return ShowDefenderPluginManualCommandRequestOptEnum{
		INSTALL: ShowDefenderPluginManualCommandRequestOpt{
			value: "install",
		},
		UNINSTALL: ShowDefenderPluginManualCommandRequestOpt{
			value: "uninstall",
		},
	}
}

func (c ShowDefenderPluginManualCommandRequestOpt) Value() string {
	return c.value
}

func (c ShowDefenderPluginManualCommandRequestOpt) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDefenderPluginManualCommandRequestOpt) UnmarshalJSON(b []byte) error {
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

type ShowDefenderPluginManualCommandRequestNetwork struct {
	value string
}

type ShowDefenderPluginManualCommandRequestNetworkEnum struct {
	PUBLIC   ShowDefenderPluginManualCommandRequestNetwork
	INTERNAL ShowDefenderPluginManualCommandRequestNetwork
}

func GetShowDefenderPluginManualCommandRequestNetworkEnum() ShowDefenderPluginManualCommandRequestNetworkEnum {
	return ShowDefenderPluginManualCommandRequestNetworkEnum{
		PUBLIC: ShowDefenderPluginManualCommandRequestNetwork{
			value: "public",
		},
		INTERNAL: ShowDefenderPluginManualCommandRequestNetwork{
			value: "internal",
		},
	}
}

func (c ShowDefenderPluginManualCommandRequestNetwork) Value() string {
	return c.value
}

func (c ShowDefenderPluginManualCommandRequestNetwork) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDefenderPluginManualCommandRequestNetwork) UnmarshalJSON(b []byte) error {
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

type ShowDefenderPluginManualCommandRequestAgentType struct {
	value string
}

type ShowDefenderPluginManualCommandRequestAgentTypeEnum struct {
	OPEN_CLAW           ShowDefenderPluginManualCommandRequestAgentType
	SELF_DEVELOPED_CLAW ShowDefenderPluginManualCommandRequestAgentType
}

func GetShowDefenderPluginManualCommandRequestAgentTypeEnum() ShowDefenderPluginManualCommandRequestAgentTypeEnum {
	return ShowDefenderPluginManualCommandRequestAgentTypeEnum{
		OPEN_CLAW: ShowDefenderPluginManualCommandRequestAgentType{
			value: "OpenClaw",
		},
		SELF_DEVELOPED_CLAW: ShowDefenderPluginManualCommandRequestAgentType{
			value: "SelfDevelopedClaw",
		},
	}
}

func (c ShowDefenderPluginManualCommandRequestAgentType) Value() string {
	return c.value
}

func (c ShowDefenderPluginManualCommandRequestAgentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDefenderPluginManualCommandRequestAgentType) UnmarshalJSON(b []byte) error {
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
