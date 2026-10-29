package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkPythonParameter **参数解释**：Flink Python作业参数，包含Python作业的入口和依赖配置。 **约束限制**：不涉及。
type FlinkPythonParameter struct {

	// **参数解释**：作业类型。 **约束限制**：固定值为 FLINK_PYTHON_JOB。 **取值范围**：不涉及。 **默认取值**：FLINK_PYTHON_JOB。
	Type *string `json:"type,omitempty"`

	// **参数解释**：Python作业文件的OBS路径，用于指定Python作业的入口脚本或作业包。 **约束限制**：以 .py 结尾时不允许配置 entry_module；非 .py 时必须配置 entry_module。 **取值范围**：长度为1~512个字符，以obs://开头。 **默认取值**：不涉及。
	PythonJobFile string `json:"python_job_file"`

	// **参数解释**：Python程序入口模块名，用于指定作业包内的入口模块。 **约束限制**：python_job_file 以 .py 结尾时必须为空；非 .py 时必填。 **取值范围**：长度为1~512个字符。 **默认取值**：不涉及。
	EntryModule *string `json:"entry_module,omitempty"`

	// **参数解释**：Python作业入口参数，多个参数之间空格分隔。 **约束限制**：参数数量不能超过100个。
	MainArgs *[]string `json:"main_args,omitempty"`

	// **参数解释**：Python作业依赖JAR包的OBS路径列表，用于指定作业运行所需的依赖JAR包，单个路径为OBS URL格式，长度为1~512个字符，以obs://开头。 **约束限制**：依赖JAR包数量不能超过100个。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// **参数解释**：Python作业依赖文件的OBS路径列表，用于指定作业运行所需的依赖文件，单个路径为OBS URL格式，长度为1~512个字符，以obs://开头。 **约束限制**：依赖文件数量不能超过100个。
	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	// **参数解释**：Python归档文件的OBS路径列表，用于指定Python运行环境或数据文件的压缩包。 **约束限制**：归档文件数量不能超过100个。
	PythonArchives *[]string `json:"python_archives,omitempty"`

	// **参数解释**：第三方Python依赖的OBS路径列表，用于指定Python UDF所需的依赖。 **约束限制**：依赖文件数量不能超过100个。
	PythonLibraries *[]string `json:"python_libraries,omitempty"`
}

func (o FlinkPythonParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkPythonParameter struct{}"
	}

	return strings.Join([]string{"FlinkPythonParameter", string(data)}, " ")
}
