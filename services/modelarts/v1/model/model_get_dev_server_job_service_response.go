package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDevServerJobServiceResponse Response Object
type GetDevServerJobServiceResponse struct {

	// **参数解释**：部署服务的id。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：部署服务名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：部署实例状态。 **取值范围**：- CREATING  - RUNNING  - FAILED  -DELETED  - ERROR。
	Status *string `json:"status,omitempty"`

	// **参数解释**：部署服务特性参数。 **取值范围**：不涉及。
	Spec map[string]string `json:"spec,omitempty"`

	// **参数解释**：部署服务实例。
	Instances *[]AiServiceInstance `json:"instances,omitempty"`

	Model          *Model `json:"model,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o GetDevServerJobServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDevServerJobServiceResponse struct{}"
	}

	return strings.Join([]string{"GetDevServerJobServiceResponse", string(data)}, " ")
}
