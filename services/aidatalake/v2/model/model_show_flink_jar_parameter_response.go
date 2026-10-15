package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFlinkJarParameterResponse **参数解释**：Flink JAR作业参数，包含JAR作业的详细配置信息。
type ShowFlinkJarParameterResponse struct {

	// **参数解释**：作业类型。 **约束限制**：不涉及。 **取值范围**：   - FLINK_JAR_JOB：Flink Jar作业。
	Type *ShowFlinkJarParameterResponseType `json:"type,omitempty"`

	// **参数解释**：JAR作业主类所在JAR包的OBS路径，用于指定作业的入口JAR包。 **取值范围**：长度为1~512个字符，以obs://开头。
	MainJar *string `json:"main_jar,omitempty"`

	// **参数解释**：JAR作业入口类名称，用于指定作业的入口类。 **取值范围**：长度不超过512个字符。
	MainClass *string `json:"main_class,omitempty"`

	// **参数解释**：JAR作业入口类参数，用于传递运行参数。
	MainArgs *[]string `json:"main_args,omitempty"`

	// **参数解释**：JAR作业依赖JAR包的OBS路径，用于指定作业运行所需的依赖JAR包。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// **参数解释**：JAR作业依赖文件的OBS路径列表，用于指定作业运行所需的依赖文件。
	DependencyFiles *[]string `json:"dependency_files,omitempty"`
}

func (o ShowFlinkJarParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkJarParameterResponse struct{}"
	}

	return strings.Join([]string{"ShowFlinkJarParameterResponse", string(data)}, " ")
}

type ShowFlinkJarParameterResponseType struct {
	value string
}

type ShowFlinkJarParameterResponseTypeEnum struct {
	FLINK_JAR_JOB ShowFlinkJarParameterResponseType
}

func GetShowFlinkJarParameterResponseTypeEnum() ShowFlinkJarParameterResponseTypeEnum {
	return ShowFlinkJarParameterResponseTypeEnum{
		FLINK_JAR_JOB: ShowFlinkJarParameterResponseType{
			value: "FLINK_JAR_JOB",
		},
	}
}

func (c ShowFlinkJarParameterResponseType) Value() string {
	return c.value
}

func (c ShowFlinkJarParameterResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlinkJarParameterResponseType) UnmarshalJSON(b []byte) error {
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
