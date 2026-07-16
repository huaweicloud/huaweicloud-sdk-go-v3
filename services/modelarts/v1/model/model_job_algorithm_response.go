package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobAlgorithmResponse 训练作业算法。目前支持三种形式： - id：只取算法的id； - subscription_id+item_version_id：取算法的订阅id和版本id； - code_dir+boot_file：取训练作业的代码目录和启动文件。
type JobAlgorithmResponse struct {

	// 训练作业算法。目前支持三种形式：   - id：只取算法的id；   - subscription_id+item_version_id：取算法的订阅id和版本id；   - code_dir+boot_file：取训练作业的代码目录和启动文件。
	Id *string `json:"id,omitempty"`

	// 算法名称。
	Name *string `json:"name,omitempty"`

	// 订阅算法的订阅ID。应与item_version_id一同出现。
	SubscriptionId *string `json:"subscription_id,omitempty"`

	// 订阅算法的版本。应与subscription_id一同出现。
	ItemVersionId *string `json:"item_version_id,omitempty"`

	// 训练作业的代码目录。如：“/usr/app/”。应与boot_file一同出现，如果boot_file已经填入id或subscription_id+item_version_id，则无需填写此参数。
	CodeDir *string `json:"code_dir,omitempty"`

	// 训练作业的代码启动文件，需要在代码目录下。如：“/usr/app/boot.py”。应与code_dir一同出现，如果code_dir已经填入id或subscription_id+item_version_id，则无需填写此参数。
	BootFile *string `json:"boot_file,omitempty"`

	// 自动化搜索作业的yaml配置路径，需要提供一个OBS路径。如：“obs://bucket/file.yaml”。
	AutosearchConfigPath *string `json:"autosearch_config_path,omitempty"`

	// 自动化搜索作业的框架代码目录，需要提供一个OBS路径。如：“obs://bucket/files/”。
	AutosearchFrameworkPath *string `json:"autosearch_framework_path,omitempty"`

	// 自定义镜像训练作业的自定义镜像的容器的启动命令。例如python train.py。
	Command *string `json:"command,omitempty"`

	// 训练作业的运行参数。
	Parameters *[]Parameter `json:"parameters,omitempty"`

	Policies *JobAlgorithmResponsePolicies `json:"policies,omitempty"`

	// **参数解释**：训练作业的数据输入。
	Inputs *[]InputResp `json:"inputs,omitempty"`

	// **参数解释**：训练作业的结果输出。
	Outputs *[]OutputResp `json:"outputs,omitempty"`

	Engine *JobEngineResp `json:"engine,omitempty"`

	// 算法的代码目录下载到训练容器内的本地路径。规则如下： - 必须为/home下的目录； - v1兼容模式下，当前字段不生效； - 当code_dir以file://为前缀时，当前字段不生效。
	LocalCodeDir *string `json:"local_code_dir,omitempty"`

	// 运行算法时所在的工作目录。规则：v1兼容模式下，当前字段不生效。
	WorkingDir *string `json:"working_dir,omitempty"`

	// 训练作业的环境变量。格式：\"key\":\"value\"，无需填写。
	Environments *[]map[string]string `json:"environments,omitempty"`

	Summary *SummaryResp `json:"summary,omitempty"`
}

func (o JobAlgorithmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobAlgorithmResponse struct{}"
	}

	return strings.Join([]string{"JobAlgorithmResponse", string(data)}, " ")
}
