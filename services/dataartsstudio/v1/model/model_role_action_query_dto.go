package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RoleActionQueryDto struct {

	// 父权限集id
	ParentPermissionSetId string `json:"parent_permission_set_id"`

	// 集群id
	ClusterId string `json:"cluster_id"`

	// 数据源类型, HIVE
	DatasourceType RoleActionQueryDtoDatasourceType `json:"datasource_type"`

	// 目前批量授权只支持单库下的多表授权，或同一集群下个多库授权，区分这两类可通过 传参中tables是否为空来判断
	DatabaseNames []string `json:"database_names"`

	// dws权限涉及 schema，预留字段，在做DWS批量授权时应保持单schema下的批量授权，或者对单库下schema批量授权
	Schemas *[]string `json:"schemas,omitempty"`

	// 数据表列表
	TableNames *[]string `json:"table_names,omitempty"`

	// 数据字段列表
	ColumnNames *[]string `json:"column_names,omitempty"`
}

func (o RoleActionQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleActionQueryDto struct{}"
	}

	return strings.Join([]string{"RoleActionQueryDto", string(data)}, " ")
}

type RoleActionQueryDtoDatasourceType struct {
	value string
}

type RoleActionQueryDtoDatasourceTypeEnum struct {
	HIVE RoleActionQueryDtoDatasourceType
}

func GetRoleActionQueryDtoDatasourceTypeEnum() RoleActionQueryDtoDatasourceTypeEnum {
	return RoleActionQueryDtoDatasourceTypeEnum{
		HIVE: RoleActionQueryDtoDatasourceType{
			value: "HIVE",
		},
	}
}

func (c RoleActionQueryDtoDatasourceType) Value() string {
	return c.value
}

func (c RoleActionQueryDtoDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RoleActionQueryDtoDatasourceType) UnmarshalJSON(b []byte) error {
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
