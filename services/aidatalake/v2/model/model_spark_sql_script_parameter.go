package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkSqlScriptParameter **参数解释**：Spark SQL脚本作业参数，用于配置SQL脚本作业的执行参数和依赖资源。 **约束限制**：不涉及。
type SparkSqlScriptParameter struct {

	// **参数解释**：作业类型。 **约束限制**：固定值为 spark_sql_scripting_job。 **取值范围**：不涉及。 **默认取值**：spark_sql_scripting_job。
	Type *string `json:"type,omitempty"`

	// **参数解释**：SQL脚本文件路径，用于指定Spark Script SQL作业的脚本文件OBS路径。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符，格式为OBS路径，例如：“obs://bucket/sparksqlscript/script.sql”。 **默认取值**：不涉及。
	SqlScriptingFile string `json:"sql_scripting_file"`

	// **参数解释**：SQL脚本占位符参数列表，用于配置SQL脚本中的参数化变量。数组中的每个元素为SparkJobSqlScriptParameter对象，包含占位符的键、值和类型信息。 **约束限制**：占位符参数数量不能超过16条。
	SqlScriptingParameters *[]SparkJobSqlScriptParameter `json:"sql_scripting_parameters,omitempty"`

	// **参数解释**：依赖Jar包列表，用于指定Spark作业依赖的Jar包OBS路径。 **约束限制**：依赖Jar包数量不能超过100个。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// **参数解释**：结果写入OBS标识，用于控制SQL脚本作业的执行结果是否写入OBS。 **约束限制**：不涉及。 **取值范围**： - true：结果写入OBS。 - false：结果不写入OBS。 **默认取值**：不涉及。
	SqlScriptingResultToObs *bool `json:"sql_scripting_result_to_obs,omitempty"`
}

func (o SparkSqlScriptParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkSqlScriptParameter struct{}"
	}

	return strings.Join([]string{"SparkSqlScriptParameter", string(data)}, " ")
}
