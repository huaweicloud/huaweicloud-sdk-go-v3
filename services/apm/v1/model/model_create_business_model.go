package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateBusinessModel struct {

	// 应用名字
	Name string `json:"name"`

	// 企业项目ID，默认值为“0”，表示默认项目的ID。
	EpsId *string `json:"eps_id,omitempty"`

	// CMDB树显示的名称
	DisplayName string `json:"display_name"`

	// 描述
	Descp string `json:"descp"`

	// 默认值为SKYWALKING。
	CmdbDatasourceType *CreateBusinessModelCmdbDatasourceType `json:"cmdb_datasource_type,omitempty"`
}

func (o CreateBusinessModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBusinessModel struct{}"
	}

	return strings.Join([]string{"CreateBusinessModel", string(data)}, " ")
}

type CreateBusinessModelCmdbDatasourceType struct {
	value string
}

type CreateBusinessModelCmdbDatasourceTypeEnum struct {
	OTEL       CreateBusinessModelCmdbDatasourceType
	SKYWALKING CreateBusinessModelCmdbDatasourceType
}

func GetCreateBusinessModelCmdbDatasourceTypeEnum() CreateBusinessModelCmdbDatasourceTypeEnum {
	return CreateBusinessModelCmdbDatasourceTypeEnum{
		OTEL: CreateBusinessModelCmdbDatasourceType{
			value: "OTEL",
		},
		SKYWALKING: CreateBusinessModelCmdbDatasourceType{
			value: "SKYWALKING",
		},
	}
}

func (c CreateBusinessModelCmdbDatasourceType) Value() string {
	return c.value
}

func (c CreateBusinessModelCmdbDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBusinessModelCmdbDatasourceType) UnmarshalJSON(b []byte) error {
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
