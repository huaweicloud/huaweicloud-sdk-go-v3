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

	// **参数解释**：JAR作业依赖JAR包的OBS路径，用于指定作业运行所需的依赖JAR包。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// **参数解释**：JAR作业主类所在JAR包的OBS路径，用于指定作业的入口JAR包。 **取值范围**：长度为1~512个字符，以obs://开头。
	MainJar *string `json:"main_jar,omitempty"`

	// **参数解释**：JAR作业入口类名称，用于指定作业的入口类。 **取值范围**：长度不超过512个字符。
	MainClass *string `json:"main_class,omitempty"`

	// **参数解释**：JAR作业入口类参数，用于传递运行参数。
	MainArgs *[]string `json:"main_args,omitempty"`

	// **参数解释**：JAR作业依赖文件的OBS路径列表，用于指定作业运行所需的依赖文件。
	DependencyFiles *[]string `json:"dependency_files,omitempty"`

	SqlGatewayConfig *ShowFlinkSqlGatewayConfigResponseBody `json:"sql_gateway_config,omitempty"`
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
