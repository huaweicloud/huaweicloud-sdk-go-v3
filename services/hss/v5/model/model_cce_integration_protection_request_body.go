package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CceIntegrationProtectionRequestBody struct {

	// cce集群类型：   existing 存量集群   adding 新增集群
	ClusterType *CceIntegrationProtectionRequestBodyClusterType `json:"cluster_type,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName string `json:"cluster_name"`

	// 付费模式：   on_demand 按需   free_security_check 免费安全体检
	ChargingMode *CceIntegrationProtectionRequestBodyChargingMode `json:"charging_mode,omitempty"`

	// cce防护类型：   cluster_level 集群级别防护   node_level 节点级别防护
	CceProtectionType *CceIntegrationProtectionRequestBodyCceProtectionType `json:"cce_protection_type,omitempty"`

	// 优先使用包周期配额；默认false
	PreferPacketCycle *bool `json:"prefer_packet_cycle,omitempty"`
}

func (o CceIntegrationProtectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CceIntegrationProtectionRequestBody struct{}"
	}

	return strings.Join([]string{"CceIntegrationProtectionRequestBody", string(data)}, " ")
}

type CceIntegrationProtectionRequestBodyClusterType struct {
	value string
}

type CceIntegrationProtectionRequestBodyClusterTypeEnum struct {
	EXISTING CceIntegrationProtectionRequestBodyClusterType
	ADDING   CceIntegrationProtectionRequestBodyClusterType
}

func GetCceIntegrationProtectionRequestBodyClusterTypeEnum() CceIntegrationProtectionRequestBodyClusterTypeEnum {
	return CceIntegrationProtectionRequestBodyClusterTypeEnum{
		EXISTING: CceIntegrationProtectionRequestBodyClusterType{
			value: "existing",
		},
		ADDING: CceIntegrationProtectionRequestBodyClusterType{
			value: "adding",
		},
	}
}

func (c CceIntegrationProtectionRequestBodyClusterType) Value() string {
	return c.value
}

func (c CceIntegrationProtectionRequestBodyClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CceIntegrationProtectionRequestBodyClusterType) UnmarshalJSON(b []byte) error {
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

type CceIntegrationProtectionRequestBodyChargingMode struct {
	value string
}

type CceIntegrationProtectionRequestBodyChargingModeEnum struct {
	ON_DEMAND           CceIntegrationProtectionRequestBodyChargingMode
	FREE_SECURITY_CHECK CceIntegrationProtectionRequestBodyChargingMode
}

func GetCceIntegrationProtectionRequestBodyChargingModeEnum() CceIntegrationProtectionRequestBodyChargingModeEnum {
	return CceIntegrationProtectionRequestBodyChargingModeEnum{
		ON_DEMAND: CceIntegrationProtectionRequestBodyChargingMode{
			value: "on_demand",
		},
		FREE_SECURITY_CHECK: CceIntegrationProtectionRequestBodyChargingMode{
			value: "free_security_check",
		},
	}
}

func (c CceIntegrationProtectionRequestBodyChargingMode) Value() string {
	return c.value
}

func (c CceIntegrationProtectionRequestBodyChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CceIntegrationProtectionRequestBodyChargingMode) UnmarshalJSON(b []byte) error {
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

type CceIntegrationProtectionRequestBodyCceProtectionType struct {
	value string
}

type CceIntegrationProtectionRequestBodyCceProtectionTypeEnum struct {
	CLUSTER_LEVEL CceIntegrationProtectionRequestBodyCceProtectionType
	NODE_LEVEL    CceIntegrationProtectionRequestBodyCceProtectionType
}

func GetCceIntegrationProtectionRequestBodyCceProtectionTypeEnum() CceIntegrationProtectionRequestBodyCceProtectionTypeEnum {
	return CceIntegrationProtectionRequestBodyCceProtectionTypeEnum{
		CLUSTER_LEVEL: CceIntegrationProtectionRequestBodyCceProtectionType{
			value: "cluster_level",
		},
		NODE_LEVEL: CceIntegrationProtectionRequestBodyCceProtectionType{
			value: "node_level",
		},
	}
}

func (c CceIntegrationProtectionRequestBodyCceProtectionType) Value() string {
	return c.value
}

func (c CceIntegrationProtectionRequestBodyCceProtectionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CceIntegrationProtectionRequestBodyCceProtectionType) UnmarshalJSON(b []byte) error {
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
