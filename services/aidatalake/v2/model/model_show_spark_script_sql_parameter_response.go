package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkScriptSqlParameterResponse **参数解释**：Spark Script SQL作业参数，用于配置SQL脚本作业的执行参数和依赖资源。
type ShowSparkScriptSqlParameterResponse struct {

	// **参数解释**：Spark Script SQL类型作业的SQL脚本文件OBS路径，用于指定作业执行的SQL脚本。 **取值范围**：OBS URL格式，长度为1~1024个字符，例如：obs://bucket/sparksql/script.sql。
	SqlScriptingFile string `json:"sql_scripting_file"`

	// **参数解释**：SQL脚本占位符参数列表，用于为SQL脚本中的占位符传递参数值。
	SqlScriptingParameters *[]SparkJobSqlScriptParameter `json:"sql_scripting_parameters,omitempty"`

	// **参数解释**：依赖Jar包列表，用于指定Spark作业依赖的Jar包OBS路径。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// **参数解释**：SQL脚本结果是否写入OBS，用于控制查询结果的输出方式。 **取值范围**： - true：结果写入OBS。 - false：结果不写入OBS。
	SqlScriptingResultToObs *bool `json:"sql_scripting_result_to_obs,omitempty"`

	Result *SparkSqlScriptingResultResponse `json:"result,omitempty"`
}

func (o ShowSparkScriptSqlParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkScriptSqlParameterResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkScriptSqlParameterResponse", string(data)}, " ")
}
