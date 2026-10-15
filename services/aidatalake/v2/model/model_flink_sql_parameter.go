package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkSqlParameter **参数解释**：Flink SQL作业参数，包含SQL作业的脚本和依赖配置。 **约束限制**：不涉及。
type FlinkSqlParameter struct {

	// **参数解释**：作业类型。 **约束限制**：固定值为 FLINK_SQL_JOB。 **取值范围**：不涉及。 **默认取值**：FLINK_SQL_JOB。
	Type *string `json:"type,omitempty"`

	// **参数解释**：SQL脚本所在的OBS文件路径，用于指定SQL作业的脚本文件。 **约束限制**：不涉及。 **取值范围**：OBS URL格式，长度为1~512个字符，以obs://开头。 **默认取值**：不涉及。
	SqlScriptFile string `json:"sql_script_file"`

	// **参数解释**：SQL作业依赖的JAR包OBS路径列表，用于指定作业运行所需的依赖JAR包，单个路径为OBS URL格式，长度为1~512个字符，以obs://开头。 **约束限制**：依赖JAR包数量不能超过100个。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`
}

func (o FlinkSqlParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSqlParameter struct{}"
	}

	return strings.Join([]string{"FlinkSqlParameter", string(data)}, " ")
}
