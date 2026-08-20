package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityMemberPermissionsByUserIdRequest Request Object
type ListSecurityMemberPermissionsByUserIdRequest struct {

	// IAM用户id
	UserId string `json:"user_id"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// 数据源类型 - HIVE数据源 - DWS数据源 - [DLI数据源](tag:nohcs)
	DatasourceType *ListSecurityMemberPermissionsByUserIdRequestDatasourceType `json:"datasource_type,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// Schema名，正向模糊匹配
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 权限账号类型 * SELF_ACCOUNT 个人账号权限 * WORKSPACE_ACCOUNT 空间调度账号权限
	AccountType *ListSecurityMemberPermissionsByUserIdRequestAccountType `json:"account_type,omitempty"`

	// 权限状态,REVOKE_FAILED,TO_BE_REVOKE,INACTIVE,PERMANENTLY_ACTIVE,ACTIVE,EXPIRE_SOON
	ExpireStatus *ListSecurityMemberPermissionsByUserIdRequestExpireStatus `json:"expire_status,omitempty"`

	// 过期时间开始时间戳，毫秒。
	StartExpireTime *int64 `json:"start_expire_time,omitempty"`

	// 过期时间结束时间戳，毫秒。
	EndExpireTime *int64 `json:"end_expire_time,omitempty"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 排序参数,EXPIRE_TIME
	OrderBy *ListSecurityMemberPermissionsByUserIdRequestOrderBy `json:"order_by,omitempty"`

	// 升序/降序。true升序，false降序
	OrderByAsc *bool `json:"order_by_asc,omitempty"`
}

func (o ListSecurityMemberPermissionsByUserIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityMemberPermissionsByUserIdRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityMemberPermissionsByUserIdRequest", string(data)}, " ")
}

type ListSecurityMemberPermissionsByUserIdRequestDatasourceType struct {
	value string
}

type ListSecurityMemberPermissionsByUserIdRequestDatasourceTypeEnum struct {
	HIVE ListSecurityMemberPermissionsByUserIdRequestDatasourceType
	DWS  ListSecurityMemberPermissionsByUserIdRequestDatasourceType
	DLI  ListSecurityMemberPermissionsByUserIdRequestDatasourceType
}

func GetListSecurityMemberPermissionsByUserIdRequestDatasourceTypeEnum() ListSecurityMemberPermissionsByUserIdRequestDatasourceTypeEnum {
	return ListSecurityMemberPermissionsByUserIdRequestDatasourceTypeEnum{
		HIVE: ListSecurityMemberPermissionsByUserIdRequestDatasourceType{
			value: "HIVE",
		},
		DWS: ListSecurityMemberPermissionsByUserIdRequestDatasourceType{
			value: "DWS",
		},
		DLI: ListSecurityMemberPermissionsByUserIdRequestDatasourceType{
			value: "DLI",
		},
	}
}

func (c ListSecurityMemberPermissionsByUserIdRequestDatasourceType) Value() string {
	return c.value
}

func (c ListSecurityMemberPermissionsByUserIdRequestDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityMemberPermissionsByUserIdRequestDatasourceType) UnmarshalJSON(b []byte) error {
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

type ListSecurityMemberPermissionsByUserIdRequestAccountType struct {
	value string
}

type ListSecurityMemberPermissionsByUserIdRequestAccountTypeEnum struct {
	SELF_ACCOUNT      ListSecurityMemberPermissionsByUserIdRequestAccountType
	WORKSPACE_ACCOUNT ListSecurityMemberPermissionsByUserIdRequestAccountType
}

func GetListSecurityMemberPermissionsByUserIdRequestAccountTypeEnum() ListSecurityMemberPermissionsByUserIdRequestAccountTypeEnum {
	return ListSecurityMemberPermissionsByUserIdRequestAccountTypeEnum{
		SELF_ACCOUNT: ListSecurityMemberPermissionsByUserIdRequestAccountType{
			value: "SELF_ACCOUNT",
		},
		WORKSPACE_ACCOUNT: ListSecurityMemberPermissionsByUserIdRequestAccountType{
			value: "WORKSPACE_ACCOUNT",
		},
	}
}

func (c ListSecurityMemberPermissionsByUserIdRequestAccountType) Value() string {
	return c.value
}

func (c ListSecurityMemberPermissionsByUserIdRequestAccountType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityMemberPermissionsByUserIdRequestAccountType) UnmarshalJSON(b []byte) error {
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

type ListSecurityMemberPermissionsByUserIdRequestExpireStatus struct {
	value string
}

type ListSecurityMemberPermissionsByUserIdRequestExpireStatusEnum struct {
	REVOKE_FAILED      ListSecurityMemberPermissionsByUserIdRequestExpireStatus
	TO_BE_REVOKE       ListSecurityMemberPermissionsByUserIdRequestExpireStatus
	INACTIVE           ListSecurityMemberPermissionsByUserIdRequestExpireStatus
	PERMANENTLY_ACTIVE ListSecurityMemberPermissionsByUserIdRequestExpireStatus
	ACTIVE             ListSecurityMemberPermissionsByUserIdRequestExpireStatus
	EXPIRE_SOON        ListSecurityMemberPermissionsByUserIdRequestExpireStatus
}

func GetListSecurityMemberPermissionsByUserIdRequestExpireStatusEnum() ListSecurityMemberPermissionsByUserIdRequestExpireStatusEnum {
	return ListSecurityMemberPermissionsByUserIdRequestExpireStatusEnum{
		REVOKE_FAILED: ListSecurityMemberPermissionsByUserIdRequestExpireStatus{
			value: "REVOKE_FAILED",
		},
		TO_BE_REVOKE: ListSecurityMemberPermissionsByUserIdRequestExpireStatus{
			value: "TO_BE_REVOKE",
		},
		INACTIVE: ListSecurityMemberPermissionsByUserIdRequestExpireStatus{
			value: "INACTIVE",
		},
		PERMANENTLY_ACTIVE: ListSecurityMemberPermissionsByUserIdRequestExpireStatus{
			value: "PERMANENTLY_ACTIVE",
		},
		ACTIVE: ListSecurityMemberPermissionsByUserIdRequestExpireStatus{
			value: "ACTIVE",
		},
		EXPIRE_SOON: ListSecurityMemberPermissionsByUserIdRequestExpireStatus{
			value: "EXPIRE_SOON",
		},
	}
}

func (c ListSecurityMemberPermissionsByUserIdRequestExpireStatus) Value() string {
	return c.value
}

func (c ListSecurityMemberPermissionsByUserIdRequestExpireStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityMemberPermissionsByUserIdRequestExpireStatus) UnmarshalJSON(b []byte) error {
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

type ListSecurityMemberPermissionsByUserIdRequestOrderBy struct {
	value string
}

type ListSecurityMemberPermissionsByUserIdRequestOrderByEnum struct {
	EXPIRE_TIME ListSecurityMemberPermissionsByUserIdRequestOrderBy
}

func GetListSecurityMemberPermissionsByUserIdRequestOrderByEnum() ListSecurityMemberPermissionsByUserIdRequestOrderByEnum {
	return ListSecurityMemberPermissionsByUserIdRequestOrderByEnum{
		EXPIRE_TIME: ListSecurityMemberPermissionsByUserIdRequestOrderBy{
			value: "EXPIRE_TIME",
		},
	}
}

func (c ListSecurityMemberPermissionsByUserIdRequestOrderBy) Value() string {
	return c.value
}

func (c ListSecurityMemberPermissionsByUserIdRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityMemberPermissionsByUserIdRequestOrderBy) UnmarshalJSON(b []byte) error {
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
