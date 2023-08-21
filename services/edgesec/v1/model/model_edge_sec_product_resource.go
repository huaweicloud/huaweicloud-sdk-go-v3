package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EdgeSecProductResource struct {

	// 购买该资源的订单ID
	OrderId *string `json:"order_id,omitempty"`

	// 云服务类型，边缘安全为hws.service.type.edgesec
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 产品ID
	ProductId *string `json:"product_id,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// region ID
	RegionId *string `json:"region_id,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源规格编码
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 扩展包资源数量
	ResourceSize *int32 `json:"resource_size,omitempty"`

	// 计费方式（0:不按照流量计费， 1:带宽峰值， 2:流量）
	BillType *EdgeSecProductResourceBillType `json:"bill_type,omitempty"`

	// 收费模式（1:一次性、包周期（包年包月）， 2:按需计费）
	ChargingMode *EdgeSecProductResourceChargingMode `json:"charging_mode,omitempty"`
}

func (o EdgeSecProductResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeSecProductResource struct{}"
	}

	return strings.Join([]string{"EdgeSecProductResource", string(data)}, " ")
}

type EdgeSecProductResourceBillType struct {
	value int32
}

type EdgeSecProductResourceBillTypeEnum struct {
	E_0 EdgeSecProductResourceBillType
	E_1 EdgeSecProductResourceBillType
	E_2 EdgeSecProductResourceBillType
}

func GetEdgeSecProductResourceBillTypeEnum() EdgeSecProductResourceBillTypeEnum {
	return EdgeSecProductResourceBillTypeEnum{
		E_0: EdgeSecProductResourceBillType{
			value: 0,
		}, E_1: EdgeSecProductResourceBillType{
			value: 1,
		}, E_2: EdgeSecProductResourceBillType{
			value: 2,
		},
	}
}

func (c EdgeSecProductResourceBillType) Value() int32 {
	return c.value
}

func (c EdgeSecProductResourceBillType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EdgeSecProductResourceBillType) UnmarshalJSON(b []byte) error {
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

type EdgeSecProductResourceChargingMode struct {
	value string
}

type EdgeSecProductResourceChargingModeEnum struct {
	E_1 EdgeSecProductResourceChargingMode
	E_2 EdgeSecProductResourceChargingMode
}

func GetEdgeSecProductResourceChargingModeEnum() EdgeSecProductResourceChargingModeEnum {
	return EdgeSecProductResourceChargingModeEnum{
		E_1: EdgeSecProductResourceChargingMode{
			value: "1",
		},
		E_2: EdgeSecProductResourceChargingMode{
			value: "2",
		},
	}
}

func (c EdgeSecProductResourceChargingMode) Value() string {
	return c.value
}

func (c EdgeSecProductResourceChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EdgeSecProductResourceChargingMode) UnmarshalJSON(b []byte) error {
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
