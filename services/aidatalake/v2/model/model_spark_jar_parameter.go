package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkJarParameter **参数解释**：Spark Jar作业参数，用于配置Spark Jar作业的执行参数和依赖资源。 **约束限制**：不涉及。
type SparkJarParameter struct {

	// **参数解释**：作业类型。 **约束限制**：固定值为 spark_jar_job。 **取值范围**：不涉及。 **默认取值**：spark_jar_job。
	Type *string `json:"type,omitempty"`

	// **参数解释**：主类名称，用于指定Spark Jar作业的入口类。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符。 **默认取值**：不涉及。
	MainClass *string `json:"main_class,omitempty"`

	// **参数解释**：主类参数列表，用于传递给Spark Jar作业入口类的参数。多个参数之间空格分隔。 **约束限制**：参数数量不能超过100个。
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
}

func (o SparkJarParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJarParameter struct{}"
	}

	return strings.Join([]string{"SparkJarParameter", string(data)}, " ")
}
