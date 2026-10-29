package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFlinkSqlParameterResponse **参数解释**：SQL作业参数，包含SQL作业的详细配置信息。
type ShowFlinkSqlParameterResponse struct {

	// **参数解释**：作业类型，用于区分SQL作业、Jar作业和SQL Gateway作业。 **约束限制**：不涉及。 **取值范围**：   - FLINK_SQL_JOB：SQL作业。
	Type *ShowFlinkSqlParameterResponseType `json:"type,omitempty"`

	// **参数解释**：SQL脚本所在的OBS文件路径，用于指定SQL作业的脚本文件。 **取值范围**：OBS URL格式，长度为1~512个字符，以obs://开头。
	SqlScriptFile *string `json:"sql_script_file,omitempty"`

	// **参数解释**：SQL作业依赖的JAR包OBS路径列表，用于指定作业运行所需的依赖JAR包，单个路径为OBS URL格式，长度为1~512个字符，以obs://开头。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// **参数解释**：SQL作业依赖文件的OBS路径列表，用于指定作业运行所需的依赖文件。
	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	PythonParameter *ShowFlinkPythonDependencyParameterResponse `json:"python_parameter,omitempty"`
}

func (o ShowFlinkSqlParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkSqlParameterResponse struct{}"
	}

	return strings.Join([]string{"ShowFlinkSqlParameterResponse", string(data)}, " ")
}

type ShowFlinkSqlParameterResponseType struct {
	value string
}

type ShowFlinkSqlParameterResponseTypeEnum struct {
	FLINK_SQL_JOB ShowFlinkSqlParameterResponseType
}

func GetShowFlinkSqlParameterResponseTypeEnum() ShowFlinkSqlParameterResponseTypeEnum {
	return ShowFlinkSqlParameterResponseTypeEnum{
		FLINK_SQL_JOB: ShowFlinkSqlParameterResponseType{
			value: "FLINK_SQL_JOB",
		},
	}
}

func (c ShowFlinkSqlParameterResponseType) Value() string {
	return c.value
}

func (c ShowFlinkSqlParameterResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlinkSqlParameterResponseType) UnmarshalJSON(b []byte) error {
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
