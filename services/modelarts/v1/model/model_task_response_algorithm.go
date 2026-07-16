package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskResponseAlgorithm 算法管理算法配置。
type TaskResponseAlgorithm struct {

	// 算法启动文件所在目录绝对路径。
	CodeDir *string `json:"code_dir,omitempty"`

	// 算法启动文件绝对路径。
	BootFile *string `json:"boot_file,omitempty"`

	Inputs *AlgorithmInput `json:"inputs,omitempty"`

	Outputs *AlgorithmOutput `json:"outputs,omitempty"`

	Engine *AlgorithmEngine `json:"engine,omitempty"`

	// 算法的代码目录下载到训练容器内的本地路径。规则如下： - 必须为/home下的目录； - v1兼容模式下，当前字段不生效； - 当code_dir以file://为前缀时，当前字段不生效。
	LocalCodeDir *string `json:"local_code_dir,omitempty"`

	// 运行算法时所在的工作目录。规则：v1兼容模式下，当前字段不生效。
	WorkingDir *string `json:"working_dir,omitempty"`

	// **参数解释**：训练作业相关的环境变量。 **取值范围**：不涉及。
	Environments map[string]string `json:"environments,omitempty"`
}

func (o TaskResponseAlgorithm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskResponseAlgorithm struct{}"
	}

	return strings.Join([]string{"TaskResponseAlgorithm", string(data)}, " ")
}
