package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceListRequest Request Object
type ListInstanceListRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ListInstanceListRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// api编号。
	ApiId string `json:"api_id"`

	// api操作。
	Action ListInstanceListRequestAction `json:"action"`

	// 全部展示(包括不可执行当前操作的实例)。
	ShowAll *bool `json:"show_all,omitempty"`

	// 校验api状态。
	CheckStatus *bool `json:"check_status,omitempty"`

	// 校验api调试状态。
	CheckDebug *bool `json:"check_debug,omitempty"`

	// app编号(用于判断授权操作app可选的实例)。
	AppId *string `json:"app_id,omitempty"`

	// limit。
	Limit *int32 `json:"limit,omitempty"`

	// offset。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListInstanceListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceListRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceListRequest", string(data)}, " ")
}

type ListInstanceListRequestDlmType struct {
	value string
}

type ListInstanceListRequestDlmTypeEnum struct {
	SHARED    ListInstanceListRequestDlmType
	EXCLUSIVE ListInstanceListRequestDlmType
}

func GetListInstanceListRequestDlmTypeEnum() ListInstanceListRequestDlmTypeEnum {
	return ListInstanceListRequestDlmTypeEnum{
		SHARED: ListInstanceListRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ListInstanceListRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ListInstanceListRequestDlmType) Value() string {
	return c.value
}

func (c ListInstanceListRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceListRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ListInstanceListRequestAction struct {
	value string
}

type ListInstanceListRequestActionEnum struct {
	PUBLISH   ListInstanceListRequestAction
	UNPUBLISH ListInstanceListRequestAction
	STOP      ListInstanceListRequestAction
	RECOVER   ListInstanceListRequestAction
	WHITELIST ListInstanceListRequestAction
	AUTHORIZE ListInstanceListRequestAction
}

func GetListInstanceListRequestActionEnum() ListInstanceListRequestActionEnum {
	return ListInstanceListRequestActionEnum{
		PUBLISH: ListInstanceListRequestAction{
			value: "PUBLISH",
		},
		UNPUBLISH: ListInstanceListRequestAction{
			value: "UNPUBLISH",
		},
		STOP: ListInstanceListRequestAction{
			value: "STOP",
		},
		RECOVER: ListInstanceListRequestAction{
			value: "RECOVER",
		},
		WHITELIST: ListInstanceListRequestAction{
			value: "WHITELIST",
		},
		AUTHORIZE: ListInstanceListRequestAction{
			value: "AUTHORIZE",
		},
	}
}

func (c ListInstanceListRequestAction) Value() string {
	return c.value
}

func (c ListInstanceListRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceListRequestAction) UnmarshalJSON(b []byte) error {
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
