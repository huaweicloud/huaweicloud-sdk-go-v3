package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDdmInstanceNameResponse Response Object
type UpdateDdmInstanceNameResponse struct {

	// **参数解释**：  工作流ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDdmInstanceNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDdmInstanceNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateDdmInstanceNameResponse", string(data)}, " ")
}
