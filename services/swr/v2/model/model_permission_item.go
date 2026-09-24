package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionItem struct {

	// 权限uuid
	Id *string `json:"id,omitempty"`

	// 权限内容。权限格式为： - iam:domain::domain_id。其中：\"iam:domain::\"为固定格式，\"domain_id\"为可连接用户的账号ID。domain_id类型支持输入包括\"a~z\"、\"A~Z\"、\"0~9\"或者\"*\"，最大长度可以传64。 - iam:domainName::domain_name_reg。其中：\"iam:domainName::\"为固定格式，\"domain_name_reg\"为可连接用户的账号名。domain_name_reg类型支持输入包括\"a~z\"、\"A~Z\"、\"0~9\", \"_+*.-?{},\"，最大长度可以传80。 - organizations:orgPath::org_path。其中: \"organizations:orgPath::\"为固定格式，org_path为可连接用户的组织路径。 org_path类型支持\"a~z\"、\"A~Z\"、\"0~9\"、\"/-?\"或者\"*\"，最大长度可以传1024。 - \"*\" (表示所有终端节点可连接)  示例： - iam:domain::6e9dfd51d1124e8d8498dce894923a0dd - iam:domainName::op_svc_vpcep.* - organizations:orgPath::o-3j59d1231uprgk9yuvlidra7zbzfi578/r-rldbu1vmxdw5ahdkknxnvd5rgag77m2z/ou-7tuddd8nh99rebxltawsm6qct5z7rklv/_* - \"*\" (表示所有终端节点可连接)
	Permission *string `json:"permission,omitempty"`

	// 权限类型。取值范围： - domainId：基于账户ID配置终端节点服务白名单 - orgPath：基于账户所在组织路径配置终端节点服务白名单
	PermissionType *PermissionItemPermissionType `json:"permission_type,omitempty"`

	// 白名单的添加时间。采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ
	CreatedAt *string `json:"created_at,omitempty"`

	// 是否为保护内网访问白名单；如果为true则不允许添加或者移除
	Protected *bool `json:"protected,omitempty"`
}

func (o PermissionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionItem struct{}"
	}

	return strings.Join([]string{"PermissionItem", string(data)}, " ")
}

type PermissionItemPermissionType struct {
	value string
}

type PermissionItemPermissionTypeEnum struct {
	DOMAIN_ID PermissionItemPermissionType
	ORG_PATH  PermissionItemPermissionType
}

func GetPermissionItemPermissionTypeEnum() PermissionItemPermissionTypeEnum {
	return PermissionItemPermissionTypeEnum{
		DOMAIN_ID: PermissionItemPermissionType{
			value: "domainId",
		},
		ORG_PATH: PermissionItemPermissionType{
			value: "orgPath",
		},
	}
}

func (c PermissionItemPermissionType) Value() string {
	return c.value
}

func (c PermissionItemPermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionItemPermissionType) UnmarshalJSON(b []byte) error {
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
