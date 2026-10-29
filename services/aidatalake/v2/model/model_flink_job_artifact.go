package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlinkJobArtifact Flink作业参数，根据作业类型自动选择对应的参数结构。
type FlinkJobArtifact struct {
	Type FlinkJobArtifactType `json:"type"`

	// **参数解释**：SQL脚本所在的OBS文件路径，用于指定SQL作业的脚本文件。 **约束限制**：不涉及。 **取值范围**：OBS URL格式，长度为1~512个字符，以obs://开头。 **默认取值**：不涉及。
	SqlScriptFile string `json:"sql_script_file"`

	// **参数解释**：Python作业依赖JAR包的OBS路径列表，用于指定作业运行所需的依赖JAR包，单个路径为OBS URL格式，长度为1~512个字符，以obs://开头。 **约束限制**：依赖JAR包数量不能超过100个。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// **参数解释**：Python作业依赖文件的OBS路径列表，用于指定作业运行所需的依赖文件，单个路径为OBS URL格式，长度为1~512个字符，以obs://开头。 **约束限制**：依赖文件数量不能超过100个。
	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	PythonParameter *FlinkPythonDependencyParameter `json:"python_parameter,omitempty"`

	// **参数解释**：JAR作业主类所在JAR包的OBS路径，用于指定作业的入口JAR包。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符，以obs://开头。 **默认取值**：不涉及。
	MainJar string `json:"main_jar"`

	// **参数解释**：JAR作业入口类全限定名，用于指定作业的入口类。当主程序包manifest中已指定主类时可省略。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符，字符组成不限制。 **默认取值**：不涉及。
	MainClass *string `json:"main_class,omitempty"`

	// **参数解释**：Python作业入口参数，多个参数之间空格分隔。 **约束限制**：参数数量不能超过100个。
	MainArgs *[]string `json:"main_args,omitempty"`

	// **参数解释**：用户自定义的SQL Gateway访问鉴权Token，调用方通过HTTP请求头 Authorization: Bearer {token} 携带。 **约束限制**：不涉及。 **取值范围**：长度为1~256个字符。 **默认取值**：不涉及。
	SqlGatewayToken *string `json:"sql_gateway_token,omitempty"`

	// **参数解释**：是否开启空闲自动释放。 **约束限制**：不涉及。 **取值范围**：   - true：开启空闲自动释放。   - false：关闭空闲自动释放。 **默认取值**：false。
	IdleEnabled *bool `json:"idle_enabled,omitempty"`

	// **参数解释**：最大空闲时长，单位为分钟，用于指定SQL Gateway空闲多长时间后自动停止。空闲指会话内没有正在运行的作业，创建会话等其他操作不计为非空闲。 **约束限制**：不涉及。 **取值范围**：10~1440。 **默认取值**：30。
	IdleTimeout *int32 `json:"idle_timeout,omitempty"`

	// **参数解释**：Python作业文件的OBS路径，用于指定Python作业的入口脚本或作业包。 **约束限制**：以 .py 结尾时不允许配置 entry_module；非 .py 时必须配置 entry_module。 **取值范围**：长度为1~512个字符，以obs://开头。 **默认取值**：不涉及。
	PythonJobFile string `json:"python_job_file"`

	// **参数解释**：Python程序入口模块名，用于指定作业包内的入口模块。 **约束限制**：python_job_file 以 .py 结尾时必须为空；非 .py 时必填。 **取值范围**：长度为1~512个字符。 **默认取值**：不涉及。
	EntryModule *string `json:"entry_module,omitempty"`

	// **参数解释**：Python归档文件的OBS路径列表，用于指定Python运行环境或数据文件的压缩包。 **约束限制**：归档文件数量不能超过100个。
	PythonArchives *[]string `json:"python_archives,omitempty"`

	// **参数解释**：第三方Python依赖的OBS路径列表，用于指定Python UDF所需的依赖。 **约束限制**：依赖文件数量不能超过100个。
	PythonLibraries *[]string `json:"python_libraries,omitempty"`
}

func (o FlinkJobArtifact) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobArtifact struct{}"
	}

	return strings.Join([]string{"FlinkJobArtifact", string(data)}, " ")
}

type FlinkJobArtifactType struct {
	value string
}

type FlinkJobArtifactTypeEnum struct {
	FLINK_SQL_JOB     FlinkJobArtifactType
	FLINK_JAR_JOB     FlinkJobArtifactType
	FLINK_SQL_GATEWAY FlinkJobArtifactType
	FLINK_PYTHON_JOB  FlinkJobArtifactType
}

func GetFlinkJobArtifactTypeEnum() FlinkJobArtifactTypeEnum {
	return FlinkJobArtifactTypeEnum{
		FLINK_SQL_JOB: FlinkJobArtifactType{
			value: "FLINK_SQL_JOB",
		},
		FLINK_JAR_JOB: FlinkJobArtifactType{
			value: "FLINK_JAR_JOB",
		},
		FLINK_SQL_GATEWAY: FlinkJobArtifactType{
			value: "FLINK_SQL_GATEWAY",
		},
		FLINK_PYTHON_JOB: FlinkJobArtifactType{
			value: "FLINK_PYTHON_JOB",
		},
	}
}

func (c FlinkJobArtifactType) Value() string {
	return c.value
}

func (c FlinkJobArtifactType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlinkJobArtifactType) UnmarshalJSON(b []byte) error {
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
