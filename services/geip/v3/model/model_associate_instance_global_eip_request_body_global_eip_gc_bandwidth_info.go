package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfo struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Type *AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoType `json:"type,omitempty"`

	// 资源的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 计费模式
	ChargeMode *AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoChargeMode `json:"charge_mode,omitempty"`

	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 大小
	Size *int32 `json:"size,omitempty"`

	// 骨干带宽的两端之一：A点
	LocalArea *string `json:"local_area,omitempty"`

	// 骨干带宽的两端之一：B点
	RemoteArea *string `json:"remote_area,omitempty"`

	// 全域弹性公网IP标签
	Tags *[]AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoTags `json:"tags,omitempty"`
}

func (o AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfo struct{}"
	}

	return strings.Join([]string{"AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfo", string(data)}, " ")
}

type AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoType struct {
	value string
}

type AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoTypeEnum struct {
	REGION AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoType
}

func GetAssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoTypeEnum() AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoTypeEnum {
	return AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoTypeEnum{
		REGION: AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoType{
			value: "Region",
		},
	}
}

func (c AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoType) Value() string {
	return c.value
}

func (c AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoType) UnmarshalJSON(b []byte) error {
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

type AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoChargeMode struct {
	value string
}

type AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoChargeModeEnum struct {
	BWD AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoChargeMode
}

func GetAssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoChargeModeEnum() AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoChargeModeEnum {
	return AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoChargeModeEnum{
		BWD: AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoChargeMode{
			value: "bwd",
		},
	}
}

func (c AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoChargeMode) Value() string {
	return c.value
}

func (c AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoChargeMode) UnmarshalJSON(b []byte) error {
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
