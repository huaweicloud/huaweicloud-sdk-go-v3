package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAutopilotClusterConfigRequest Request Object
type ShowAutopilotClusterConfigRequest struct {

	// **参数解释**： 组件类型，不填写则查询全部类型。  **约束限制**： 合法取值为control，audit，system-addon  **取值范围**： - control: 控制面组件日志。 - audit: 控制面审计日志。 - system-addon: 系统插件日志。  **默认取值**： 无
	Type *ShowAutopilotClusterConfigRequestType `json:"type,omitempty"`

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`
}

func (o ShowAutopilotClusterConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutopilotClusterConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowAutopilotClusterConfigRequest", string(data)}, " ")
}

type ShowAutopilotClusterConfigRequestType struct {
	value string
}

type ShowAutopilotClusterConfigRequestTypeEnum struct {
	CONTROL      ShowAutopilotClusterConfigRequestType
	AUDIT        ShowAutopilotClusterConfigRequestType
	SYSTEM_ADDON ShowAutopilotClusterConfigRequestType
}

func GetShowAutopilotClusterConfigRequestTypeEnum() ShowAutopilotClusterConfigRequestTypeEnum {
	return ShowAutopilotClusterConfigRequestTypeEnum{
		CONTROL: ShowAutopilotClusterConfigRequestType{
			value: "control",
		},
		AUDIT: ShowAutopilotClusterConfigRequestType{
			value: "audit",
		},
		SYSTEM_ADDON: ShowAutopilotClusterConfigRequestType{
			value: "system-addon",
		},
	}
}

func (c ShowAutopilotClusterConfigRequestType) Value() string {
	return c.value
}

func (c ShowAutopilotClusterConfigRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAutopilotClusterConfigRequestType) UnmarshalJSON(b []byte) error {
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
