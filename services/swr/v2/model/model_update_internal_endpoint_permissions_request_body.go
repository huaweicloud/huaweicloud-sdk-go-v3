package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateInternalEndpointPermissionsRequestBody struct {

	// 权限格式为： - iam:domain::domain_id。其中：\"iam:domain::\"为固定格式，\"domain_id\"为可连接用户的账号ID。domain_id类型支持输入包括\"a~z\"、\"A~Z\"、\"0~9\"或者\"*\"，最大长度可以传64。 - iam:domainName::domain_name_reg。其中：\"iam:domainName::\"为固定格式，\"domain_name_reg\"为可连接用户的账号名。domain_name_reg类型支持输入包括\"a~z\"、\"A~Z\"、\"0~9\", \"_+*.-?{},\"，最大长度可以传80。 - organizations:orgPath::org_path。其中: \"organizations:orgPath::\"为固定格式，org_path为可连接用户的组织路径。 org_path类型支持\"a~z\"、\"A~Z\"、\"0~9\"、\"/-?\"或者\"*\"，最大长度可以传1024。 - \"*\" (表示所有终端节点可连接)  示例： - iam:domain::6e9dfd51d1124e8d8498dce894923a0dd - iam:domainName::op_svc_vpcep.* - organizations:orgPath::o-3j59d1231uprgk9yuvlidra7zbzfi578/r-rldbu1vmxdw5ahdkknxnvd5rgag77m2z/ou-7tuddd8nh99rebxltawsm6qct5z7rklv/_* - \"*\"（表示所有终端节点可连接）
	Permissions []string `json:"permissions"`

	// 操作类型。取值范围: - add:增加内网白名单操作 - remove:删除内网白名单操作
	Action UpdateInternalEndpointPermissionsRequestBodyAction `json:"action"`

	// 权限类型。取值范围: - domainId：基于账户ID配置终端节点服务白名单 - orgPath：基于账户所在组织路径配置终端节点服务白名单
	PermissionType *UpdateInternalEndpointPermissionsRequestBodyPermissionType `json:"permission_type,omitempty"`
}

func (o UpdateInternalEndpointPermissionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInternalEndpointPermissionsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInternalEndpointPermissionsRequestBody", string(data)}, " ")
}

type UpdateInternalEndpointPermissionsRequestBodyAction struct {
	value string
}

type UpdateInternalEndpointPermissionsRequestBodyActionEnum struct {
	ADD    UpdateInternalEndpointPermissionsRequestBodyAction
	REMOVE UpdateInternalEndpointPermissionsRequestBodyAction
}

func GetUpdateInternalEndpointPermissionsRequestBodyActionEnum() UpdateInternalEndpointPermissionsRequestBodyActionEnum {
	return UpdateInternalEndpointPermissionsRequestBodyActionEnum{
		ADD: UpdateInternalEndpointPermissionsRequestBodyAction{
			value: "add",
		},
		REMOVE: UpdateInternalEndpointPermissionsRequestBodyAction{
			value: "remove",
		},
	}
}

func (c UpdateInternalEndpointPermissionsRequestBodyAction) Value() string {
	return c.value
}

func (c UpdateInternalEndpointPermissionsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInternalEndpointPermissionsRequestBodyAction) UnmarshalJSON(b []byte) error {
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

type UpdateInternalEndpointPermissionsRequestBodyPermissionType struct {
	value string
}

type UpdateInternalEndpointPermissionsRequestBodyPermissionTypeEnum struct {
	DOMAIN_ID UpdateInternalEndpointPermissionsRequestBodyPermissionType
	ORG_PATH  UpdateInternalEndpointPermissionsRequestBodyPermissionType
}

func GetUpdateInternalEndpointPermissionsRequestBodyPermissionTypeEnum() UpdateInternalEndpointPermissionsRequestBodyPermissionTypeEnum {
	return UpdateInternalEndpointPermissionsRequestBodyPermissionTypeEnum{
		DOMAIN_ID: UpdateInternalEndpointPermissionsRequestBodyPermissionType{
			value: "domainId",
		},
		ORG_PATH: UpdateInternalEndpointPermissionsRequestBodyPermissionType{
			value: "orgPath",
		},
	}
}

func (c UpdateInternalEndpointPermissionsRequestBodyPermissionType) Value() string {
	return c.value
}

func (c UpdateInternalEndpointPermissionsRequestBodyPermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInternalEndpointPermissionsRequestBodyPermissionType) UnmarshalJSON(b []byte) error {
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
