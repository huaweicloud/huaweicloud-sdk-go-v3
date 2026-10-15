package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkJarParameterResponse **参数解释**：Spark Jar作业参数，用于配置Jar作业的执行参数和依赖资源。
type ShowSparkJarParameterResponse struct {

	// **参数解释**：Spark Jar作业入口类，用于指定作业的主类名称。 **取值范围**：完整的Java类路径格式，例如：com.example.SparkJob。
	MainClass *string `json:"main_class,omitempty"`

	// **参数解释**：Spark Jar作业入口类参数列表，用于传递作业执行参数。
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
}

func (o ShowSparkJarParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJarParameterResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkJarParameterResponse", string(data)}, " ")
}
