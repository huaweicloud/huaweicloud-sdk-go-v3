package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAclRulesRequest Request Object
type ListAclRulesRequest struct {

	// **参数解释**： 防护对象ID，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志ID，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得 **约束限制**： type为0时，object_id为互联网边界防护对象ID，type为1时，object_id为VPC边界防护对象ID，type可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得 **取值范围**： 32位UUID **默认取值**： 不涉及
	ObjectId string `json:"object_id"`

	// **参数解释**： 规则类型，用于区分不同防护对象设置规则类型。 **约束限制**： 不涉及 **取值范围**： 0：互联网边界规则，源（source）和目的（destination）地址需要为公网IP或域名； 1：VPC间规则，源（source）和目的（destination）地址需要为私有ip； 2：NAT规则，源（source）地址需要为私网IP，目的地址为公网IP或域名。 **默认取值**： 不涉及
	Type *ListAclRulesRequestType `json:"type,omitempty"`

	// **参数解释**： IP地址信息 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Ip *string `json:"ip,omitempty"`

	// **参数解释**： 规则名称，由用户定义，用于标识规则 **约束限制**： 字符串长度为0到255 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 规则方向，用于指定规则是从云上至云下，还是云下至云上 **约束限制**： 当规则type=0（互联网规则）或者type= 2（nat规则）时方向值必填 **取值范围**： 0表示外到内（云下到云上），1表示内到外（云上到云下）， **默认取值**： 不涉及
	Direction *int32 `json:"direction,omitempty"`

	// **参数解释**： 规则启用状态，用于区分规则是否启用 **约束限制**： 仅能使用数字0和1 **取值范围**： 0表示启用，1表示禁用 **默认取值**： 不涉及
	Status *ListAclRulesRequestStatus `json:"status,omitempty"`

	// **参数解释**： 规则动作类型，用于区分规则对流量的动作 **约束限制**： 仅能使用数字0和1 **取值范围**： 0表示允许通行（permit），1表示拒绝通行（deny） **默认取值**： 不涉及
	ActionType *ListAclRulesRequestActionType `json:"action_type,omitempty"`

	// **参数解释**： IP地址的互联网协议类型，用于指定IP地址的互联网协议，由客户指定 **约束限制**： 不涉及 **取值范围**： 0表示IPv4，1表示IPv6 **默认取值**： 不涉及
	AddressType *ListAclRulesRequestAddressType `json:"address_type,omitempty"`

	// **参数解释**： 每页显示个数 **约束限制**： 必须为数字 **取值范围**： 1-1024 **默认取值**： 不涉及
	Limit int32 `json:"limit"`

	// **参数解释**： 偏移量：指定返回记录的开始位置 **约束限制**： 必须为数字 **取值范围**： 大于或等于0 **默认取值**： 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**： 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，配置后可根据企业项目过滤不同企业项目下的资产，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**：  0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// **参数解释**： 规则标签id，创建规则时产生。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	TagsId *string `json:"tags_id,omitempty"`

	// **参数解释**： 源地址。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Source *string `json:"source,omitempty"`

	// **参数解释**： 目的地址。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Destination *string `json:"destination,omitempty"`

	// **参数解释**： 服务端口 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Service *string `json:"service,omitempty"`

	// **参数解释**： 规则应用协议列表 **约束限制**： 不涉及 **取值范围**： 规则应用类型包括：“HTTP”，\"HTTPS\"，\"TLS1\"，“DNS”，“SSH”，“MYSQL”，“SMTP”，“RDP”，“RDPS”，“VNC”，“POP3”，“IMAP4”，“SMTPS”，“POP3S”，“FTPS”，“ANY”，“BGP”等。 **默认取值**： 不涉及
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
}

func GetListAclRulesRequestAddressTypeEnum() ListAclRulesRequestAddressTypeEnum {
	return ListAclRulesRequestAddressTypeEnum{
		E_0: ListAclRulesRequestAddressType{
			value: 0,
		}, E_1: ListAclRulesRequestAddressType{
			value: 1,
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
