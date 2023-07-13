package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PermissionObject permission列表。
type PermissionObject struct {

	// permission的ID，唯一标识。
	Id *string `json:"id,omitempty"`

	// permission列表。 权限格式为：iam:domain::domain_id或者organizations:orgPath::org_path其中， ● “iam:domain::”和“organizations:orgPath::”为固定格式。 ● “domain_id”为可连接用户的帐号ID，org_path可连接用户的组织路径 domain_id类型支持输入包括“a~z”、“A~Z”、“0~9”或者“*”，org_path类型支持“a~z”、“A~Z”、“0~9”、“/-*?”或者“*”。 “*”表示所有终端节点可连接。 例如：iam:domain::6e9dfd51d1124e8d8498dce894923a0dd或者organizations:orgPath::o-3j59d1231uprgk9yuvlidra7zbzfi578/r-rldbu1vmxdw5ahdkknxnvd5rgag77m2z/ou-7tuddd8nh99rebxltawsm6qct5z7rklv/_*
	Permission *string `json:"permission,omitempty"`

	// 终端节点服务白名单类型。 ● domainId：基于账户ID配置终端节点服务白名单。 ● orgPath：基于账户所在组织路径配置终端节点服务白名单。
	PermissionType *PermissionObjectPermissionType `json:"permission_type,omitempty"`

	// 白名单的添加时间。 采用UTC时间格式，格式为：YYYY-MMDDTHH:MM:SSZ
	CreatedAt *string `json:"created_at,omitempty"`
}

func (o PermissionObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionObject struct{}"
	}

	return strings.Join([]string{"PermissionObject", string(data)}, " ")
}

type PermissionObjectPermissionType struct {
	value string
}

type PermissionObjectPermissionTypeEnum struct {
	DOMAIN_ID PermissionObjectPermissionType
	ORG_PATH  PermissionObjectPermissionType
}

func GetPermissionObjectPermissionTypeEnum() PermissionObjectPermissionTypeEnum {
	return PermissionObjectPermissionTypeEnum{
		DOMAIN_ID: PermissionObjectPermissionType{
			value: "domainId",
		},
		ORG_PATH: PermissionObjectPermissionType{
			value: "orgPath",
		},
	}
}

func (c PermissionObjectPermissionType) Value() string {
	return c.value
}

func (c PermissionObjectPermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionObjectPermissionType) UnmarshalJSON(b []byte) error {
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
