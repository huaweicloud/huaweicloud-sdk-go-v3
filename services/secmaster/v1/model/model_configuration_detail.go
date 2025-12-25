package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ConfigurationDetail struct {

	// **参数解释**: 节点部署状态 - UN_SAVED 待保存 - SAVE_AND_UN_APPLY 待应用 - MOVE_AND_UN_APPLY 待移除 - APPLYING 应用中 - FAIL_APPLY 应用失败 - APPLIED 应用完成  **约束限制** 不涉及 **取值范围**: - UN_SAVED - SAVE_AND_UN_APPLY - MOVE_AND_UN_APPLY - APPLYING - FAIL_APPLY - APPLIED  **默认值** 不涉及
	ConfigStatus *ConfigurationDetailConfigStatus `json:"config_status,omitempty"`

	// 文件配置列表
	List []FileConfiguration `json:"list"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 节点名称
	NodeName *string `json:"node_name,omitempty"`

	// 规范
	Specification *string `json:"specification,omitempty"`
}

func (o ConfigurationDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationDetail struct{}"
	}

	return strings.Join([]string{"ConfigurationDetail", string(data)}, " ")
}

type ConfigurationDetailConfigStatus struct {
	value string
}

type ConfigurationDetailConfigStatusEnum struct {
	UN_SAVED          ConfigurationDetailConfigStatus
	SAVE_AND_UN_APPLY ConfigurationDetailConfigStatus
	MOVE_AND_UN_APPLY ConfigurationDetailConfigStatus
	APPLYING          ConfigurationDetailConfigStatus
	FAIL_APPLY        ConfigurationDetailConfigStatus
	APPLIED           ConfigurationDetailConfigStatus
}

func GetConfigurationDetailConfigStatusEnum() ConfigurationDetailConfigStatusEnum {
	return ConfigurationDetailConfigStatusEnum{
		UN_SAVED: ConfigurationDetailConfigStatus{
			value: "UN_SAVED",
		},
		SAVE_AND_UN_APPLY: ConfigurationDetailConfigStatus{
			value: "SAVE_AND_UN_APPLY",
		},
		MOVE_AND_UN_APPLY: ConfigurationDetailConfigStatus{
			value: "MOVE_AND_UN_APPLY",
		},
		APPLYING: ConfigurationDetailConfigStatus{
			value: "APPLYING",
		},
		FAIL_APPLY: ConfigurationDetailConfigStatus{
			value: "FAIL_APPLY",
		},
		APPLIED: ConfigurationDetailConfigStatus{
			value: "APPLIED",
		},
	}
}

func (c ConfigurationDetailConfigStatus) Value() string {
	return c.value
}

func (c ConfigurationDetailConfigStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigurationDetailConfigStatus) UnmarshalJSON(b []byte) error {
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
