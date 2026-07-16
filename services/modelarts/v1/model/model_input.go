package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Input struct {

	// 数据输入通道名称。
	Name string `json:"name"`

	// 数据输入通道描述信息。
	Description *string `json:"description,omitempty"`

	// 数据输入通道映射的容器本地路径。例如，“/home/ma-user/modelarts/inputs/data_url_0”。
	LocalDir *string `json:"local_dir,omitempty"`

	// **参数解释**：数据输入通道路径（local_dir）的下发方式。 **约束限制**：不涉及。 **取值范围**： - parameter：超参形式 - env：环境变量形式  **默认取值**：默认超参形式。
	AccessMethod *string `json:"access_method,omitempty"`

	Remote *InputDataInfo `json:"remote"`

	// 数据输入约束。
	RemoteConstraint *[]InputRemoteConstraint `json:"remote_constraint,omitempty"`
}

func (o Input) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Input struct{}"
	}

	return strings.Join([]string{"Input", string(data)}, " ")
}
