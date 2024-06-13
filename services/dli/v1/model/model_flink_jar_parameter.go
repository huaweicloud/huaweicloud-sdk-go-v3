package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkJarParameter Flink Jar 作业参数。
type FlinkJarParameter struct {

	// Flink Jar 作业入口类。
	MainClass *string `json:"main_class,omitempty"`

	// Flink Jar 作业入口类参数，多个参数之间空格分隔。
	MainArgs *string `json:"main_args,omitempty"`

	// Flink Jar 作业主类所在 Jar 包的 OBS 路径。
	MainJar *string `json:"main_jar,omitempty"`

	// Flink Jar 作业依赖 Jar 包的 OBS 路径数组。示例：[obs://bucket/demo/test1.jar,obs://bucket/demo/test2.jar]
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// Flink Jar 作业依赖文件的 OBS 路径数组。示例：[obs://bucket/demo/test1.csv,obs://bucket/demo/test2.csv]
	DependencyFiles *[]string `json:"dependency_files,omitempty"`
}

func (o FlinkJarParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJarParameter struct{}"
	}

	return strings.Join([]string{"FlinkJarParameter", string(data)}, " ")
}
