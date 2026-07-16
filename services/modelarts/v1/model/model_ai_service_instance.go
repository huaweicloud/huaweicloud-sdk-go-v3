package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiServiceInstance Lite Server部署服务的实例信息
type AiServiceInstance struct {

	// **参数解释**：部署实例id。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：部署实例状态。 **取值范围**：- CREATING - RUNNING  - FAILED  -DELETED - ERROR
	Status *string `json:"status,omitempty"`

	// **参数解释**：调用方式信息。 **取值范围**：不涉及。
	Endpoints *string `json:"endpoints,omitempty"`
}

func (o AiServiceInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiServiceInstance struct{}"
	}

	return strings.Join([]string{"AiServiceInstance", string(data)}, " ")
}
