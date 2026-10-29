package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFlinkJobArtifactResponse Flink作业参数响应，根据作业类型自动选择对应的参数结构。
type ShowFlinkJobArtifactResponse struct {
	Type ShowFlinkJobArtifactResponseType `json:"type"`

	// **参数解释**：SQL脚本所在的OBS文件路径，用于指定SQL作业的脚本文件。 **取值范围**：OBS URL格式，长度为1~512个字符，以obs://开头。
	SqlScriptFile *string `json:"sql_script_file,omitempty"`

	// **参数解释**：Python作业依赖的第三方Java JAR包OBS路径列表。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// **参数解释**：Python作业附加依赖文件的OBS路径列表。
	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	PythonParameter *ShowFlinkPythonDependencyParameterResponse `json:"python_parameter,omitempty"`

	// **参数解释**：JAR作业主类所在JAR包的OBS路径，用于指定作业的入口JAR包。 **取值范围**：长度为1~512个字符，以obs://开头。
	MainJar *string `json:"main_jar,omitempty"`

	// **参数解释**：JAR作业入口类名称，用于指定作业的入口类。 **取值范围**：长度不超过512个字符。
	MainClass *string `json:"main_class,omitempty"`

	// **参数解释**：Python作业入口参数，用于传递运行参数。
	MainArgs *[]string `json:"main_args,omitempty"`

	SqlGatewayConfig *ShowFlinkSqlGatewayConfigResponseBody `json:"sql_gateway_config,omitempty"`

	// **参数解释**：是否开启空闲自动释放。 **取值范围**：   - true：开启空闲自动释放。   - false：关闭空闲自动释放。
	IdleEnabled *bool `json:"idle_enabled,omitempty"`

	// **参数解释**：最大空闲时长，单位为分钟。空闲指会话内没有正在运行的作业，创建会话等其他操作不计为非空闲。 **取值范围**：10~1440。
	IdleTimeout *int32 `json:"idle_timeout,omitempty"`

	// **参数解释**：Python作业文件的OBS路径，用于指定Python作业的入口脚本或作业包。 **取值范围**：长度为1~512个字符，以obs://开头。
	PythonJobFile *string `json:"python_job_file,omitempty"`

	// **参数解释**：Python程序入口模块名，用于指定作业包内的入口模块。 **取值范围**：长度为1~512个字符。
	EntryModule *string `json:"entry_module,omitempty"`

	// **参数解释**：Python归档文件的OBS路径列表，用于指定Python运行环境或数据文件的压缩包。
	PythonArchives *[]string `json:"python_archives,omitempty"`

	// **参数解释**：第三方Python依赖的OBS路径列表，用于指定Python UDF所需的依赖。
	PythonLibraries *[]string `json:"python_libraries,omitempty"`
}

func (o ShowFlinkJobArtifactResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkJobArtifactResponse struct{}"
	}

	return strings.Join([]string{"ShowFlinkJobArtifactResponse", string(data)}, " ")
}

type ShowFlinkJobArtifactResponseType struct {
	value string
}

type ShowFlinkJobArtifactResponseTypeEnum struct {
	FLINK_SQL_JOB     ShowFlinkJobArtifactResponseType
	FLINK_JAR_JOB     ShowFlinkJobArtifactResponseType
	FLINK_SQL_GATEWAY ShowFlinkJobArtifactResponseType
	FLINK_PYTHON_JOB  ShowFlinkJobArtifactResponseType
}

func GetShowFlinkJobArtifactResponseTypeEnum() ShowFlinkJobArtifactResponseTypeEnum {
	return ShowFlinkJobArtifactResponseTypeEnum{
		FLINK_SQL_JOB: ShowFlinkJobArtifactResponseType{
			value: "FLINK_SQL_JOB",
		},
		FLINK_JAR_JOB: ShowFlinkJobArtifactResponseType{
			value: "FLINK_JAR_JOB",
		},
		FLINK_SQL_GATEWAY: ShowFlinkJobArtifactResponseType{
			value: "FLINK_SQL_GATEWAY",
		},
		FLINK_PYTHON_JOB: ShowFlinkJobArtifactResponseType{
			value: "FLINK_PYTHON_JOB",
		},
	}
}

func (c ShowFlinkJobArtifactResponseType) Value() string {
	return c.value
}

func (c ShowFlinkJobArtifactResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlinkJobArtifactResponseType) UnmarshalJSON(b []byte) error {
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
