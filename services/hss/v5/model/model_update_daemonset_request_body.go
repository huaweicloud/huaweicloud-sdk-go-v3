package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateDaemonsetRequestBody struct {

	// agent版本
	AgentVersion *string `json:"agent_version,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 开启agent自动升级
	AutoUpgrade *bool `json:"auto_upgrade,omitempty"`

	// 容器运行时配置
	RuntimeInfo *[]RuntimeRequestBody `json:"runtime_info,omitempty"`

	ScheduleInfo *UpdateDaemonsetRequestBodyScheduleInfo `json:"schedule_info,omitempty"`

	// 调用服务，默认hss，cce集成防护调用场景使用:   - hss：hss服务    - cce：cce服务
	InvokedService *string `json:"invoked_service,omitempty"`

	// 付费模式，cce集成防护调用场景使用:   - on_demand:按需    - free_security_check:免费安全体检
	ChargingMode *UpdateDaemonsetRequestBodyChargingMode `json:"charging_mode,omitempty"`

	// cce防护类型，cce集成防护调用场景使用:   - cluster_level:集群级别防护    - node_level:节点级别防护
	CceProtectionType *UpdateDaemonsetRequestBodyCceProtectionType `json:"cce_protection_type,omitempty"`

	// 优先使用包周期配额，cce集成防护调用场景使用，默认false
	PreferPacketCycle *bool `json:"prefer_packet_cycle,omitempty"`
}

func (o UpdateDaemonsetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDaemonsetRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDaemonsetRequestBody", string(data)}, " ")
}

type UpdateDaemonsetRequestBodyChargingMode struct {
	value string
}

type UpdateDaemonsetRequestBodyChargingModeEnum struct {
	ON_DEMAND           UpdateDaemonsetRequestBodyChargingMode
	FREE_SECURITY_CHECK UpdateDaemonsetRequestBodyChargingMode
}

func GetUpdateDaemonsetRequestBodyChargingModeEnum() UpdateDaemonsetRequestBodyChargingModeEnum {
	return UpdateDaemonsetRequestBodyChargingModeEnum{
		ON_DEMAND: UpdateDaemonsetRequestBodyChargingMode{
			value: "on_demand",
		},
		FREE_SECURITY_CHECK: UpdateDaemonsetRequestBodyChargingMode{
			value: "free_security_check",
		},
	}
}

func (c UpdateDaemonsetRequestBodyChargingMode) Value() string {
	return c.value
}

func (c UpdateDaemonsetRequestBodyChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDaemonsetRequestBodyChargingMode) UnmarshalJSON(b []byte) error {
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

type UpdateDaemonsetRequestBodyCceProtectionType struct {
	value string
}

type UpdateDaemonsetRequestBodyCceProtectionTypeEnum struct {
	CLUSTER_LEVEL UpdateDaemonsetRequestBodyCceProtectionType
	NODE_LEVEL    UpdateDaemonsetRequestBodyCceProtectionType
}

func GetUpdateDaemonsetRequestBodyCceProtectionTypeEnum() UpdateDaemonsetRequestBodyCceProtectionTypeEnum {
	return UpdateDaemonsetRequestBodyCceProtectionTypeEnum{
		CLUSTER_LEVEL: UpdateDaemonsetRequestBodyCceProtectionType{
			value: "cluster_level",
		},
		NODE_LEVEL: UpdateDaemonsetRequestBodyCceProtectionType{
			value: "node_level",
		},
	}
}

func (c UpdateDaemonsetRequestBodyCceProtectionType) Value() string {
	return c.value
}

func (c UpdateDaemonsetRequestBodyCceProtectionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDaemonsetRequestBodyCceProtectionType) UnmarshalJSON(b []byte) error {
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
