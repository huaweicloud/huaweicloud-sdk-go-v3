package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DataConnectorDetail struct {

	// 数据连接名称。
	ConnectorName string `json:"connector_name"`

	// 数据连接类型。 - RDS_POSTGRES：RDS服务PostgreSQL数据库 - RDS_MYSQL：RDS服务MySQL数据库 - gaussdb-mysql：云数据库GaussDB(for MySQL)
	SourceType DataConnectorDetailSourceType `json:"source_type"`

	// 数据源信息，为json格式，不同数据连接有不同的信息，各数据源的source_info请求内容可参见示例。
	SourceInfo string `json:"source_info"`

	// 数据连接ID
	ConnectorId *string `json:"connector_id,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 最后更新时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 创建时间
	CreateBy *string `json:"create_by,omitempty"`

	// 创建用户
	CreateUser *string `json:"create_user,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 最后更新时间
	LastUpdateBy *string `json:"last_update_by,omitempty"`

	// 数据连接状态
	Status *int32 `json:"status,omitempty"`

	// 使用集群
	UsedClusters *string `json:"used_clusters,omitempty"`

	// 加密类型
	EncryptType *int32 `json:"encrypt_type,omitempty"`
}

func (o DataConnectorDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataConnectorDetail struct{}"
	}

	return strings.Join([]string{"DataConnectorDetail", string(data)}, " ")
}

type DataConnectorDetailSourceType struct {
	value string
}

type DataConnectorDetailSourceTypeEnum struct {
	RDS_MYSQL     DataConnectorDetailSourceType
	RDS_POSTGRES  DataConnectorDetailSourceType
	GAUSSDB_MYSQL DataConnectorDetailSourceType
}

func GetDataConnectorDetailSourceTypeEnum() DataConnectorDetailSourceTypeEnum {
	return DataConnectorDetailSourceTypeEnum{
		RDS_MYSQL: DataConnectorDetailSourceType{
			value: "RDS_MYSQL",
		},
		RDS_POSTGRES: DataConnectorDetailSourceType{
			value: "RDS_POSTGRES",
		},
		GAUSSDB_MYSQL: DataConnectorDetailSourceType{
			value: "gaussdb-mysql",
		},
	}
}

func (c DataConnectorDetailSourceType) Value() string {
	return c.value
}

func (c DataConnectorDetailSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataConnectorDetailSourceType) UnmarshalJSON(b []byte) error {
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
