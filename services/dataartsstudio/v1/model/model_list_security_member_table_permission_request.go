package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityMemberTablePermissionRequest Request Object
type ListSecurityMemberTablePermissionRequest struct {

	// IAM用户id
	UserId string `json:"user_id"`

	// 权限清单场景类型，PERMISSION_LIST
	Feature *ListSecurityMemberTablePermissionRequestFeature `json:"feature,omitempty"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// 数据源类型,hive,dws[,dli](tag:nohcs)
	DatasourceType *ListSecurityMemberTablePermissionRequestDatasourceType `json:"datasource_type,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 表名（模糊匹配）
	FuzzyTableName *string `json:"fuzzy_table_name,omitempty"`

	// 工作空间id列表
	WorkspaceIds *[]string `json:"workspace_ids,omitempty"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o ListSecurityMemberTablePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityMemberTablePermissionRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityMemberTablePermissionRequest", string(data)}, " ")
}

type ListSecurityMemberTablePermissionRequestFeature struct {
	value string
}

type ListSecurityMemberTablePermissionRequestFeatureEnum struct {
	PERMISSION_LIST ListSecurityMemberTablePermissionRequestFeature
}

func GetListSecurityMemberTablePermissionRequestFeatureEnum() ListSecurityMemberTablePermissionRequestFeatureEnum {
	return ListSecurityMemberTablePermissionRequestFeatureEnum{
		PERMISSION_LIST: ListSecurityMemberTablePermissionRequestFeature{
			value: "PERMISSION_LIST",
		},
	}
}

func (c ListSecurityMemberTablePermissionRequestFeature) Value() string {
	return c.value
}

func (c ListSecurityMemberTablePermissionRequestFeature) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityMemberTablePermissionRequestFeature) UnmarshalJSON(b []byte) error {
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

type ListSecurityMemberTablePermissionRequestDatasourceType struct {
	value string
}

type ListSecurityMemberTablePermissionRequestDatasourceTypeEnum struct {
	HIVE ListSecurityMemberTablePermissionRequestDatasourceType
	DWS  ListSecurityMemberTablePermissionRequestDatasourceType
	DLI  ListSecurityMemberTablePermissionRequestDatasourceType
}

func GetListSecurityMemberTablePermissionRequestDatasourceTypeEnum() ListSecurityMemberTablePermissionRequestDatasourceTypeEnum {
	return ListSecurityMemberTablePermissionRequestDatasourceTypeEnum{
		HIVE: ListSecurityMemberTablePermissionRequestDatasourceType{
			value: "HIVE",
		},
		DWS: ListSecurityMemberTablePermissionRequestDatasourceType{
			value: "DWS",
		},
		DLI: ListSecurityMemberTablePermissionRequestDatasourceType{
			value: "DLI",
		},
	}
}

func (c ListSecurityMemberTablePermissionRequestDatasourceType) Value() string {
	return c.value
}

func (c ListSecurityMemberTablePermissionRequestDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityMemberTablePermissionRequestDatasourceType) UnmarshalJSON(b []byte) error {
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
