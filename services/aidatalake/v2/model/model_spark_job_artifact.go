package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SparkJobArtifact Spark作业参数，根据作业类型自动选择对应的参数结构。
type SparkJobArtifact struct {
	Type SparkJobArtifactType `json:"type"`

	// **参数解释**：主类名称，用于指定Spark Jar作业的入口类。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符。 **默认取值**：不涉及。
	MainClass *string `json:"main_class,omitempty"`

	// **参数解释**：主类参数列表，用于传递给Python Spark作业的参数。多个参数之间空格分隔。 **约束限制**：参数数量不能超过100个。
	MainArgs *[]string `json:"main_args,omitempty"`

	// **参数解释**：主Jar包路径，用于指定Spark Jar作业主类所在Jar包的OBS路径。当作业类型为“spark_jar_job”时，此参数必填。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符，格式为OBS路径，例如：“obs://bucket_name/path/to/jar.jar”。 **默认取值**：不涉及。
	MainJar *string `json:"main_jar,omitempty"`

	// **参数解释**：依赖Jar包列表，用于指定Spark作业依赖的Jar包OBS路径。 **约束限制**：依赖Jar包数量不能超过100个。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// **参数解释**：依赖文件列表，用于指定Spark作业依赖的文件OBS路径。 **约束限制**：依赖文件数量不能超过100个。
	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	// **参数解释**：依赖归档包列表，用于指定Spark作业依赖的归档包OBS路径。 **约束限制**：依赖归档包数量不能超过100个。
	DependencyArchives *[]string `json:"dependency_archives,omitempty"`

	// **参数解释**：依赖Python文件列表，用于指定Spark作业依赖的Python包OBS路径。 **约束限制**：依赖Python文件数量不能超过100个。
	DependencyPyFiles *[]string `json:"dependency_py_files,omitempty"`

	// **参数解释**：主Python文件路径，用于指定Python Spark作业的主文件OBS路径。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符，格式为OBS路径，例如：“obs://bucket/pyspark/pySpark_udf_python.py”。 **默认取值**：不涉及。
	MainPythonFile string `json:"main_python_file"`

	// **参数解释**：SQL脚本文件路径，用于指定Spark Script SQL作业的脚本文件OBS路径。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符，格式为OBS路径，例如：“obs://bucket/sparksqlscript/script.sql”。 **默认取值**：不涉及。
	SqlScriptingFile string `json:"sql_scripting_file"`

	// **参数解释**：SQL脚本占位符参数列表，用于配置SQL脚本中的参数化变量。数组中的每个元素为SparkJobSqlScriptParameter对象，包含占位符的键、值和类型信息。 **约束限制**：占位符参数数量不能超过16条。
	SqlScriptingParameters *[]SparkJobSqlScriptParameter `json:"sql_scripting_parameters,omitempty"`

	// **参数解释**：结果写入OBS标识，用于控制SQL脚本作业的执行结果是否写入OBS。 **约束限制**：不涉及。 **取值范围**： - true：结果写入OBS。 - false：结果不写入OBS。 **默认取值**：不涉及。
	SqlScriptingResultToObs *bool `json:"sql_scripting_result_to_obs,omitempty"`
}

func (o SparkJobArtifact) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJobArtifact struct{}"
	}

	return strings.Join([]string{"SparkJobArtifact", string(data)}, " ")
}

type SparkJobArtifactType struct {
	value string
}

type SparkJobArtifactTypeEnum struct {
	SPARK_JAR_JOB           SparkJobArtifactType
	SPARK_PYTHON_JOB        SparkJobArtifactType
	SPARK_SQL_SCRIPTING_JOB SparkJobArtifactType
}

func GetSparkJobArtifactTypeEnum() SparkJobArtifactTypeEnum {
	return SparkJobArtifactTypeEnum{
		SPARK_JAR_JOB: SparkJobArtifactType{
			value: "spark_jar_job",
		},
		SPARK_PYTHON_JOB: SparkJobArtifactType{
			value: "spark_python_job",
		},
		SPARK_SQL_SCRIPTING_JOB: SparkJobArtifactType{
			value: "spark_sql_scripting_job",
		},
	}
}

func (c SparkJobArtifactType) Value() string {
	return c.value
}

func (c SparkJobArtifactType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SparkJobArtifactType) UnmarshalJSON(b []byte) error {
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
