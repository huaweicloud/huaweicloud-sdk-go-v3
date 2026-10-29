package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkPythonDependencyParameter **参数解释**：JAR/SQL 作业的 Python UDF 依赖。 **约束限制**：不涉及。
type FlinkPythonDependencyParameter struct {

	// **参数解释**：Python归档文件的OBS路径列表，用于指定Python运行环境或数据文件的压缩包。 **约束限制**：归档文件数量不能超过100个。单个路径为OBS URL格式，长度为1~512个字符，以obs://开头。
	PythonArchives *[]string `json:"python_archives,omitempty"`

	// **参数解释**：第三方Python依赖的OBS路径列表，用于指定Python UDF所需的依赖。 **约束限制**：依赖文件数量不能超过100个。单个路径为OBS URL格式，长度为1~512个字符，以obs://开头。
	PythonLibraries *[]string `json:"python_libraries,omitempty"`
}

func (o FlinkPythonDependencyParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkPythonDependencyParameter struct{}"
	}

	return strings.Join([]string{"FlinkPythonDependencyParameter", string(data)}, " ")
}
