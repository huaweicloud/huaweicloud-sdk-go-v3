package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScriptInfo 脚本信息
type ScriptInfo struct {

	// 脚本名称，只能包含五种字符：英文字母、数字、中文、中划线和点号，且长度小于等于128个字符。脚本名称不能重复。
	Name *string `json:"name,omitempty"`

	// 脚本类型 - FlinkSQL: Flink SQL - DLISQL：DLI SQL - SparkSQL: MRS Spark SQL - HiveSQL: MRS Hive SQL - DWSSQL: DWS SQL - RDSSQL: RDS SQL - Shell: Shell 脚本 - PYTHON：Python 脚本 - PRESTO: MRS Presto SQL - ClickHouseSQL: MRS ClickHouse SQL - HetuEngineSQL: MRS HetuEngine SQL - ImpalaSQL: MRS Impala SQL - SparkPython: MRS Spark Python脚本
	Type *ScriptInfoType `json:"type,omitempty"`

	// 脚本所在目录路径 通过DataArts Studio管理控制台 > 数据开发，左侧列表选择“数据开发 > 脚本开发”。在脚本的目录树上，可以查看到当前已经创建的目录，默认在根目录/。
	Directory *string `json:"directory,omitempty"`

	// 脚本id。
	Id *string `json:"id,omitempty"`

	// 脚本创建人
	CreateUser *string `json:"create_user,omitempty"`

	// 脚本关联的连接名称。当type参数值为DLISQL、SparkSQL、HiveSQL、DWSSQL、Shell、PRESTO、ClickHouseSQL、ImpalaSQL、HetuEngineSQL、RDSSQL其中之一时，这个参数是必选的。用户可以通过查询连接列表（待下线）接口获取当前系统中已经存在的连接。默认值为空。
	ConnectionName *string `json:"connection_name,omitempty"`

	// 执行SQL语句所关联的数据库，当type参数值为DLISQL、SparkSQL、HiveSQL、DWSSQL、PRESTO、ClickHouseSQL、ImpalaSQL、HetuEngineSQL、RDSSQL其中之一时，才支持此参数。 - type为DLISQL时，可以通过查看所有数据库接口获取数据库信息。 - type为其他类型的时候，需要通过JDBC方式连上集群，查询数据库信息。默认值为空。
	Database *string `json:"database,omitempty"`

	// DLI资源队列名称，当type参数值为DLISQL时，才支持此参数。可以通过查询队列列表接口获取队列信息。默认值为空。
	QueueName *string `json:"queue_name,omitempty"`

	// 用户定义适用于此作业的配置参数，当type参数值为DLISQL时存在。当前支持的配置项列表请参考DLI的conf参数说明。默认值为空。
	Configuration *interface{} `json:"configuration,omitempty"`

	// 描述，长度不能超过255个字符。
	Description *string `json:"description,omitempty"`

	// 责任人名称。
	Owner *string `json:"owner,omitempty"`

	// 脚本最新提交版本。
	Version *int32 `json:"version,omitempty"`
}

func (o ScriptInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptInfo struct{}"
	}

	return strings.Join([]string{"ScriptInfo", string(data)}, " ")
}

type ScriptInfoType struct {
	value string
}

type ScriptInfoTypeEnum struct {
	FLINK_SQL       ScriptInfoType
	DLISQL          ScriptInfoType
	SPARK_SQL       ScriptInfoType
	HIVE_SQL        ScriptInfoType
	DWSSQL          ScriptInfoType
	RDSSQL          ScriptInfoType
	SHELL           ScriptInfoType
	PYTHON          ScriptInfoType
	PRESTO          ScriptInfoType
	CLICK_HOUSE_SQL ScriptInfoType
	HETU_ENGINE_SQL ScriptInfoType
	IMPALA_SQL      ScriptInfoType
	SPARK_PYTHON    ScriptInfoType
}

func GetScriptInfoTypeEnum() ScriptInfoTypeEnum {
	return ScriptInfoTypeEnum{
		FLINK_SQL: ScriptInfoType{
			value: "FlinkSQL",
		},
		DLISQL: ScriptInfoType{
			value: "DLISQL",
		},
		SPARK_SQL: ScriptInfoType{
			value: "SparkSQL",
		},
		HIVE_SQL: ScriptInfoType{
			value: "HiveSQL",
		},
		DWSSQL: ScriptInfoType{
			value: "DWSSQL",
		},
		RDSSQL: ScriptInfoType{
			value: "RDSSQL",
		},
		SHELL: ScriptInfoType{
			value: "Shell",
		},
		PYTHON: ScriptInfoType{
			value: "PYTHON",
		},
		PRESTO: ScriptInfoType{
			value: "PRESTO",
		},
		CLICK_HOUSE_SQL: ScriptInfoType{
			value: "ClickHouseSQL",
		},
		HETU_ENGINE_SQL: ScriptInfoType{
			value: "HetuEngineSQL",
		},
		IMPALA_SQL: ScriptInfoType{
			value: "ImpalaSQL",
		},
		SPARK_PYTHON: ScriptInfoType{
			value: "SparkPython",
		},
	}
}

func (c ScriptInfoType) Value() string {
	return c.value
}

func (c ScriptInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScriptInfoType) UnmarshalJSON(b []byte) error {
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
