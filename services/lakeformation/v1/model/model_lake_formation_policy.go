package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LakeFormationPolicy LakeFormationPolicy
type LakeFormationPolicy struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 主体类型，USER-用户,GROUP-组,ROLE-角色,SHARE-共享,OTHER-其它
	PrincipalType LakeFormationPolicyPrincipalType `json:"principal_type"`

	// 主体来源，IAM-云,SAML-联邦,LDAP-权限策略,LOCAL-本地,OTHER-其它
	PrincipalSource LakeFormationPolicyPrincipalSource `json:"principal_source"`

	// 主体名
	PrincipalName string `json:"principal_name"`

	Resource *ResourceInfo `json:"resource,omitempty"`

	// 要求用点分格式进行分割
	ResourceName string `json:"resource_name"`

	// 权限列表
	Permissions []LakeFormationPolicyPermissions `json:"permissions"`

	// 可以传递的权限列表
	GrantAblePermissions *[]LakeFormationPolicyGrantAblePermissions `json:"grant_able_permissions,omitempty"`

	// 创建时间
	CreatedTime int64 `json:"created_time"`

	// 条件信息
	Condition *string `json:"condition,omitempty"`

	// obligation，义务，当前包含data filter和data mask
	Obligation *string `json:"obligation,omitempty"`

	// 授权路径列表
	AuthorizationPaths *[]string `json:"authorization_paths,omitempty"`
}

func (o LakeFormationPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LakeFormationPolicy struct{}"
	}

	return strings.Join([]string{"LakeFormationPolicy", string(data)}, " ")
}

type LakeFormationPolicyPrincipalType struct {
	value string
}

type LakeFormationPolicyPrincipalTypeEnum struct {
	USER  LakeFormationPolicyPrincipalType
	GROUP LakeFormationPolicyPrincipalType
	ROLE  LakeFormationPolicyPrincipalType
	SHARE LakeFormationPolicyPrincipalType
	OTHER LakeFormationPolicyPrincipalType
}

func GetLakeFormationPolicyPrincipalTypeEnum() LakeFormationPolicyPrincipalTypeEnum {
	return LakeFormationPolicyPrincipalTypeEnum{
		USER: LakeFormationPolicyPrincipalType{
			value: "USER",
		},
		GROUP: LakeFormationPolicyPrincipalType{
			value: "GROUP",
		},
		ROLE: LakeFormationPolicyPrincipalType{
			value: "ROLE",
		},
		SHARE: LakeFormationPolicyPrincipalType{
			value: "SHARE",
		},
		OTHER: LakeFormationPolicyPrincipalType{
			value: "OTHER",
		},
	}
}

func (c LakeFormationPolicyPrincipalType) Value() string {
	return c.value
}

func (c LakeFormationPolicyPrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LakeFormationPolicyPrincipalType) UnmarshalJSON(b []byte) error {
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

type LakeFormationPolicyPrincipalSource struct {
	value string
}

type LakeFormationPolicyPrincipalSourceEnum struct {
	IAM   LakeFormationPolicyPrincipalSource
	SAML  LakeFormationPolicyPrincipalSource
	LDAP  LakeFormationPolicyPrincipalSource
	LOCAL LakeFormationPolicyPrincipalSource
	OTHER LakeFormationPolicyPrincipalSource
}

func GetLakeFormationPolicyPrincipalSourceEnum() LakeFormationPolicyPrincipalSourceEnum {
	return LakeFormationPolicyPrincipalSourceEnum{
		IAM: LakeFormationPolicyPrincipalSource{
			value: "IAM",
		},
		SAML: LakeFormationPolicyPrincipalSource{
			value: "SAML",
		},
		LDAP: LakeFormationPolicyPrincipalSource{
			value: "LDAP",
		},
		LOCAL: LakeFormationPolicyPrincipalSource{
			value: "LOCAL",
		},
		OTHER: LakeFormationPolicyPrincipalSource{
			value: "OTHER",
		},
	}
}

func (c LakeFormationPolicyPrincipalSource) Value() string {
	return c.value
}

func (c LakeFormationPolicyPrincipalSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LakeFormationPolicyPrincipalSource) UnmarshalJSON(b []byte) error {
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

type LakeFormationPolicyPermissions struct {
	value string
}

type LakeFormationPolicyPermissionsEnum struct {
	ALL             LakeFormationPolicyPermissions
	CREATE          LakeFormationPolicyPermissions
	ALTER           LakeFormationPolicyPermissions
	DROP            LakeFormationPolicyPermissions
	DESCRIBE        LakeFormationPolicyPermissions
	EXEC            LakeFormationPolicyPermissions
	CREATE_DATABASE LakeFormationPolicyPermissions
	LIST_DATABASE   LakeFormationPolicyPermissions
	CREATE_TABLE    LakeFormationPolicyPermissions
	LIST_TABLE      LakeFormationPolicyPermissions
	CREATE_FUNC     LakeFormationPolicyPermissions
	LIST_FUNC       LakeFormationPolicyPermissions
	REGISTER_MODEL  LakeFormationPolicyPermissions
	LIST_MODEL      LakeFormationPolicyPermissions
	INSERT          LakeFormationPolicyPermissions
	UPDATE          LakeFormationPolicyPermissions
	DELETE          LakeFormationPolicyPermissions
	SELECT          LakeFormationPolicyPermissions
	READ            LakeFormationPolicyPermissions
	WRITE           LakeFormationPolicyPermissions
	OPERATE         LakeFormationPolicyPermissions
}

func GetLakeFormationPolicyPermissionsEnum() LakeFormationPolicyPermissionsEnum {
	return LakeFormationPolicyPermissionsEnum{
		ALL: LakeFormationPolicyPermissions{
			value: "ALL",
		},
		CREATE: LakeFormationPolicyPermissions{
			value: "CREATE",
		},
		ALTER: LakeFormationPolicyPermissions{
			value: "ALTER",
		},
		DROP: LakeFormationPolicyPermissions{
			value: "DROP",
		},
		DESCRIBE: LakeFormationPolicyPermissions{
			value: "DESCRIBE",
		},
		EXEC: LakeFormationPolicyPermissions{
			value: "EXEC",
		},
		CREATE_DATABASE: LakeFormationPolicyPermissions{
			value: "CREATE_DATABASE",
		},
		LIST_DATABASE: LakeFormationPolicyPermissions{
			value: "LIST_DATABASE",
		},
		CREATE_TABLE: LakeFormationPolicyPermissions{
			value: "CREATE_TABLE",
		},
		LIST_TABLE: LakeFormationPolicyPermissions{
			value: "LIST_TABLE",
		},
		CREATE_FUNC: LakeFormationPolicyPermissions{
			value: "CREATE_FUNC",
		},
		LIST_FUNC: LakeFormationPolicyPermissions{
			value: "LIST_FUNC",
		},
		REGISTER_MODEL: LakeFormationPolicyPermissions{
			value: "REGISTER_MODEL",
		},
		LIST_MODEL: LakeFormationPolicyPermissions{
			value: "LIST_MODEL",
		},
		INSERT: LakeFormationPolicyPermissions{
			value: "INSERT",
		},
		UPDATE: LakeFormationPolicyPermissions{
			value: "UPDATE",
		},
		DELETE: LakeFormationPolicyPermissions{
			value: "DELETE",
		},
		SELECT: LakeFormationPolicyPermissions{
			value: "SELECT",
		},
		READ: LakeFormationPolicyPermissions{
			value: "READ",
		},
		WRITE: LakeFormationPolicyPermissions{
			value: "WRITE",
		},
		OPERATE: LakeFormationPolicyPermissions{
			value: "OPERATE",
		},
	}
}

func (c LakeFormationPolicyPermissions) Value() string {
	return c.value
}

func (c LakeFormationPolicyPermissions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LakeFormationPolicyPermissions) UnmarshalJSON(b []byte) error {
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

type LakeFormationPolicyGrantAblePermissions struct {
	value string
}

type LakeFormationPolicyGrantAblePermissionsEnum struct {
	ALL             LakeFormationPolicyGrantAblePermissions
	CREATE          LakeFormationPolicyGrantAblePermissions
	ALTER           LakeFormationPolicyGrantAblePermissions
	DROP            LakeFormationPolicyGrantAblePermissions
	DESCRIBE        LakeFormationPolicyGrantAblePermissions
	EXEC            LakeFormationPolicyGrantAblePermissions
	CREATE_DATABASE LakeFormationPolicyGrantAblePermissions
	LIST_DATABASE   LakeFormationPolicyGrantAblePermissions
	CREATE_TABLE    LakeFormationPolicyGrantAblePermissions
	LIST_TABLE      LakeFormationPolicyGrantAblePermissions
	CREATE_FUNC     LakeFormationPolicyGrantAblePermissions
	LIST_FUNC       LakeFormationPolicyGrantAblePermissions
	REGISTER_MODEL  LakeFormationPolicyGrantAblePermissions
	LIST_MODEL      LakeFormationPolicyGrantAblePermissions
	INSERT          LakeFormationPolicyGrantAblePermissions
	UPDATE          LakeFormationPolicyGrantAblePermissions
	DELETE          LakeFormationPolicyGrantAblePermissions
	SELECT          LakeFormationPolicyGrantAblePermissions
	READ            LakeFormationPolicyGrantAblePermissions
	WRITE           LakeFormationPolicyGrantAblePermissions
	OPERATE         LakeFormationPolicyGrantAblePermissions
}

func GetLakeFormationPolicyGrantAblePermissionsEnum() LakeFormationPolicyGrantAblePermissionsEnum {
	return LakeFormationPolicyGrantAblePermissionsEnum{
		ALL: LakeFormationPolicyGrantAblePermissions{
			value: "ALL",
		},
		CREATE: LakeFormationPolicyGrantAblePermissions{
			value: "CREATE",
		},
		ALTER: LakeFormationPolicyGrantAblePermissions{
			value: "ALTER",
		},
		DROP: LakeFormationPolicyGrantAblePermissions{
			value: "DROP",
		},
		DESCRIBE: LakeFormationPolicyGrantAblePermissions{
			value: "DESCRIBE",
		},
		EXEC: LakeFormationPolicyGrantAblePermissions{
			value: "EXEC",
		},
		CREATE_DATABASE: LakeFormationPolicyGrantAblePermissions{
			value: "CREATE_DATABASE",
		},
		LIST_DATABASE: LakeFormationPolicyGrantAblePermissions{
			value: "LIST_DATABASE",
		},
		CREATE_TABLE: LakeFormationPolicyGrantAblePermissions{
			value: "CREATE_TABLE",
		},
		LIST_TABLE: LakeFormationPolicyGrantAblePermissions{
			value: "LIST_TABLE",
		},
		CREATE_FUNC: LakeFormationPolicyGrantAblePermissions{
			value: "CREATE_FUNC",
		},
		LIST_FUNC: LakeFormationPolicyGrantAblePermissions{
			value: "LIST_FUNC",
		},
		REGISTER_MODEL: LakeFormationPolicyGrantAblePermissions{
			value: "REGISTER_MODEL",
		},
		LIST_MODEL: LakeFormationPolicyGrantAblePermissions{
			value: "LIST_MODEL",
		},
		INSERT: LakeFormationPolicyGrantAblePermissions{
			value: "INSERT",
		},
		UPDATE: LakeFormationPolicyGrantAblePermissions{
			value: "UPDATE",
		},
		DELETE: LakeFormationPolicyGrantAblePermissions{
			value: "DELETE",
		},
		SELECT: LakeFormationPolicyGrantAblePermissions{
			value: "SELECT",
		},
		READ: LakeFormationPolicyGrantAblePermissions{
			value: "READ",
		},
		WRITE: LakeFormationPolicyGrantAblePermissions{
			value: "WRITE",
		},
		OPERATE: LakeFormationPolicyGrantAblePermissions{
			value: "OPERATE",
		},
	}
}

func (c LakeFormationPolicyGrantAblePermissions) Value() string {
	return c.value
}

func (c LakeFormationPolicyGrantAblePermissions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LakeFormationPolicyGrantAblePermissions) UnmarshalJSON(b []byte) error {
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
