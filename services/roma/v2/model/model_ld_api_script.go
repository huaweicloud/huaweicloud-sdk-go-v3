package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type LdApiScript struct {

	// 数据源编号，当api_type = data时，必选
	DsId *string `json:"ds_id,omitempty"`

	// 数据源名称
	DsName *string `json:"ds_name,omitempty"`

	// 数据源类型： - oracle：oracle数据源类型 - mysql：mysql数据源类型 - mongodb：mongodb数据源类型 - redis：redis数据源类型 - postgresql：postgresql数据源类型 - hive：hive数据源类型 - mssql：sqlserver数据源类型 - sqlserver：sqlserver数据源类型 - gauss200：gauss200数据源类型 - dws：dws数据源类型 - gauss100：gauss100数据源类型 - zenith：zenith数据源类型
	DsType *LdApiScriptDsType `json:"ds_type,omitempty"`

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

	// 后端API脚本创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 后端API脚本修改时间
	ModifiedTime *sdktime.SdkTime `json:"modified_time,omitempty"`
}

func (o LdApiScript) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdApiScript struct{}"
	}

	return strings.Join([]string{"LdApiScript", string(data)}, " ")
}

type LdApiScriptDsType struct {
	value string
}

type LdApiScriptDsTypeEnum struct {
	ORACLE     LdApiScriptDsType
	MYSQL      LdApiScriptDsType
	MONGODB    LdApiScriptDsType
	REDIS      LdApiScriptDsType
	POSTGRESQL LdApiScriptDsType
	HIVE       LdApiScriptDsType
	MSSQL      LdApiScriptDsType
	SQLSERVER  LdApiScriptDsType
	GAUSS200   LdApiScriptDsType
	DWS        LdApiScriptDsType
	GAUSS100   LdApiScriptDsType
	ZENITH     LdApiScriptDsType
}

func GetLdApiScriptDsTypeEnum() LdApiScriptDsTypeEnum {
	return LdApiScriptDsTypeEnum{
		ORACLE: LdApiScriptDsType{
			value: "oracle",
		},
		MYSQL: LdApiScriptDsType{
			value: "mysql",
		},
		MONGODB: LdApiScriptDsType{
			value: "mongodb",
		},
		REDIS: LdApiScriptDsType{
			value: "redis",
		},
		POSTGRESQL: LdApiScriptDsType{
			value: "postgresql",
		},
		HIVE: LdApiScriptDsType{
			value: "hive",
		},
		MSSQL: LdApiScriptDsType{
			value: "mssql",
		},
		SQLSERVER: LdApiScriptDsType{
			value: "sqlserver",
		},
		GAUSS200: LdApiScriptDsType{
			value: "gauss200",
		},
		DWS: LdApiScriptDsType{
			value: "dws",
		},
		GAUSS100: LdApiScriptDsType{
			value: "gauss100",
		},
		ZENITH: LdApiScriptDsType{
			value: "zenith",
		},
	}
}

func (c LdApiScriptDsType) Value() string {
	return c.value
}

func (c LdApiScriptDsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdApiScriptDsType) UnmarshalJSON(b []byte) error {
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
