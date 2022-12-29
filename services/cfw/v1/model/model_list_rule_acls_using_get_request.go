package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListRuleAclsUsingGetRequest struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用查询防火墙实例接口获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。具体可参考APIExlorer和帮助中心FAQ。
	ObjectId string `json:"object_id"`

	// 规则Type0：互联网规则,1：vpc规则, 2:nat规则
	Type *ListRuleAclsUsingGetRequestType `json:"type,omitempty"`

	// 协议类型:TCP为6, UDP为17,ICMP为1,ICMPV6为58,ANY为-1
	Protocol *ListRuleAclsUsingGetRequestProtocol `json:"protocol,omitempty"`

	// ip地址
	Ip *string `json:"ip,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 方向0：外到内1：内到外
	Direction *int32 `json:"direction,omitempty"`

	// 规则下发状态 0：禁用,1：启用
	Status *ListRuleAclsUsingGetRequestStatus `json:"status,omitempty"`

	// 动作0：permit,1：deny
	ActionType *ListRuleAclsUsingGetRequestActionType `json:"action_type,omitempty"`

	// 地址类型0 ipv4,1 ipv6,2 domain
	AddressType *ListRuleAclsUsingGetRequestAddressType `json:"address_type,omitempty"`

	// 每页显示个数
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用查询防火墙实例接口获得。具体可参考APIExlorer和帮助中心FAQ。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`
}

func (o ListRuleAclsUsingGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleAclsUsingGetRequest struct{}"
	}

	return strings.Join([]string{"ListRuleAclsUsingGetRequest", string(data)}, " ")
}

type ListRuleAclsUsingGetRequestType struct {
	value int32
}

type ListRuleAclsUsingGetRequestTypeEnum struct {
	E_0 ListRuleAclsUsingGetRequestType
	E_1 ListRuleAclsUsingGetRequestType
	E_2 ListRuleAclsUsingGetRequestType
}

func GetListRuleAclsUsingGetRequestTypeEnum() ListRuleAclsUsingGetRequestTypeEnum {
	return ListRuleAclsUsingGetRequestTypeEnum{
		E_0: ListRuleAclsUsingGetRequestType{
			value: 0,
		}, E_1: ListRuleAclsUsingGetRequestType{
			value: 1,
		}, E_2: ListRuleAclsUsingGetRequestType{
			value: 2,
		},
	}
}

func (c ListRuleAclsUsingGetRequestType) Value() int32 {
	return c.value
}

func (c ListRuleAclsUsingGetRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRuleAclsUsingGetRequestType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ListRuleAclsUsingGetRequestProtocol struct {
	value int32
}

type ListRuleAclsUsingGetRequestProtocolEnum struct {
	E_6  ListRuleAclsUsingGetRequestProtocol
	E_17 ListRuleAclsUsingGetRequestProtocol
	E_1  ListRuleAclsUsingGetRequestProtocol
	E_58 ListRuleAclsUsingGetRequestProtocol
}

func GetListRuleAclsUsingGetRequestProtocolEnum() ListRuleAclsUsingGetRequestProtocolEnum {
	return ListRuleAclsUsingGetRequestProtocolEnum{
		E_6: ListRuleAclsUsingGetRequestProtocol{
			value: 6,
		}, E_17: ListRuleAclsUsingGetRequestProtocol{
			value: 17,
		}, E_1: ListRuleAclsUsingGetRequestProtocol{
			value: 1,
		}, E_58: ListRuleAclsUsingGetRequestProtocol{
			value: 58,
		},
	}
}

func (c ListRuleAclsUsingGetRequestProtocol) Value() int32 {
	return c.value
}

func (c ListRuleAclsUsingGetRequestProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRuleAclsUsingGetRequestProtocol) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ListRuleAclsUsingGetRequestStatus struct {
	value int32
}

type ListRuleAclsUsingGetRequestStatusEnum struct {
	E_0 ListRuleAclsUsingGetRequestStatus
	E_1 ListRuleAclsUsingGetRequestStatus
}

func GetListRuleAclsUsingGetRequestStatusEnum() ListRuleAclsUsingGetRequestStatusEnum {
	return ListRuleAclsUsingGetRequestStatusEnum{
		E_0: ListRuleAclsUsingGetRequestStatus{
			value: 0,
		}, E_1: ListRuleAclsUsingGetRequestStatus{
			value: 1,
		},
	}
}

func (c ListRuleAclsUsingGetRequestStatus) Value() int32 {
	return c.value
}

func (c ListRuleAclsUsingGetRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRuleAclsUsingGetRequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ListRuleAclsUsingGetRequestActionType struct {
	value int32
}

type ListRuleAclsUsingGetRequestActionTypeEnum struct {
	E_0 ListRuleAclsUsingGetRequestActionType
	E_1 ListRuleAclsUsingGetRequestActionType
}

func GetListRuleAclsUsingGetRequestActionTypeEnum() ListRuleAclsUsingGetRequestActionTypeEnum {
	return ListRuleAclsUsingGetRequestActionTypeEnum{
		E_0: ListRuleAclsUsingGetRequestActionType{
			value: 0,
		}, E_1: ListRuleAclsUsingGetRequestActionType{
			value: 1,
		},
	}
}

func (c ListRuleAclsUsingGetRequestActionType) Value() int32 {
	return c.value
}

func (c ListRuleAclsUsingGetRequestActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRuleAclsUsingGetRequestActionType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ListRuleAclsUsingGetRequestAddressType struct {
	value int32
}

type ListRuleAclsUsingGetRequestAddressTypeEnum struct {
	E_0 ListRuleAclsUsingGetRequestAddressType
	E_1 ListRuleAclsUsingGetRequestAddressType
	E_2 ListRuleAclsUsingGetRequestAddressType
}

func GetListRuleAclsUsingGetRequestAddressTypeEnum() ListRuleAclsUsingGetRequestAddressTypeEnum {
	return ListRuleAclsUsingGetRequestAddressTypeEnum{
		E_0: ListRuleAclsUsingGetRequestAddressType{
			value: 0,
		}, E_1: ListRuleAclsUsingGetRequestAddressType{
			value: 1,
		}, E_2: ListRuleAclsUsingGetRequestAddressType{
			value: 2,
		},
	}
}

func (c ListRuleAclsUsingGetRequestAddressType) Value() int32 {
	return c.value
}

func (c ListRuleAclsUsingGetRequestAddressType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRuleAclsUsingGetRequestAddressType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
