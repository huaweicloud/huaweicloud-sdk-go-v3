package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSparkJobArtifactResponse **参数解释**：作业配置信息，包含作业类型、入口参数、依赖包等信息。
type ShowSparkJobArtifactResponse struct {

	// **参数解释**：Spark Jar作业入口类，用于指定作业的主类名称。 **取值范围**：完整的Java类路径格式，例如：com.example.SparkJob。
	MainClass *string `json:"main_class,omitempty"`

	// **参数解释**：Spark Python作业入口类参数列表，用于传递作业执行参数。
	MainArgs *[]string `json:"main_args,omitempty"`

	// **参数解释**：Spark Jar作业主类所在Jar包的OBS路径，用于指定作业的主Jar包。 **取值范围**：OBS URL格式，长度为1~1024个字符，例如：obs://bucket/jars/spark-job.jar。
	MainJar *string `json:"main_jar,omitempty"`

	// **参数解释**：依赖Jar包列表，用于指定Spark作业依赖的Jar包OBS路径。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// **参数解释**：依赖文件列表，用于指定Spark作业依赖的文件OBS路径。
	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	// **参数解释**：依赖归档包列表，用于指定Spark作业依赖的归档包OBS路径。
	DependencyArchives *[]string `json:"dependency_archives,omitempty"`

	// **参数解释**：依赖Python文件列表，用于指定Spark作业依赖的Python包OBS路径。
	DependencyPyFiles *[]string `json:"dependency_py_files,omitempty"`

	// **参数解释**：主Python文件路径，用于指定Python Spark作业的主文件OBS路径。 **取值范围**：OBS URL格式，长度为1~512个字符，例如：obs://bucket/pyspark/pySpark_udf_python.py.
	MainPythonFile *string `json:"main_python_file,omitempty"`

	// **参数解释**：Spark Script SQL类型作业的SQL脚本文件OBS路径，用于指定作业执行的SQL脚本。 **取值范围**：OBS URL格式，长度为1~1024个字符，例如：obs://bucket/sparksql/script.sql。
	SqlScriptingFile string `json:"sql_scripting_file"`

	// **参数解释**：SQL脚本占位符参数列表，用于为SQL脚本中的占位符传递参数值。
	SqlScriptingParameters *[]SparkJobSqlScriptParameter `json:"sql_scripting_parameters,omitempty"`

	// **参数解释**：SQL脚本结果是否写入OBS，用于控制查询结果的输出方式。 **取值范围**： - true：结果写入OBS。 - false：结果不写入OBS。
	SqlScriptingResultToObs *bool `json:"sql_scripting_result_to_obs,omitempty"`

	Result *SparkSqlScriptingResultResponse `json:"result,omitempty"`

	Type ShowSparkJobArtifactResponseType `json:"type"`
}

func (o ShowSparkJobArtifactResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobArtifactResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkJobArtifactResponse", string(data)}, " ")
}

type ShowSparkJobArtifactResponseType struct {
	value string
}

type ShowSparkJobArtifactResponseTypeEnum struct {
	SPARK_JAR_JOB           ShowSparkJobArtifactResponseType
	SPARK_PYTHON_JOB        ShowSparkJobArtifactResponseType
	SPARK_SQL_SCRIPTING_JOB ShowSparkJobArtifactResponseType
}

func GetShowSparkJobArtifactResponseTypeEnum() ShowSparkJobArtifactResponseTypeEnum {
	return ShowSparkJobArtifactResponseTypeEnum{
		SPARK_JAR_JOB: ShowSparkJobArtifactResponseType{
			value: "spark_jar_job",
		},
		SPARK_PYTHON_JOB: ShowSparkJobArtifactResponseType{
			value: "spark_python_job",
		},
		SPARK_SQL_SCRIPTING_JOB: ShowSparkJobArtifactResponseType{
			value: "spark_sql_scripting_job",
		},
	}
}

func (c ShowSparkJobArtifactResponseType) Value() string {
	return c.value
}

func (c ShowSparkJobArtifactResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSparkJobArtifactResponseType) UnmarshalJSON(b []byte) error {
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
