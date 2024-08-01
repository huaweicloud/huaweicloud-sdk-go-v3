package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddOrRemoveServicePermissionsRequestBody 操作权限请求结构体
type AddOrRemoveServicePermissionsRequestBody struct {

	// permission列表。 权限格式为： - iam:domain::domain_id。其中： “iam:domain::”为固定格式，“domain_id”为可连接用户的账号ID。 domain_id类型支持输入包括“a~z”、“A~Z”、“0~9”或者“*”，最大长度可以传64。  - organizations:orgPath::org_path。其中： “organizations:orgPath::”为固定格式，org_path为可连接用户的组织路径。 org_path类型支持“a~z”、“A~Z”、“0~9”、“/-*?”或者“*”，最大长度可以传1024。  “*”表示所有终端节点可连接。 示例： - iam:domain::6e9dfd51d1124e8d8498dce894923a0dd - organizations:orgPath::o-3j59d1231uprgk9yuvlidra7zbzfi578/r-rldbu1vmxdw5ahdkknxnvd5rgag77m2z/ou-7tuddd8nh99rebxltawsm6qct5z7rklv/_*
	Permissions []string `json:"permissions"`

	// 终端节点服务白名单类型。  - domainId：基于账户ID配置终端节点服务白名单。  - orgPath：基于账户所在组织路径配置终端节点服务白名单。
	PermissionType *AddOrRemoveServicePermissionsRequestBodyPermissionType `json:"permission_type,omitempty"`

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

type AddOrRemoveServicePermissionsRequestBodyPermissionType struct {
	value string
}

type AddOrRemoveServicePermissionsRequestBodyPermissionTypeEnum struct {
	DOMAIN_ID AddOrRemoveServicePermissionsRequestBodyPermissionType
	ORG_PATH  AddOrRemoveServicePermissionsRequestBodyPermissionType
}

func GetAddOrRemoveServicePermissionsRequestBodyPermissionTypeEnum() AddOrRemoveServicePermissionsRequestBodyPermissionTypeEnum {
	return AddOrRemoveServicePermissionsRequestBodyPermissionTypeEnum{
		DOMAIN_ID: AddOrRemoveServicePermissionsRequestBodyPermissionType{
			value: "domainId",
		},
		ORG_PATH: AddOrRemoveServicePermissionsRequestBodyPermissionType{
			value: "orgPath",
		},
	}
}

func (c AddOrRemoveServicePermissionsRequestBodyPermissionType) Value() string {
	return c.value
}

func (c AddOrRemoveServicePermissionsRequestBodyPermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddOrRemoveServicePermissionsRequestBodyPermissionType) UnmarshalJSON(b []byte) error {
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

func (c AddOrRemoveServicePermissionsRequestBodyAction) Value() string {
	return c.value
}

func (c AddOrRemoveServicePermissionsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddOrRemoveServicePermissionsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
