package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type LdApiScriptBase struct {

	// 数据源编号，当api_type = data时，必选
	DsId *string `json:"ds_id,omitempty"`

	// 数据源名称
	DsName *string `json:"ds_name,omitempty"`

	// 数据源类型： - oracle：oracle数据源类型 - mysql：mysql数据源类型 - mongodb：mongodb数据源类型 - redis：redis数据源类型 - postgresql：postgresql数据源类型 - hive：hive数据源类型 - mssql：sqlserver数据源类型 - sqlserver：sqlserver数据源类型 - gauss200：gauss200数据源类型 - dws：dws数据源类型 - gauss100：gauss100数据源类型 - zenith：zenith数据源类型
	DsType *LdApiScriptBaseDsType `json:"ds_type,omitempty"`

	// 脚本类型 - SQL：sql语句 - SP：存储过程
	Type *string `json:"type,omitempty"`

	// 返回对象。  当api_type = data时，必选
	ObjectName *string `json:"object_name,omitempty"`

	// API脚本内容  请对脚本进行base64编码
	Content string `json:"content"`

	// 数据脚本是否结果分页，当api_type = data时有效
	EnableResultPaging *bool `json:"enable_result_paging,omitempty"`

	// 数据脚本是否预编译，当api_type = data时有效
	EnablePreparestatement *bool `json:"enable_preparestatement,omitempty"`
}

func (o LdApiScriptBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdApiScriptBase struct{}"
	}

	return strings.Join([]string{"LdApiScriptBase", string(data)}, " ")
}

type LdApiScriptBaseDsType struct {
	value string
}

type LdApiScriptBaseDsTypeEnum struct {
	ORACLE     LdApiScriptBaseDsType
	MYSQL      LdApiScriptBaseDsType
	MONGODB    LdApiScriptBaseDsType
	REDIS      LdApiScriptBaseDsType
	POSTGRESQL LdApiScriptBaseDsType
	HIVE       LdApiScriptBaseDsType
	MSSQL      LdApiScriptBaseDsType
	SQLSERVER  LdApiScriptBaseDsType
	GAUSS200   LdApiScriptBaseDsType
	DWS        LdApiScriptBaseDsType
	GAUSS100   LdApiScriptBaseDsType
	ZENITH     LdApiScriptBaseDsType
}

func GetLdApiScriptBaseDsTypeEnum() LdApiScriptBaseDsTypeEnum {
	return LdApiScriptBaseDsTypeEnum{
		ORACLE: LdApiScriptBaseDsType{
			value: "oracle",
		},
		MYSQL: LdApiScriptBaseDsType{
			value: "mysql",
		},
		MONGODB: LdApiScriptBaseDsType{
			value: "mongodb",
		},
		REDIS: LdApiScriptBaseDsType{
			value: "redis",
		},
		POSTGRESQL: LdApiScriptBaseDsType{
			value: "postgresql",
		},
		HIVE: LdApiScriptBaseDsType{
			value: "hive",
		},
		MSSQL: LdApiScriptBaseDsType{
			value: "mssql",
		},
		SQLSERVER: LdApiScriptBaseDsType{
			value: "sqlserver",
		},
		GAUSS200: LdApiScriptBaseDsType{
			value: "gauss200",
		},
		DWS: LdApiScriptBaseDsType{
			value: "dws",
		},
		GAUSS100: LdApiScriptBaseDsType{
			value: "gauss100",
		},
		ZENITH: LdApiScriptBaseDsType{
			value: "zenith",
		},
	}
}

func (c LdApiScriptBaseDsType) Value() string {
	return c.value
}

func (c LdApiScriptBaseDsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdApiScriptBaseDsType) UnmarshalJSON(b []byte) error {
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
