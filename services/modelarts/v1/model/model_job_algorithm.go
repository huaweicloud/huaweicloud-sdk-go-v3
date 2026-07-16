package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobAlgorithm 训练作业算法。目前支持三种形式： - id：只取算法的id； - subscription_id+item_version_id：取算法的订阅id和版本id； - code_dir+boot_file：取训练作业的代码目录和启动文件。
type JobAlgorithm struct {

	// **参数解释**：算法管理的算法id。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：算法名称。无需填写。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：订阅算法的订阅ID。 **约束限制**：应与item_version_id一同出现。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SubscriptionId *string `json:"subscription_id,omitempty"`

	// **参数解释**：订阅算法的版本。 **约束限制**：应与subscription_id一同出现。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ItemVersionId *string `json:"item_version_id,omitempty"`

	// **参数解释**：训练作业的代码目录。如：“/usr/app/”。 **约束限制**：应与boot_file一同出现，如果boot_file填入id或subscription_id+item_version_id，则此参数无需填写。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CodeDir *string `json:"code_dir,omitempty"`

	// **参数解释**：训练作业的代码启动文件，需要在代码目录下。如：“/usr/app/boot.py”。 **约束限制**：应与code_dir一同出现，如果code_dir填入id或subscription_id+item_version_id，则此参数无需填写。 **取值范围**：不涉及。 **默认取值**：不涉及。
	BootFile *string `json:"boot_file,omitempty"`

	// **参数解释**：自动化搜索作业的yaml配置路径，需要提供一个OBS路径。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AutosearchConfigPath *string `json:"autosearch_config_path,omitempty"`

	// **参数解释**：自动化搜索作业的框架代码目录，需要提供一个OBS路径。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AutosearchFrameworkPath *string `json:"autosearch_framework_path,omitempty"`

	// **参数解释**：自定义镜像场景下，训练作业的自定义镜像的容器的启动命令。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Command *string `json:"command,omitempty"`

	// **参数解释**：训练作业的运行参数。 **约束限制**：不涉及。
	Parameters *[]Parameters `json:"parameters,omitempty"`

	Policies *JobPolicies `json:"policies,omitempty"`

	// **参数解释**：训练作业的数据输入。 **约束限制**：不涉及。
	Inputs *[]Input `json:"inputs,omitempty"`

	// **参数解释**：训练作业的结果输出。 **约束限制**：不涉及。
	Outputs *[]Output `json:"outputs,omitempty"`

	Engine *JobEngine `json:"engine,omitempty"`

	// **参数解释**：算法的代码目录下载到训练容器内的本地路径。 **约束限制**： - 必须为/home下的目录。 - v1兼容模式下，当前字段不生效。 - 当code_dir以file://为前缀时，当前字段不生效。 - 不支持配置成/home/ma-user/modelarts，/home/ma-user/modelarts-dev，/home/ma-user/infer以及它们底下的目录，也不支持配置成/home/ma-user  **取值范围**：不涉及。 **默认取值**：不涉及。
	LocalCodeDir *string `json:"local_code_dir,omitempty"`

	// **参数解释**：运行算法时所在的工作目录。 **约束限制**：v1兼容模式下，当前字段不生效。 **取值范围**：不涉及。 **默认取值**：不涉及。
	WorkingDir *string `json:"working_dir,omitempty"`

	// **参数解释**：训练作业的环境变量。格式：\"key\":\"value\"。 **约束限制**：其中key最大允许填写8192字符，value最大允许填写4096字符，最多允许100对环境变量。变量名应该仅包含字母、数字、下划线，且以字母或下划线开头。 注：不支持使用符号 $ 引用变量。
	Environments map[string]string `json:"environments,omitempty"`

	Summary *Summary `json:"summary,omitempty"`
}

func (o JobAlgorithm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobAlgorithm struct{}"
	}

	return strings.Join([]string{"JobAlgorithm", string(data)}, " ")
}
