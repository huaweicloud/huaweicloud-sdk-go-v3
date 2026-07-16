package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InputResp struct {

	// **参数解释**：数据输入通道名称。 **取值范围**：不涉及。
	Name string `json:"name"`

	// **参数解释**：数据输入通道描述信息。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：数据输入通道映射的容器本地路径。例如，“/home/ma-user/modelarts/inputs/data_url_0”。 **取值范围**：不涉及。
	LocalDir *string `json:"local_dir,omitempty"`

	// **参数解释**：数据输入通道路径（local_dir）的下发方式。 **取值范围**： - parameter：超参形式 - env：环境变量形式
	AccessMethod *string `json:"access_method,omitempty"`

	Remote *InputDataInfoResp `json:"remote"`

	// **参数解释**：数据输入约束。
	RemoteConstraint *[]InputRespRemoteConstraint `json:"remote_constraint,omitempty"`
}

func (o InputResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputResp struct{}"
	}

	return strings.Join([]string{"InputResp", string(data)}, " ")
}
