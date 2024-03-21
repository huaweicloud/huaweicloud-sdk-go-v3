package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfo 骨干带宽的信息
type AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfo struct {

	// 骨干带宽的ID
	Id *string `json:"id,omitempty"`

	// - 功能说明：骨干带宽的名称 - 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// 骨干带宽描述信息
	Description *string `json:"description,omitempty"`

	// 骨干带宽类型
	Type *AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoType `json:"type,omitempty"`

	// - 企业项目ID。最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。 - 创建全域弹性公网IP时，给全域弹性公网IP绑定企业项目ID。 - 不指定该参数时，默认值是 0 - 关于企业项目ID的获取及企业项目特性的详细信息，请参见[《企业管理用户指南》](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0126101490.html)。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 计费模式
	ChargeMode *AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoChargeMode `json:"charge_mode,omitempty"`

	// 域间带宽值
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 域间带宽大小
	Size *int32 `json:"size,omitempty"`

	// 骨干带宽的两端之一：A点
	LocalArea *string `json:"local_area,omitempty"`

	// 骨干带宽的两端之一：B点
	RemoteArea *string `json:"remote_area,omitempty"`

	// 全域弹性公网IP标签
	Tags *[]CreateGlobalEipRequestBodyGlobalEipTags `json:"tags,omitempty"`
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
