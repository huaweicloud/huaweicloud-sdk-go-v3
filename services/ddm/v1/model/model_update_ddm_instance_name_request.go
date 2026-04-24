package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDdmInstanceNameRequest Request Object
type UpdateDdmInstanceNameRequest struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in09，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	Body *UpdateDdmInstanceNameRequestBody `json:"body,omitempty"`
}

func (o UpdateDdmInstanceNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDdmInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateDdmInstanceNameRequest", string(data)}, " ")
}
