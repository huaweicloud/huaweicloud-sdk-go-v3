package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityUnreasonablePermissionsRequest Request Object
type ListSecurityUnreasonablePermissionsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// 诊断任务id，通过调用查询数据权限控制模块诊断结果接口获取。
	DiagnoseId string `json:"diagnose_id"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 数据源类型 - HIVE数据源 - DWS数据源 - [DLI数据源](tag:nohcs)
	DatasourceType *ListSecurityUnreasonablePermissionsRequestDatasourceType `json:"datasource_type,omitempty"`
}

func (o ListSecurityUnreasonablePermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityUnreasonablePermissionsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityUnreasonablePermissionsRequest", string(data)}, " ")
}

type ListSecurityUnreasonablePermissionsRequestDatasourceType struct {
	value string
}

type ListSecurityUnreasonablePermissionsRequestDatasourceTypeEnum struct {
	HIVE ListSecurityUnreasonablePermissionsRequestDatasourceType
	DWS  ListSecurityUnreasonablePermissionsRequestDatasourceType
	DLI  ListSecurityUnreasonablePermissionsRequestDatasourceType
}

func GetListSecurityUnreasonablePermissionsRequestDatasourceTypeEnum() ListSecurityUnreasonablePermissionsRequestDatasourceTypeEnum {
	return ListSecurityUnreasonablePermissionsRequestDatasourceTypeEnum{
		HIVE: ListSecurityUnreasonablePermissionsRequestDatasourceType{
			value: "HIVE",
		},
		DWS: ListSecurityUnreasonablePermissionsRequestDatasourceType{
			value: "DWS",
		},
		DLI: ListSecurityUnreasonablePermissionsRequestDatasourceType{
			value: "DLI",
		},
	}
}

func (c ListSecurityUnreasonablePermissionsRequestDatasourceType) Value() string {
	return c.value
}

func (c ListSecurityUnreasonablePermissionsRequestDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityUnreasonablePermissionsRequestDatasourceType) UnmarshalJSON(b []byte) error {
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
