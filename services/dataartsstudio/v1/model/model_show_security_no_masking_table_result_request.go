package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSecurityNoMaskingTableResultRequest Request Object
type ShowSecurityNoMaskingTableResultRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// 诊断任务id，通过调用查询数据权限控制模块诊断结果接口获取。
	DiagnoseId string `json:"diagnose_id"`

	// 数据源类型,HIVE
	DatasourceType *ShowSecurityNoMaskingTableResultRequestDatasourceType `json:"datasource_type,omitempty"`

	// 数据库名称
	Database *string `json:"database,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`
}

func (o ShowSecurityNoMaskingTableResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityNoMaskingTableResultRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityNoMaskingTableResultRequest", string(data)}, " ")
}

type ShowSecurityNoMaskingTableResultRequestDatasourceType struct {
	value string
}

type ShowSecurityNoMaskingTableResultRequestDatasourceTypeEnum struct {
	HIVE ShowSecurityNoMaskingTableResultRequestDatasourceType
}

func GetShowSecurityNoMaskingTableResultRequestDatasourceTypeEnum() ShowSecurityNoMaskingTableResultRequestDatasourceTypeEnum {
	return ShowSecurityNoMaskingTableResultRequestDatasourceTypeEnum{
		HIVE: ShowSecurityNoMaskingTableResultRequestDatasourceType{
			value: "HIVE",
		},
	}
}

func (c ShowSecurityNoMaskingTableResultRequestDatasourceType) Value() string {
	return c.value
}

func (c ShowSecurityNoMaskingTableResultRequestDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecurityNoMaskingTableResultRequestDatasourceType) UnmarshalJSON(b []byte) error {
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
