package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBindingGeip Returns geip and its binding status
type ListBindingGeip struct {

	// geip的id
	GlobalEipId *string `json:"global_eip_id,omitempty"`

	// 网段geip的id
	GlobalEipSegmentId *string `json:"global_eip_segment_id,omitempty"`

	// geip的绑定状态
	Status *ListBindingGeipStatus `json:"status,omitempty"`

	// geip类型：IP_ADDRESS/IP_SEGMENT
	Type *ListBindingGeipType `json:"type,omitempty"`

	// geip绑定失败的原因
	ErrorMessage *string `json:"error_message,omitempty"`

	// geip的地址ip/mask
	Cidr *string `json:"cidr,omitempty"`

	// geip的地址簇
	AddressFamily *string `json:"address_family,omitempty"`

	// CloudPond的集群vtepIp
	IeVtepIp *string `json:"ie_vtep_ip,omitempty"`

	// geip绑定时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 带宽包的id
	GcbId *string `json:"gcb_id,omitempty"`
}

func (o ListBindingGeip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBindingGeip struct{}"
	}

	return strings.Join([]string{"ListBindingGeip", string(data)}, " ")
}

type ListBindingGeipStatus struct {
	value string
}

type ListBindingGeipStatusEnum struct {
	ACTIVE ListBindingGeipStatus
	ERROR  ListBindingGeipStatus
}

func GetListBindingGeipStatusEnum() ListBindingGeipStatusEnum {
	return ListBindingGeipStatusEnum{
		ACTIVE: ListBindingGeipStatus{
			value: "ACTIVE",
		},
		ERROR: ListBindingGeipStatus{
			value: "ERROR",
		},
	}
}

func (c ListBindingGeipStatus) Value() string {
	return c.value
}

func (c ListBindingGeipStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBindingGeipStatus) UnmarshalJSON(b []byte) error {
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

type ListBindingGeipType struct {
	value string
}

type ListBindingGeipTypeEnum struct {
	IP_ADDRESS ListBindingGeipType
	IP_SEGMENT ListBindingGeipType
}

func GetListBindingGeipTypeEnum() ListBindingGeipTypeEnum {
	return ListBindingGeipTypeEnum{
		IP_ADDRESS: ListBindingGeipType{
			value: "IP_ADDRESS",
		},
		IP_SEGMENT: ListBindingGeipType{
			value: "IP_SEGMENT",
		},
	}
}

func (c ListBindingGeipType) Value() string {
	return c.value
}

func (c ListBindingGeipType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBindingGeipType) UnmarshalJSON(b []byte) error {
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
