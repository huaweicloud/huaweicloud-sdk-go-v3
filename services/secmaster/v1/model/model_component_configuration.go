package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ComponentConfiguration 插件配置定义简介
type ComponentConfiguration struct {

	// 组件id.
	ComponentId *string `json:"component_id,omitempty"`

	// 组件名称
	ComponentName *string `json:"component_name,omitempty"`

	// **参数解释**: 节点配置状态 - UN_SAVED 待保存 - SAVE_AND_UN_APPLY 待应用 - MOVE_AND_UN_APPLY 待移除 - APPLYING 应用中 - FAIL_APPLY 应用失败 - APPLIED 应用完成  **约束限制** 不涉及 **取值范围**: - UN_SAVED - SAVE_AND_UN_APPLY - MOVE_AND_UN_APPLY - APPLYING - FAIL_APPLY - APPLIED  **默认值** 不涉及
	ConfigStatus *ComponentConfigurationConfigStatus `json:"config_status,omitempty"`

	// 毫秒时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 部署失败的消息
	FailDeployMessage *string `json:"fail_deploy_message,omitempty"`

	// IP地址
	IpAddress *string `json:"ip_address,omitempty"`

	// 组件配置参数列表
	List *[]ComponentConfigurationParam `json:"list,omitempty"`

	Monitor *Monitor `json:"monitor,omitempty"`

	NodeApplyFailEnum *NodeApplyFailEnum `json:"node_apply_fail_enum,omitempty"`

	NodeExpansion *IsapNodeExpansion `json:"node_expansion,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 节点名称
	NodeName *string `json:"node_name,omitempty"`

	// IP地址
	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	// 地区
	Region *string `json:"region,omitempty"`

	// 规范
	Specification *string `json:"specification,omitempty"`

	// VPC端点地址
	VpcEndpointAddress *string `json:"vpc_endpoint_address,omitempty"`

	// VPC端点ID
	VpcEndpointId *string `json:"vpc_endpoint_id,omitempty"`
}

func (o ComponentConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentConfiguration struct{}"
	}

	return strings.Join([]string{"ComponentConfiguration", string(data)}, " ")
}

type ComponentConfigurationConfigStatus struct {
	value string
}

type ComponentConfigurationConfigStatusEnum struct {
	UN_SAVED          ComponentConfigurationConfigStatus
	SAVE_AND_UN_APPLY ComponentConfigurationConfigStatus
	MOVE_AND_UN_APPLY ComponentConfigurationConfigStatus
	APPLYING          ComponentConfigurationConfigStatus
	FAIL_APPLY        ComponentConfigurationConfigStatus
	APPLIED           ComponentConfigurationConfigStatus
}

func GetComponentConfigurationConfigStatusEnum() ComponentConfigurationConfigStatusEnum {
	return ComponentConfigurationConfigStatusEnum{
		UN_SAVED: ComponentConfigurationConfigStatus{
			value: "UN_SAVED",
		},
		SAVE_AND_UN_APPLY: ComponentConfigurationConfigStatus{
			value: "SAVE_AND_UN_APPLY",
		},
		MOVE_AND_UN_APPLY: ComponentConfigurationConfigStatus{
			value: "MOVE_AND_UN_APPLY",
		},
		APPLYING: ComponentConfigurationConfigStatus{
			value: "APPLYING",
		},
		FAIL_APPLY: ComponentConfigurationConfigStatus{
			value: "FAIL_APPLY",
		},
		APPLIED: ComponentConfigurationConfigStatus{
			value: "APPLIED",
		},
	}
}

func (c ComponentConfigurationConfigStatus) Value() string {
	return c.value
}

func (c ComponentConfigurationConfigStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentConfigurationConfigStatus) UnmarshalJSON(b []byte) error {
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
