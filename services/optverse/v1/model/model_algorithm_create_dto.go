package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmCreateDto 用户创建算法必要的结构体
type AlgorithmCreateDto struct {

	// **参数解释**： 算法名称 **约束限制**： 不涉及 **取值范围**： 长度[0,128] **默认取值**： 不涉及
	Name string `json:"name"`

	// **参数解释**： 算法构建命令 **约束限制**： 不涉及 **取值范围**： 长度[0,256] **默认取值**： 不涉及
	BuildCommand *string `json:"build_command,omitempty"`

	// **参数解释**： 算法预处理命令，bash脚本，python为pip install等预处理过程。 **约束限制**： 不涉及 **取值范围**： 长度[0,256] **默认取值**： 不涉及
	Env *string `json:"env,omitempty"`

	// **参数解释**： 算法默认启动指令 **约束限制**： 不涉及 **取值范围**： 长度[0,256] **默认取值**： 不涉及
	Command *string `json:"command,omitempty"`

	// **参数解释**： 算法描述。 **约束限制**： 不涉及 **取值范围**： 长度[0,32768] **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 算法项目的图标 **约束限制**： 不涉及 **取值范围**： [0,65536] **默认取值**： 无
	Picture *string `json:"picture,omitempty"`

	Lang *ProgramLang `json:"lang"`
}

func (o AlgorithmCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmCreateDto struct{}"
	}

	return strings.Join([]string{"AlgorithmCreateDto", string(data)}, " ")
}
