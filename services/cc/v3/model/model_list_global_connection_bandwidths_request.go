package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGlobalConnectionBandwidthsRequest Request Object
type ListGlobalConnectionBandwidthsRequest struct {

	// 每页返回的个数。 取值范围：1~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向翻页。 翻页过程中，查询条件不能修改，包括过滤条件，排序条件，limit。
	Marker *string `json:"marker,omitempty"`

	// 根据id查询，可查询多个id。
	Id *[]string `json:"id,omitempty"`

	// 根据名字查询，可查询多个名字。
	Name *[]string `json:"name,omitempty"`

	// 根据企业项目ID过滤列表。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 根据绑定实例id过滤全域互联带宽列表。
	InstanceId *[]string `json:"instance_id,omitempty"`

	// 根据绑定实例类型过滤全域互联带宽列表。实例类型： - CC: 云连接 - GEIP: 全域弹性公网IP - GCN: 中心网络 - GSN: 分支网络
	InstanceType *[]ListGlobalConnectionBandwidthsRequestInstanceType `json:"instance_type,omitempty"`

	// 根据支持绑定实例类型过滤全域互联带宽列表。实例类型： - CC: 云连接 - GEIP: 全域弹性公网IP - GCN: 中心网络 - GSN: 分支网络
	BindingService *[]ListGlobalConnectionBandwidthsRequestBindingService `json:"binding_service,omitempty"`

	// 根据带宽类型过滤全域互联带宽列表。带宽类型： - TrsArea: 跨区带宽 - Area: 大区带宽 - SubArea: 区域带宽 - Region: 城域带宽
	Type *[]ListGlobalConnectionBandwidthsRequestType `json:"type,omitempty"`

	// 根据带宽状态过滤全域互联带宽列表： - NORMAL: 正常 - FREEZED: 冻结
	AdminState *[]ListGlobalConnectionBandwidthsRequestAdminState `json:"admin_state,omitempty"`

	// 根据计费方式过滤全域互联带宽列表： - bwd: 按带宽计费 - 95: 按传统型95计费 - 95avr (日95计费)
	ChargeMode *[]ListGlobalConnectionBandwidthsRequestChargeMode `json:"charge_mode,omitempty"`
}

func (o ListGlobalConnectionBandwidthsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalConnectionBandwidthsRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalConnectionBandwidthsRequest", string(data)}, " ")
}

type ListGlobalConnectionBandwidthsRequestInstanceType struct {
	value string
}

type ListGlobalConnectionBandwidthsRequestInstanceTypeEnum struct {
	CC   ListGlobalConnectionBandwidthsRequestInstanceType
	GEIP ListGlobalConnectionBandwidthsRequestInstanceType
	GCN  ListGlobalConnectionBandwidthsRequestInstanceType
	GSN  ListGlobalConnectionBandwidthsRequestInstanceType
}

func GetListGlobalConnectionBandwidthsRequestInstanceTypeEnum() ListGlobalConnectionBandwidthsRequestInstanceTypeEnum {
	return ListGlobalConnectionBandwidthsRequestInstanceTypeEnum{
		CC: ListGlobalConnectionBandwidthsRequestInstanceType{
			value: "CC",
		},
		GEIP: ListGlobalConnectionBandwidthsRequestInstanceType{
			value: "GEIP",
		},
		GCN: ListGlobalConnectionBandwidthsRequestInstanceType{
			value: "GCN",
		},
		GSN: ListGlobalConnectionBandwidthsRequestInstanceType{
			value: "GSN",
		},
	}
}

func (c ListGlobalConnectionBandwidthsRequestInstanceType) Value() string {
	return c.value
}

func (c ListGlobalConnectionBandwidthsRequestInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalConnectionBandwidthsRequestInstanceType) UnmarshalJSON(b []byte) error {
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

type ListGlobalConnectionBandwidthsRequestBindingService struct {
	value string
}

type ListGlobalConnectionBandwidthsRequestBindingServiceEnum struct {
	CC   ListGlobalConnectionBandwidthsRequestBindingService
	GEIP ListGlobalConnectionBandwidthsRequestBindingService
	GCN  ListGlobalConnectionBandwidthsRequestBindingService
	GSN  ListGlobalConnectionBandwidthsRequestBindingService
}

func GetListGlobalConnectionBandwidthsRequestBindingServiceEnum() ListGlobalConnectionBandwidthsRequestBindingServiceEnum {
	return ListGlobalConnectionBandwidthsRequestBindingServiceEnum{
		CC: ListGlobalConnectionBandwidthsRequestBindingService{
			value: "CC",
		},
		GEIP: ListGlobalConnectionBandwidthsRequestBindingService{
			value: "GEIP",
		},
		GCN: ListGlobalConnectionBandwidthsRequestBindingService{
			value: "GCN",
		},
		GSN: ListGlobalConnectionBandwidthsRequestBindingService{
			value: "GSN",
		},
	}
}

func (c ListGlobalConnectionBandwidthsRequestBindingService) Value() string {
	return c.value
}

func (c ListGlobalConnectionBandwidthsRequestBindingService) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalConnectionBandwidthsRequestBindingService) UnmarshalJSON(b []byte) error {
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

type ListGlobalConnectionBandwidthsRequestType struct {
	value string
}

type ListGlobalConnectionBandwidthsRequestTypeEnum struct {
	TRS_AREA ListGlobalConnectionBandwidthsRequestType
	AREA     ListGlobalConnectionBandwidthsRequestType
	SUB_AREA ListGlobalConnectionBandwidthsRequestType
	REGION   ListGlobalConnectionBandwidthsRequestType
}

func GetListGlobalConnectionBandwidthsRequestTypeEnum() ListGlobalConnectionBandwidthsRequestTypeEnum {
	return ListGlobalConnectionBandwidthsRequestTypeEnum{
		TRS_AREA: ListGlobalConnectionBandwidthsRequestType{
			value: "TrsArea",
		},
		AREA: ListGlobalConnectionBandwidthsRequestType{
			value: "Area",
		},
		SUB_AREA: ListGlobalConnectionBandwidthsRequestType{
			value: "SubArea",
		},
		REGION: ListGlobalConnectionBandwidthsRequestType{
			value: "Region",
		},
	}
}

func (c ListGlobalConnectionBandwidthsRequestType) Value() string {
	return c.value
}

func (c ListGlobalConnectionBandwidthsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalConnectionBandwidthsRequestType) UnmarshalJSON(b []byte) error {
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

type ListGlobalConnectionBandwidthsRequestAdminState struct {
	value string
}

type ListGlobalConnectionBandwidthsRequestAdminStateEnum struct {
	NORMAL  ListGlobalConnectionBandwidthsRequestAdminState
	FREEZED ListGlobalConnectionBandwidthsRequestAdminState
}

func GetListGlobalConnectionBandwidthsRequestAdminStateEnum() ListGlobalConnectionBandwidthsRequestAdminStateEnum {
	return ListGlobalConnectionBandwidthsRequestAdminStateEnum{
		NORMAL: ListGlobalConnectionBandwidthsRequestAdminState{
			value: "NORMAL",
		},
		FREEZED: ListGlobalConnectionBandwidthsRequestAdminState{
			value: "FREEZED",
		},
	}
}

func (c ListGlobalConnectionBandwidthsRequestAdminState) Value() string {
	return c.value
}

func (c ListGlobalConnectionBandwidthsRequestAdminState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalConnectionBandwidthsRequestAdminState) UnmarshalJSON(b []byte) error {
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

type ListGlobalConnectionBandwidthsRequestChargeMode struct {
	value string
}

type ListGlobalConnectionBandwidthsRequestChargeModeEnum struct {
	BWD     ListGlobalConnectionBandwidthsRequestChargeMode
	E_95    ListGlobalConnectionBandwidthsRequestChargeMode
	E_95AVR ListGlobalConnectionBandwidthsRequestChargeMode
}

func GetListGlobalConnectionBandwidthsRequestChargeModeEnum() ListGlobalConnectionBandwidthsRequestChargeModeEnum {
	return ListGlobalConnectionBandwidthsRequestChargeModeEnum{
		BWD: ListGlobalConnectionBandwidthsRequestChargeMode{
			value: "bwd",
		},
		E_95: ListGlobalConnectionBandwidthsRequestChargeMode{
			value: "95",
		},
		E_95AVR: ListGlobalConnectionBandwidthsRequestChargeMode{
			value: "95avr",
		},
	}
}

func (c ListGlobalConnectionBandwidthsRequestChargeMode) Value() string {
	return c.value
}

func (c ListGlobalConnectionBandwidthsRequestChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalConnectionBandwidthsRequestChargeMode) UnmarshalJSON(b []byte) error {
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
