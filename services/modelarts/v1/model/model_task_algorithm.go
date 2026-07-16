package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskAlgorithm **参数解释**：算法管理算法配置。 **约束限制**：不涉及。
type TaskAlgorithm struct {
	JobConfig *TaskAlgorithmJobConfig `json:"job_config,omitempty"`

	// **参数解释**：算法的代码目录。如：“/usr/app/”。 **约束限制**：应与boot_file一同出现。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CodeDir *string `json:"code_dir,omitempty"`

	// **参数解释**：算法的代码启动文件，需要在代码目录下。如：“/usr/app/boot.py”。 **约束限制**：应与code_dir一同出现。 **取值范围**：不涉及。 **默认取值**：不涉及。
	BootFile *string `json:"boot_file,omitempty"`

	Engine *TaskAlgorithmEngine `json:"engine,omitempty"`

	// **参数解释**：算法的数据输入。 **约束限制**：不涉及。
	Inputs *[]TaskAlgorithmInputs `json:"inputs,omitempty"`

	// **参数解释**：算法的数据输出。 **约束限制**：不涉及。
	Outputs *[]TaskAlgorithmOutputs `json:"outputs,omitempty"`

	// **参数解释**：算法的代码目录下载到训练容器内的本地路径。 **约束限制**： - 必须为/home下的目录； - v1兼容模式下，当前字段不生效； - 当code_dir以file://为前缀时，当前字段不生效。  **取值范围**：不涉及。 **默认取值**：不涉及。
	LocalCodeDir *string `json:"local_code_dir,omitempty"`

	// **参数解释**：运行算法时所在的工作目录。 **约束限制**：v1兼容模式下，当前字段不生效。 **取值范围**：不涉及。 **默认取值**：不涉及。
	WorkingDir *string `json:"working_dir,omitempty"`

	// **参数解释**：训练作业环境变量。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Environments map[string]string `json:"environments,omitempty"`
}

func (o TaskAlgorithm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskAlgorithm struct{}"
	}

	return strings.Join([]string{"TaskAlgorithm", string(data)}, " ")
}
