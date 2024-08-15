package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AuthorizeTicketInfo struct {

	// 关联单号
	TicketId *string `json:"ticket_id,omitempty"`

	// 目标 ID，一般为应用id
	TargetId *string `json:"target_id,omitempty"`

	// scope ID，一般为region id
	ScopeId *string `json:"scope_id,omitempty"`

	// 工单名称
	Title *string `json:"title,omitempty"`

	// 授权单类型，取值：CHANGE/INCIDENT/WAR_ROOM/ALARM
	TicketType *AuthorizeTicketInfoTicketType `json:"ticket_type,omitempty"`

	// 当前责任人
	Owner *[]string `json:"owner,omitempty"`

	// 级别
	Level *string `json:"level,omitempty"`

	// 状态，取值：open/close
	Status *AuthorizeTicketInfoStatus `json:"status,omitempty"`

	// 起始时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o AuthorizeTicketInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeTicketInfo struct{}"
	}

	return strings.Join([]string{"AuthorizeTicketInfo", string(data)}, " ")
}

type AuthorizeTicketInfoTicketType struct {
	value string
}

type AuthorizeTicketInfoTicketTypeEnum struct {
	CHANGE   AuthorizeTicketInfoTicketType
	INCIDENT AuthorizeTicketInfoTicketType
	WAR_ROOM AuthorizeTicketInfoTicketType
	ALARM    AuthorizeTicketInfoTicketType
}

func GetAuthorizeTicketInfoTicketTypeEnum() AuthorizeTicketInfoTicketTypeEnum {
	return AuthorizeTicketInfoTicketTypeEnum{
		CHANGE: AuthorizeTicketInfoTicketType{
			value: "CHANGE",
		},
		INCIDENT: AuthorizeTicketInfoTicketType{
			value: "INCIDENT",
		},
		WAR_ROOM: AuthorizeTicketInfoTicketType{
			value: "WAR_ROOM",
		},
		ALARM: AuthorizeTicketInfoTicketType{
			value: "ALARM",
		},
	}
}

func (c AuthorizeTicketInfoTicketType) Value() string {
	return c.value
}

func (c AuthorizeTicketInfoTicketType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizeTicketInfoTicketType) UnmarshalJSON(b []byte) error {
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

type AuthorizeTicketInfoStatus struct {
	value string
}

type AuthorizeTicketInfoStatusEnum struct {
	OPEN  AuthorizeTicketInfoStatus
	CLOSE AuthorizeTicketInfoStatus
}

func GetAuthorizeTicketInfoStatusEnum() AuthorizeTicketInfoStatusEnum {
	return AuthorizeTicketInfoStatusEnum{
		OPEN: AuthorizeTicketInfoStatus{
			value: "open",
		},
		CLOSE: AuthorizeTicketInfoStatus{
			value: "close",
		},
	}
}

func (c AuthorizeTicketInfoStatus) Value() string {
	return c.value
}

func (c AuthorizeTicketInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizeTicketInfoStatus) UnmarshalJSON(b []byte) error {
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
