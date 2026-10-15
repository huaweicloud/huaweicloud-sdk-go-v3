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

	// **参数解释**：JAR作业依赖JAR包的OBS路径列表，用于指定作业运行所需的依赖JAR包，单个路径为OBS URL格式，长度为1~512个字符，以obs://开头。 **约束限制**：依赖JAR包数量不能超过100个。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// **参数解释**：JAR作业主类所在JAR包的OBS路径，用于指定作业的入口JAR包。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符，以obs://开头。 **默认取值**：不涉及。
	MainJar string `json:"main_jar"`

	// **参数解释**：JAR作业入口类全限定名，用于指定作业的入口类。当主程序包manifest中已指定主类时可省略。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符，字符组成不限制。 **默认取值**：不涉及。
	MainClass *string `json:"main_class,omitempty"`

	// **参数解释**：JAR作业入口类参数列表，按数组顺序传入。 **约束限制**：参数数量不能超过100个。
	MainArgs *[]string `json:"main_args,omitempty"`

	// **参数解释**：JAR作业依赖文件的OBS路径列表，用于指定作业运行所需的依赖文件，单个路径为OBS URL格式，长度为1~512个字符，以obs://开头。 **约束限制**：依赖文件数量不能超过100个。
	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	// **参数解释**：用户自定义的SQL Gateway访问鉴权Token，调用方通过HTTP请求头 Authorization: Bearer {token} 携带。 **约束限制**：不涉及。 **取值范围**：长度为1~256个字符。 **默认取值**：不涉及。
	SqlGatewayToken *string `json:"sql_gateway_token,omitempty"`
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
