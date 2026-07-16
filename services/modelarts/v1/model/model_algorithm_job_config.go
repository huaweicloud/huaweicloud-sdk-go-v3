package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmJobConfig 算法配置信息，如启动文件等。
type AlgorithmJobConfig struct {

	// 算法的代码目录。如：“/usr/app/”。应与boot_file一同出现。
	CodeDir *string `json:"code_dir,omitempty"`

	// 算法的代码启动文件，需要在代码目录下。如：“/usr/app/boot.py”。应与code_dir一同出现。
	BootFile *string `json:"boot_file,omitempty"`

	// 自定义镜像算法的容器启动命令。
	Command *string `json:"command,omitempty"`

	// 算法的运行参数。
	Parameters *[]Parameters `json:"parameters,omitempty"`

	// 算法的数据输入。
	Inputs *[]AlgorithmCreateInput `json:"inputs,omitempty"`

	// 算法的数据输出。
	Outputs *[]AlgorithmCreateOutput `json:"outputs,omitempty"`

	Engine *AlgorithmCreateEngine `json:"engine,omitempty"`

	// 算法是否允许创建训练作业时自定义超参。
	ParametersCustomization *bool `json:"parameters_customization,omitempty"`
}

func (o AlgorithmJobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmJobConfig struct{}"
	}

	return strings.Join([]string{"AlgorithmJobConfig", string(data)}, " ")
}
