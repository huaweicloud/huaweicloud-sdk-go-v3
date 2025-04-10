package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityMemberPermissionRequest Request Object
type ListSecurityMemberPermissionRequest struct {

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// 数据源类型,HIVE
	DatasourceType *ListSecurityMemberPermissionRequestDatasourceType `json:"datasource_type,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 权限账号类型 * SELF_ACCOUNT 个人账号权限 * WORKSPACE_ACCOUNT 空间调度账号权限
	AccountType *ListSecurityMemberPermissionRequestAccountType `json:"account_type,omitempty"`

	// 权限状态,REVOKE_FAILED,TO_BE_REVOKE,INACTIVE,PERMANENTLY_ACTIVE,ACTIVE,EXPIRE_SOON
	ExpireStatus *ListSecurityMemberPermissionRequestExpireStatus `json:"expire_status,omitempty"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 排序参数,EXPIRE_TIME
	OrderBy *ListSecurityMemberPermissionRequestOrderBy `json:"order_by,omitempty"`

	// 升序/降序。true升序，fasle降序
	OrderByAsc *bool `json:"order_by_asc,omitempty"`
}

func (o ListSecurityMemberPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityMemberPermissionRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityMemberPermissionRequest", string(data)}, " ")
}

type ListSecurityMemberPermissionRequestDatasourceType struct {
	value string
}

type ListSecurityMemberPermissionRequestDatasourceTypeEnum struct {
	HIVE ListSecurityMemberPermissionRequestDatasourceType
}

func GetListSecurityMemberPermissionRequestDatasourceTypeEnum() ListSecurityMemberPermissionRequestDatasourceTypeEnum {
	return ListSecurityMemberPermissionRequestDatasourceTypeEnum{
		HIVE: ListSecurityMemberPermissionRequestDatasourceType{
			value: "HIVE",
		},
	}
}

func (c ListSecurityMemberPermissionRequestDatasourceType) Value() string {
	return c.value
}

func (c ListSecurityMemberPermissionRequestDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityMemberPermissionRequestDatasourceType) UnmarshalJSON(b []byte) error {
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

type ListSecurityMemberPermissionRequestAccountType struct {
	value string
}

type ListSecurityMemberPermissionRequestAccountTypeEnum struct {
	SELF_ACCOUNT      ListSecurityMemberPermissionRequestAccountType
	WORKSPACE_ACCOUNT ListSecurityMemberPermissionRequestAccountType
}

func GetListSecurityMemberPermissionRequestAccountTypeEnum() ListSecurityMemberPermissionRequestAccountTypeEnum {
	return ListSecurityMemberPermissionRequestAccountTypeEnum{
		SELF_ACCOUNT: ListSecurityMemberPermissionRequestAccountType{
			value: "SELF_ACCOUNT",
		},
		WORKSPACE_ACCOUNT: ListSecurityMemberPermissionRequestAccountType{
			value: "WORKSPACE_ACCOUNT",
		},
	}
}

func (c ListSecurityMemberPermissionRequestAccountType) Value() string {
	return c.value
}

func (c ListSecurityMemberPermissionRequestAccountType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityMemberPermissionRequestAccountType) UnmarshalJSON(b []byte) error {
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

type ListSecurityMemberPermissionRequestExpireStatus struct {
	value string
}

type ListSecurityMemberPermissionRequestExpireStatusEnum struct {
	REVOKE_FAILED      ListSecurityMemberPermissionRequestExpireStatus
	TO_BE_REVOKE       ListSecurityMemberPermissionRequestExpireStatus
	INACTIVE           ListSecurityMemberPermissionRequestExpireStatus
	PERMANENTLY_ACTIVE ListSecurityMemberPermissionRequestExpireStatus
	ACTIVE             ListSecurityMemberPermissionRequestExpireStatus
	EXPIRE_SOON        ListSecurityMemberPermissionRequestExpireStatus
}

func GetListSecurityMemberPermissionRequestExpireStatusEnum() ListSecurityMemberPermissionRequestExpireStatusEnum {
	return ListSecurityMemberPermissionRequestExpireStatusEnum{
		REVOKE_FAILED: ListSecurityMemberPermissionRequestExpireStatus{
			value: "REVOKE_FAILED",
		},
		TO_BE_REVOKE: ListSecurityMemberPermissionRequestExpireStatus{
			value: "TO_BE_REVOKE",
		},
		INACTIVE: ListSecurityMemberPermissionRequestExpireStatus{
			value: "INACTIVE",
		},
		PERMANENTLY_ACTIVE: ListSecurityMemberPermissionRequestExpireStatus{
			value: "PERMANENTLY_ACTIVE",
		},
		ACTIVE: ListSecurityMemberPermissionRequestExpireStatus{
			value: "ACTIVE",
		},
		EXPIRE_SOON: ListSecurityMemberPermissionRequestExpireStatus{
			value: "EXPIRE_SOON",
		},
	}
}

func (c ListSecurityMemberPermissionRequestExpireStatus) Value() string {
	return c.value
}

func (c ListSecurityMemberPermissionRequestExpireStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityMemberPermissionRequestExpireStatus) UnmarshalJSON(b []byte) error {
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

type ListSecurityMemberPermissionRequestOrderBy struct {
	value string
}

type ListSecurityMemberPermissionRequestOrderByEnum struct {
	EXPIRE_TIME ListSecurityMemberPermissionRequestOrderBy
}

func GetListSecurityMemberPermissionRequestOrderByEnum() ListSecurityMemberPermissionRequestOrderByEnum {
	return ListSecurityMemberPermissionRequestOrderByEnum{
		EXPIRE_TIME: ListSecurityMemberPermissionRequestOrderBy{
			value: "EXPIRE_TIME",
		},
	}
}

func (c ListSecurityMemberPermissionRequestOrderBy) Value() string {
	return c.value
}

func (c ListSecurityMemberPermissionRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityMemberPermissionRequestOrderBy) UnmarshalJSON(b []byte) error {
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
