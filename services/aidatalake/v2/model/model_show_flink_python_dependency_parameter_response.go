package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkPythonDependencyParameterResponse **参数解释**：JAR/SQL 作业的 Python UDF 依赖。
type ShowFlinkPythonDependencyParameterResponse struct {

	// **参数解释**：Python归档文件的OBS路径列表，用于指定Python运行环境或数据文件的压缩包。
	PythonArchives *[]string `json:"python_archives,omitempty"`

	// **参数解释**：第三方Python依赖的OBS路径列表，用于指定Python UDF所需的依赖。
	PythonLibraries *[]string `json:"python_libraries,omitempty"`
}

func (o ShowFlinkPythonDependencyParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkPythonDependencyParameterResponse struct{}"
	}

	return strings.Join([]string{"ShowFlinkPythonDependencyParameterResponse", string(data)}, " ")
}
