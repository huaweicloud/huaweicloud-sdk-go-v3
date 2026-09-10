package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelServiceTaskInputs 模型服务任务输入
type ModelServiceTaskInputs struct {

	// **参数解释**： 创建沙箱任务输入数据文件路径。支持求解助手类型的模型服务任务时使用，优先级高于model_request。 **约束限制**： 不涉及 **取值范围**： 长度为[0-2048]个字符。 **默认取值**： 不涉及
	ModelData *string `json:"model_data,omitempty"`

	// **参数解释**： 创建沙箱任务输入请求体。 **约束限制**： 不涉及 **取值范围**： 长度为[0-12582912]个字符，上限12MB。 **默认取值**： 不涉及
	ModelRequest *string `json:"model_request,omitempty"`
}

func (o ModelServiceTaskInputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelServiceTaskInputs struct{}"
	}

	return strings.Join([]string{"ModelServiceTaskInputs", string(data)}, " ")
}
