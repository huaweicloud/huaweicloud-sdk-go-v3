package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Output struct {

	// 数据输出通道名称。
	Name string `json:"name"`

	// 数据输出通道描述信息。
	Description *string `json:"description,omitempty"`

	// 数据输出通道映射的容器本地路径。
	LocalDir *string `json:"local_dir,omitempty"`

	// **参数解释**：数据输出通道路径（local_dir）的下发方式。 **约束限制**：不涉及。 **取值范围**： - parameter：超参形式 - env：环境变量形式  **默认取值**：默认超参形式。
	AccessMethod *string `json:"access_method,omitempty"`

	Remote *Remote `json:"remote"`
}

func (o Output) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Output struct{}"
	}

	return strings.Join([]string{"Output", string(data)}, " ")
}
