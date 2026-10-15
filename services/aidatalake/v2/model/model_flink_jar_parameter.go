package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkJarParameter **参数解释**：Flink JAR作业参数，包含JAR作业的入口和依赖配置。 **约束限制**：不涉及。
type FlinkJarParameter struct {

	// **参数解释**：作业类型。 **约束限制**：固定值为 FLINK_JAR_JOB。 **取值范围**：不涉及。 **默认取值**：FLINK_JAR_JOB。
	Type *string `json:"type,omitempty"`

	// **参数解释**：JAR作业主类所在JAR包的OBS路径，用于指定作业的入口JAR包。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符，以obs://开头。 **默认取值**：不涉及。
	MainJar string `json:"main_jar"`

	// **参数解释**：JAR作业入口类全限定名，用于指定作业的入口类。当主程序包manifest中已指定主类时可省略。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符，字符组成不限制。 **默认取值**：不涉及。
	MainClass *string `json:"main_class,omitempty"`

	// **参数解释**：JAR作业入口类参数列表，按数组顺序传入。 **约束限制**：参数数量不能超过100个。
	MainArgs *[]string `json:"main_args,omitempty"`

	// **参数解释**：JAR作业依赖JAR包的OBS路径列表，用于指定作业运行所需的依赖JAR包，单个路径为OBS URL格式，长度为1~512个字符，以obs://开头。 **约束限制**：依赖JAR包数量不能超过100个。
	DependencyJars *[]string `json:"dependency_jars,omitempty"`

	// **参数解释**：JAR作业依赖文件的OBS路径列表，用于指定作业运行所需的依赖文件，单个路径为OBS URL格式，长度为1~512个字符，以obs://开头。 **约束限制**：依赖文件数量不能超过100个。
	DependencyFiles *[]string `json:"dependency_files,omitempty"`
}

func (o FlinkJarParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJarParameter struct{}"
	}

	return strings.Join([]string{"FlinkJarParameter", string(data)}, " ")
}
