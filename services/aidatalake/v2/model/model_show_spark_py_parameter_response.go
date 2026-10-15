package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkPyParameterResponse **参数解释**：Spark Python作业参数，用于配置Python Spark作业的执行参数和依赖资源。
type ShowSparkPyParameterResponse struct {

	// **参数解释**：主Python文件路径，用于指定Python Spark作业的主文件OBS路径。 **取值范围**：OBS URL格式，长度为1~512个字符，例如：obs://bucket/pyspark/pySpark_udf_python.py.
	MainPythonFile *string `json:"main_python_file,omitempty"`

	// **参数解释**：Spark Python作业入口类参数列表，用于传递作业执行参数。
	MainArgs *[]string `json:"main_args,omitempty"`

	// **参数解释**：依赖Jar包列表，用于指定Spark作业依赖的Jar包OBS路径。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// **参数解释**：依赖文件列表，用于指定Spark作业依赖的文件OBS路径。
	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	// **参数解释**：依赖归档包列表，用于指定Spark作业依赖的归档包OBS路径。
	DependencyArchives *[]string `json:"dependency_archives,omitempty"`

	// **参数解释**：依赖Python文件列表，用于指定Spark作业依赖的Python包OBS路径。
	DependencyPyFiles *[]string `json:"dependency_py_files,omitempty"`
}

func (o ShowSparkPyParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkPyParameterResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkPyParameterResponse", string(data)}, " ")
}
