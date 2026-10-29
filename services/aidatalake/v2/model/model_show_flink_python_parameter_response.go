package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFlinkPythonParameterResponse **参数解释**：Flink Python作业参数，包含Python作业的详细配置信息。
type ShowFlinkPythonParameterResponse struct {

	// **参数解释**：作业类型。 **约束限制**：不涉及。 **取值范围**：   - FLINK_PYTHON_JOB：Flink Python作业。
	Type *ShowFlinkPythonParameterResponseType `json:"type,omitempty"`

	// **参数解释**：Python作业文件的OBS路径，用于指定Python作业的入口脚本或作业包。 **取值范围**：长度为1~512个字符，以obs://开头。
	PythonJobFile *string `json:"python_job_file,omitempty"`

	// **参数解释**：Python程序入口模块名，用于指定作业包内的入口模块。 **取值范围**：长度为1~512个字符。
	EntryModule *string `json:"entry_module,omitempty"`

	// **参数解释**：Python作业入口参数，用于传递运行参数。
	MainArgs *[]string `json:"main_args,omitempty"`

	// **参数解释**：Python作业依赖的第三方Java JAR包OBS路径列表。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// **参数解释**：Python作业附加依赖文件的OBS路径列表。
	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	// **参数解释**：Python归档文件的OBS路径列表，用于指定Python运行环境或数据文件的压缩包。
	PythonArchives *[]string `json:"python_archives,omitempty"`

	// **参数解释**：第三方Python依赖的OBS路径列表，用于指定Python UDF所需的依赖。
	PythonLibraries *[]string `json:"python_libraries,omitempty"`
}

func (o ShowFlinkPythonParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkPythonParameterResponse struct{}"
	}

	return strings.Join([]string{"ShowFlinkPythonParameterResponse", string(data)}, " ")
}

type ShowFlinkPythonParameterResponseType struct {
	value string
}

type ShowFlinkPythonParameterResponseTypeEnum struct {
	FLINK_PYTHON_JOB ShowFlinkPythonParameterResponseType
}

func GetShowFlinkPythonParameterResponseTypeEnum() ShowFlinkPythonParameterResponseTypeEnum {
	return ShowFlinkPythonParameterResponseTypeEnum{
		FLINK_PYTHON_JOB: ShowFlinkPythonParameterResponseType{
			value: "FLINK_PYTHON_JOB",
		},
	}
}

func (c ShowFlinkPythonParameterResponseType) Value() string {
	return c.value
}

func (c ShowFlinkPythonParameterResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlinkPythonParameterResponseType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
