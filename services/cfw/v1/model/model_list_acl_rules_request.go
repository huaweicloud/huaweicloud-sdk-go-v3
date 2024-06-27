package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAclRulesRequest Request Object
type ListAclRulesRequest struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。
	ObjectId string `json:"object_id"`

	// 规则Type0：互联网规则,1：vpc规则, 2:nat规则
	Type *ListAclRulesRequestType `json:"type,omitempty"`

	// ip地址
	Ip *string `json:"ip,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 方向0：外到内1：内到外
	Direction *int32 `json:"direction,omitempty"`

	// 规则下发状态 0：禁用,1：启用
	Status *ListAclRulesRequestStatus `json:"status,omitempty"`

	// 动作0：permit,1：deny
	ActionType *ListAclRulesRequestActionType `json:"action_type,omitempty"`

	// 地址类型0 ipv4
	AddressType *ListAclRulesRequestAddressType `json:"address_type,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，默认情况下，fw_instance_Id为空时，返回账号下第一个墙的信息；fw_instance_Id非空时，返回与fw_instance_Id对应墙的信息。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 标签id
	TagsId *string `json:"tags_id,omitempty"`

	// 源地址
	Source *string `json:"source,omitempty"`

	// 目的地址
	Destination *string `json:"destination,omitempty"`

	// 服务端口
	Service *string `json:"service,omitempty"`

	// 应用
	Application *string `json:"application,omitempty"`
}

func (o ListAclRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAclRulesRequest struct{}"
	}

	return strings.Join([]string{"ListAclRulesRequest", string(data)}, " ")
}

type ListAclRulesRequestType struct {
	value int32
}

type ListAclRulesRequestTypeEnum struct {
	E_0 ListAclRulesRequestType
	E_1 ListAclRulesRequestType
	E_2 ListAclRulesRequestType
}

func GetListAclRulesRequestTypeEnum() ListAclRulesRequestTypeEnum {
	return ListAclRulesRequestTypeEnum{
		E_0: ListAclRulesRequestType{
			value: 0,
		}, E_1: ListAclRulesRequestType{
			value: 1,
		}, E_2: ListAclRulesRequestType{
			value: 2,
		},
	}
}

func (c ListAclRulesRequestType) Value() int32 {
	return c.value
}

func (c ListAclRulesRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAclRulesRequestType) UnmarshalJSON(b []byte) error {
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

type ListAclRulesRequestStatus struct {
	value int32
}

type ListAclRulesRequestStatusEnum struct {
	E_0 ListAclRulesRequestStatus
	E_1 ListAclRulesRequestStatus
}

func GetListAclRulesRequestStatusEnum() ListAclRulesRequestStatusEnum {
	return ListAclRulesRequestStatusEnum{
		E_0: ListAclRulesRequestStatus{
			value: 0,
		}, E_1: ListAclRulesRequestStatus{
			value: 1,
		},
	}
}

func (c ListAclRulesRequestStatus) Value() int32 {
	return c.value
}

func (c ListAclRulesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAclRulesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListAclRulesRequestActionType struct {
	value int32
}

type ListAclRulesRequestActionTypeEnum struct {
	E_0 ListAclRulesRequestActionType
	E_1 ListAclRulesRequestActionType
}

func GetListAclRulesRequestActionTypeEnum() ListAclRulesRequestActionTypeEnum {
	return ListAclRulesRequestActionTypeEnum{
		E_0: ListAclRulesRequestActionType{
			value: 0,
		}, E_1: ListAclRulesRequestActionType{
			value: 1,
		},
	}
}

func (c ListAclRulesRequestActionType) Value() int32 {
	return c.value
}

func (c ListAclRulesRequestActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAclRulesRequestActionType) UnmarshalJSON(b []byte) error {
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

type ListAclRulesRequestAddressType struct {
	value int32
}

type ListAclRulesRequestAddressTypeEnum struct {
	E_0 ListAclRulesRequestAddressType
	E_1 ListAclRulesRequestAddressType
	E_2 ListAclRulesRequestAddressType
}

func GetListAclRulesRequestAddressTypeEnum() ListAclRulesRequestAddressTypeEnum {
	return ListAclRulesRequestAddressTypeEnum{
		E_0: ListAclRulesRequestAddressType{
			value: 0,
		}, E_1: ListAclRulesRequestAddressType{
			value: 1,
		}, E_2: ListAclRulesRequestAddressType{
			value: 2,
		},
	}
}

func (c ListAclRulesRequestAddressType) Value() int32 {
	return c.value
}

func (c ListAclRulesRequestAddressType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAclRulesRequestAddressType) UnmarshalJSON(b []byte) error {
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
