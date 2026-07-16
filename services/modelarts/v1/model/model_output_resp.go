package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OutputResp struct {

	// **参数解释**：数据输出通道名称。 **取值范围**：不涉及。
	Name string `json:"name"`

	// **参数解释**：数据输出通道描述信息。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：数据输出通道映射的容器本地路径。 **取值范围**：不涉及。
	LocalDir *string `json:"local_dir,omitempty"`

	// **参数解释**：数据输入通道路径（local_dir）的下发方式。 **取值范围**： - parameter：超参形式 - env：环境变量形式
	AccessMethod *string `json:"access_method,omitempty"`

	Remote *RemoteResp `json:"remote"`
}

func (o OutputResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputResp struct{}"
	}

	return strings.Join([]string{"OutputResp", string(data)}, " ")
}
