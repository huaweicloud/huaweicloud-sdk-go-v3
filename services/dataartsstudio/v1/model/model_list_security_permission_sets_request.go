package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityPermissionSetsRequest Request Object
type ListSecurityPermissionSetsRequest struct {

	// workspace 信息
	Workspace string `json:"workspace"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 父权限集id
	ParentId *string `json:"parent_id,omitempty"`

	// 权限集类型过滤,TOP_PERMISSION_SET,SUB_PERMISSION_SET,ALL_PERMISSION_SET
	TypeFilter *ListSecurityPermissionSetsRequestTypeFilter `json:"type_filter,omitempty"`

	// 管理员id
	ManagerId *string `json:"manager_id,omitempty"`

	// 管理员名称
	ManagerName *string `json:"manager_name,omitempty"`

	// 管理员类型,USER,USER_GROUP
	ManagerType *ListSecurityPermissionSetsRequestManagerType `json:"manager_type,omitempty"`

	// 数据源类型,HIVE
	DatasourceType *ListSecurityPermissionSetsRequestDatasourceType `json:"datasource_type,omitempty"`

	// 同步状态,UNKNOWN,NOT_SYNC,SYNCING,SYNC_SUCCESS,SYNC_FAIL
	SyncStatus *ListSecurityPermissionSetsRequestSyncStatus `json:"sync_status,omitempty"`

	// 排序参数, NAME,CREATE_TIME,UPDATE_TIME
	OrderBy *ListSecurityPermissionSetsRequestOrderBy `json:"order_by,omitempty"`

	// 是否升序（仅指定排序参数时有效）
	OrderByAsc *bool `json:"order_by_asc,omitempty"`
}

func (o ListSecurityPermissionSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityPermissionSetsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityPermissionSetsRequest", string(data)}, " ")
}

type ListSecurityPermissionSetsRequestTypeFilter struct {
	value string
}

type ListSecurityPermissionSetsRequestTypeFilterEnum struct {
	TOP_PERMISSION_SET ListSecurityPermissionSetsRequestTypeFilter
	SUB_PERMISSION_SET ListSecurityPermissionSetsRequestTypeFilter
	ALL_PERMISSION_SET ListSecurityPermissionSetsRequestTypeFilter
}

func GetListSecurityPermissionSetsRequestTypeFilterEnum() ListSecurityPermissionSetsRequestTypeFilterEnum {
	return ListSecurityPermissionSetsRequestTypeFilterEnum{
		TOP_PERMISSION_SET: ListSecurityPermissionSetsRequestTypeFilter{
			value: "TOP_PERMISSION_SET",
		},
		SUB_PERMISSION_SET: ListSecurityPermissionSetsRequestTypeFilter{
			value: "SUB_PERMISSION_SET",
		},
		ALL_PERMISSION_SET: ListSecurityPermissionSetsRequestTypeFilter{
			value: "ALL_PERMISSION_SET",
		},
	}
}

func (c ListSecurityPermissionSetsRequestTypeFilter) Value() string {
	return c.value
}

func (c ListSecurityPermissionSetsRequestTypeFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityPermissionSetsRequestTypeFilter) UnmarshalJSON(b []byte) error {
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

type ListSecurityPermissionSetsRequestManagerType struct {
	value string
}

type ListSecurityPermissionSetsRequestManagerTypeEnum struct {
	USER       ListSecurityPermissionSetsRequestManagerType
	USER_GROUP ListSecurityPermissionSetsRequestManagerType
}

func GetListSecurityPermissionSetsRequestManagerTypeEnum() ListSecurityPermissionSetsRequestManagerTypeEnum {
	return ListSecurityPermissionSetsRequestManagerTypeEnum{
		USER: ListSecurityPermissionSetsRequestManagerType{
			value: "USER",
		},
		USER_GROUP: ListSecurityPermissionSetsRequestManagerType{
			value: "USER_GROUP",
		},
	}
}

func (c ListSecurityPermissionSetsRequestManagerType) Value() string {
	return c.value
}

func (c ListSecurityPermissionSetsRequestManagerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityPermissionSetsRequestManagerType) UnmarshalJSON(b []byte) error {
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

type ListSecurityPermissionSetsRequestDatasourceType struct {
	value string
}

type ListSecurityPermissionSetsRequestDatasourceTypeEnum struct {
	HIVE ListSecurityPermissionSetsRequestDatasourceType
}

func GetListSecurityPermissionSetsRequestDatasourceTypeEnum() ListSecurityPermissionSetsRequestDatasourceTypeEnum {
	return ListSecurityPermissionSetsRequestDatasourceTypeEnum{
		HIVE: ListSecurityPermissionSetsRequestDatasourceType{
			value: "HIVE",
		},
	}
}

func (c ListSecurityPermissionSetsRequestDatasourceType) Value() string {
	return c.value
}

func (c ListSecurityPermissionSetsRequestDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityPermissionSetsRequestDatasourceType) UnmarshalJSON(b []byte) error {
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

type ListSecurityPermissionSetsRequestSyncStatus struct {
	value string
}

type ListSecurityPermissionSetsRequestSyncStatusEnum struct {
	UNKNOWN      ListSecurityPermissionSetsRequestSyncStatus
	NOT_SYNC     ListSecurityPermissionSetsRequestSyncStatus
	SYNCING      ListSecurityPermissionSetsRequestSyncStatus
	SYNC_SUCCESS ListSecurityPermissionSetsRequestSyncStatus
	SYNC_FAIL    ListSecurityPermissionSetsRequestSyncStatus
}

func GetListSecurityPermissionSetsRequestSyncStatusEnum() ListSecurityPermissionSetsRequestSyncStatusEnum {
	return ListSecurityPermissionSetsRequestSyncStatusEnum{
		UNKNOWN: ListSecurityPermissionSetsRequestSyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: ListSecurityPermissionSetsRequestSyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: ListSecurityPermissionSetsRequestSyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: ListSecurityPermissionSetsRequestSyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: ListSecurityPermissionSetsRequestSyncStatus{
			value: "SYNC_FAIL",
		},
	}
}

func (c ListSecurityPermissionSetsRequestSyncStatus) Value() string {
	return c.value
}

func (c ListSecurityPermissionSetsRequestSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityPermissionSetsRequestSyncStatus) UnmarshalJSON(b []byte) error {
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

type ListSecurityPermissionSetsRequestOrderBy struct {
	value string
}

type ListSecurityPermissionSetsRequestOrderByEnum struct {
	NAME        ListSecurityPermissionSetsRequestOrderBy
	CREATE_TIME ListSecurityPermissionSetsRequestOrderBy
	UPDATE_TIME ListSecurityPermissionSetsRequestOrderBy
}

func GetListSecurityPermissionSetsRequestOrderByEnum() ListSecurityPermissionSetsRequestOrderByEnum {
	return ListSecurityPermissionSetsRequestOrderByEnum{
		NAME: ListSecurityPermissionSetsRequestOrderBy{
			value: "NAME",
		},
		CREATE_TIME: ListSecurityPermissionSetsRequestOrderBy{
			value: "CREATE_TIME",
		},
		UPDATE_TIME: ListSecurityPermissionSetsRequestOrderBy{
			value: "UPDATE_TIME",
		},
	}
}

func (c ListSecurityPermissionSetsRequestOrderBy) Value() string {
	return c.value
}

func (c ListSecurityPermissionSetsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityPermissionSetsRequestOrderBy) UnmarshalJSON(b []byte) error {
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
