package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DynamicMaskingPolicyCreateDto struct {

	// 策略名称。英文和汉字开头, 支持英文、汉字、数字、下划线, 2-64字符。
	Name string `json:"name"`

	// 数据源类型 - HIVE数据源 - DWS数据源 - [DLI数据源](tag:nohcs)
	DatasourceType DynamicMaskingPolicyCreateDtoDatasourceType `json:"datasource_type"`

	// 集群id。请于集群管理页面查看集群ID信息。[当数据源类型为DLI时，该参数需要填写为DLI](tag:nohcs)。
	ClusterId string `json:"cluster_id"`

	// 集群名称。请于集群管理页面查看集群名称信息。[当数据源类型为DLI时，该参数需要填写为DLI](tag:nohcs)。
	ClusterName string `json:"cluster_name"`

	// 数据库名称。获取方法请参见[获取数据源中的表](getDataTables.html)。
	DatabaseName string `json:"database_name"`

	// 数据表id，获取方法请参见[获取数据源中的表](getDataTables.html)。
	TableId *string `json:"table_id,omitempty"`

	// 数据表名称, 获取方法请参见[获取数据源中的表](getDataTables.html)。
	TableName string `json:"table_name"`

	// 用户组列表，用户组名称逗号分隔（非必填项，但用户、用户组必须二选其一进行配置）。例如：\"userGroup1,userGroup2\"。
	UserGroups *string `json:"user_groups,omitempty"`

	// 用户列表，用户名称逗号分隔（非必填项，但用户、用户组必须二选其一进行配置），例如：\"user1,user2\"。
	Users *string `json:"users,omitempty"`

	// 数据连接名称，获取方法请参见[查询数据连接列表](ListDataconnections.html)。
	ConnName string `json:"conn_name"`

	// 数据连接id，获取方法请参见[查询数据连接列表](ListDataconnections.html)。
	ConnId string `json:"conn_id"`

	// DWS数据源的模式名称。
	SchemaName *string `json:"schema_name,omitempty"`

	// 动态数据脱敏策略列表。
	PolicyList []DynamicMaskingPolicyCreate `json:"policy_list"`
}

func (o DynamicMaskingPolicyCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DynamicMaskingPolicyCreateDto struct{}"
	}

	return strings.Join([]string{"DynamicMaskingPolicyCreateDto", string(data)}, " ")
}

type DynamicMaskingPolicyCreateDtoDatasourceType struct {
	value string
}

type DynamicMaskingPolicyCreateDtoDatasourceTypeEnum struct {
	HIVE DynamicMaskingPolicyCreateDtoDatasourceType
	DWS  DynamicMaskingPolicyCreateDtoDatasourceType
	DLI  DynamicMaskingPolicyCreateDtoDatasourceType
}

func GetDynamicMaskingPolicyCreateDtoDatasourceTypeEnum() DynamicMaskingPolicyCreateDtoDatasourceTypeEnum {
	return DynamicMaskingPolicyCreateDtoDatasourceTypeEnum{
		HIVE: DynamicMaskingPolicyCreateDtoDatasourceType{
			value: "HIVE",
		},
		DWS: DynamicMaskingPolicyCreateDtoDatasourceType{
			value: "DWS",
		},
		DLI: DynamicMaskingPolicyCreateDtoDatasourceType{
			value: "DLI",
		},
	}
}

func (c DynamicMaskingPolicyCreateDtoDatasourceType) Value() string {
	return c.value
}

func (c DynamicMaskingPolicyCreateDtoDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DynamicMaskingPolicyCreateDtoDatasourceType) UnmarshalJSON(b []byte) error {
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
