package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AuthorizeTicketCommonInfo struct {

	// 关联单号
	TicketId *string `json:"ticket_id,omitempty"`

	// 目标 ID，一般为应用id
	TargetId *string `json:"target_id,omitempty"`

	// scope ID，一般为region id
	ScopeId *string `json:"scope_id,omitempty"`

	// 工单名称
	Title *string `json:"title,omitempty"`

	// 授权单类型，取值：CHANGE/INCIDENT/WAR_ROOM/ALARM
	TicketType *AuthorizeTicketCommonInfoTicketType `json:"ticket_type,omitempty"`

	// 当前责任人
	Owner *[]string `json:"owner,omitempty"`

	// 级别
	Level *string `json:"level,omitempty"`

	// 状态，取值：open/close
	Status *AuthorizeTicketCommonInfoStatus `json:"status,omitempty"`

	// 起始时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o AuthorizeTicketCommonInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeTicketCommonInfo struct{}"
	}

	return strings.Join([]string{"AuthorizeTicketCommonInfo", string(data)}, " ")
}

type AuthorizeTicketCommonInfoTicketType struct {
	value string
}

type AuthorizeTicketCommonInfoTicketTypeEnum struct {
	CHANGE   AuthorizeTicketCommonInfoTicketType
	INCIDENT AuthorizeTicketCommonInfoTicketType
	WAR_ROOM AuthorizeTicketCommonInfoTicketType
	ALARM    AuthorizeTicketCommonInfoTicketType
}

func GetAuthorizeTicketCommonInfoTicketTypeEnum() AuthorizeTicketCommonInfoTicketTypeEnum {
	return AuthorizeTicketCommonInfoTicketTypeEnum{
		CHANGE: AuthorizeTicketCommonInfoTicketType{
			value: "CHANGE",
		},
		INCIDENT: AuthorizeTicketCommonInfoTicketType{
			value: "INCIDENT",
		},
		WAR_ROOM: AuthorizeTicketCommonInfoTicketType{
			value: "WAR_ROOM",
		},
		ALARM: AuthorizeTicketCommonInfoTicketType{
			value: "ALARM",
		},
	}
}

func (c AuthorizeTicketCommonInfoTicketType) Value() string {
	return c.value
}

func (c AuthorizeTicketCommonInfoTicketType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizeTicketCommonInfoTicketType) UnmarshalJSON(b []byte) error {
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

type AuthorizeTicketCommonInfoStatus struct {
	value string
}

type AuthorizeTicketCommonInfoStatusEnum struct {
	OPEN  AuthorizeTicketCommonInfoStatus
	CLOSE AuthorizeTicketCommonInfoStatus
}

func GetAuthorizeTicketCommonInfoStatusEnum() AuthorizeTicketCommonInfoStatusEnum {
	return AuthorizeTicketCommonInfoStatusEnum{
		OPEN: AuthorizeTicketCommonInfoStatus{
			value: "open",
		},
		CLOSE: AuthorizeTicketCommonInfoStatus{
			value: "close",
		},
	}
}

func (c AuthorizeTicketCommonInfoStatus) Value() string {
	return c.value
}

func (c AuthorizeTicketCommonInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizeTicketCommonInfoStatus) UnmarshalJSON(b []byte) error {
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
