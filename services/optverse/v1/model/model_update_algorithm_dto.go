package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAlgorithmDto struct {

	// **参数解释**： 项目ID，您可以从[获取项目ID](ai4sservice_03_0033.xml)中获取。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 算法id **约束限制**： 不涉及 **取值范围**： 长度[0,64] **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 算法名称 **约束限制**： 不涉及 **取值范围**： 长度[0,128] **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 算法构建命令 **约束限制**： 不涉及 **取值范围**： 长度[0,256] **默认取值**： 不涉及
	BuildCommand *string `json:"build_command,omitempty"`

	// **参数解释**： 算法预处理命令，bash脚本，python为pip install等预处理过程。 **约束限制**： 不涉及 **取值范围**： 长度[0,256] **默认取值**： 不涉及
	Env *string `json:"env,omitempty"`

	// **参数解释**： 算法描述。 **约束限制**： 不涉及 **取值范围**： 长度[0,32768] **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 算法默认启动指令 **约束限制**： 不涉及 **取值范围**： 长度[0,256] **默认取值**： 不涉及
	Command *string `json:"command,omitempty"`

	// **参数解释**： 算法的创建时间 **约束限制**： 不涉及 **取值范围**： [0,9999999999999] **默认取值**： 无
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**： 算法项目的图标 **约束限制**： 不涉及 **取值范围**： [0,65536] **默认取值**： 无
	Picture *string `json:"picture,omitempty"`

	Lang *ProgramLang `json:"lang,omitempty"`
}

func (o UpdateAlgorithmDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlgorithmDto struct{}"
	}

	return strings.Join([]string{"UpdateAlgorithmDto", string(data)}, " ")
}
