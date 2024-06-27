package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// EpsPermission 终端节点服务白名单响应体。
type EpsPermission struct {

	// 白名单表主键ID
	Id *string `json:"id,omitempty"`

	// 权限格式为：iam:domain::domain_id或者organizations:orgPath::org_path其中，  - “iam:domain::”和“organizations:orgPath::”为固定格式。  - “domain_id”为可连接用户的账号ID，org_path可连接用户的组织路径 domain_id类型支持输入包括“a~z”、“A~Z”、“0~9”或者“*”，org_path类型支持“a~z”、“A~Z”、“0~9”、“/-*?”或者“*”。 “*”表示所有终端节点可连接。 例如：iam:domain::6e9dfd51d1124e8d8498dce894923a0dd或者organizations:orgPath::o-3j59d1231uprgk9yuvlidra7zbzfi578/r-rldbu1vmxdw5ahdkknxnvd5rgag77m2z/ou-7tuddd8nh99rebxltawsm6qct5z7rklv/_*
	Permission *string `json:"permission,omitempty"`

	// 终端节点服务白名单类型。  - domainId：基于账户ID配置终端节点服务白名单。  - orgPath：基于账户所在组织路径配置终端节点服务白名单。
	PermissionType *EpsPermissionPermissionType `json:"permission_type,omitempty"`

	// 终端节点服务白名单描述
	Description *string `json:"description,omitempty"`

	// 白名单创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
}

func (o EpsPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpsPermission struct{}"
	}

	return strings.Join([]string{"EpsPermission", string(data)}, " ")
}

type EpsPermissionPermissionType struct {
	value string
}

type EpsPermissionPermissionTypeEnum struct {
	DOMAIN_ID EpsPermissionPermissionType
	ORG_PATH  EpsPermissionPermissionType
}

func GetEpsPermissionPermissionTypeEnum() EpsPermissionPermissionTypeEnum {
	return EpsPermissionPermissionTypeEnum{
		DOMAIN_ID: EpsPermissionPermissionType{
			value: "domainId",
		},
		ORG_PATH: EpsPermissionPermissionType{
			value: "orgPath",
		},
	}
}

func (c EpsPermissionPermissionType) Value() string {
	return c.value
}

func (c EpsPermissionPermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EpsPermissionPermissionType) UnmarshalJSON(b []byte) error {
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
