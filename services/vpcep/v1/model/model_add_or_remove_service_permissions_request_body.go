package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 操作权限请求结构体
type AddOrRemoveServicePermissionsRequestBody struct {
	// permission列表。 权限格式为：iam:domain::domain_id 其中， ● “iam:domain::”为固定格式。 ● “domain_id”为可连接用户的帐号ID。 支持输入1~64个字符，包括“a~z”、 “A~Z”、“0~9”或者“*”。“*”表示 所有终端节点可连接。 例如：iam:domain:: 6e9dfd51d1124e8d8498dce894923a0dd

	Permissions []string `json:"permissions"`
	// 要执行的操作。 add/remove。

	Action AddOrRemoveServicePermissionsRequestBodyAction `json:"action"`
}

func (o AddOrRemoveServicePermissionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOrRemoveServicePermissionsRequestBody struct{}"
	}

	return strings.Join([]string{"AddOrRemoveServicePermissionsRequestBody", string(data)}, " ")
}

type AddOrRemoveServicePermissionsRequestBodyAction struct {
	value string
}

type AddOrRemoveServicePermissionsRequestBodyActionEnum struct {
	ADD    AddOrRemoveServicePermissionsRequestBodyAction
	REMOVE AddOrRemoveServicePermissionsRequestBodyAction
}

func GetAddOrRemoveServicePermissionsRequestBodyActionEnum() AddOrRemoveServicePermissionsRequestBodyActionEnum {
	return AddOrRemoveServicePermissionsRequestBodyActionEnum{
		ADD: AddOrRemoveServicePermissionsRequestBodyAction{
			value: "add",
		},
		REMOVE: AddOrRemoveServicePermissionsRequestBodyAction{
			value: "remove",
		},
	}
}

func (c AddOrRemoveServicePermissionsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddOrRemoveServicePermissionsRequestBodyAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
