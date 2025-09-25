package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSqlLimitTaskResponse Response Object
type UpdateSqlLimitTaskResponse struct {

	// **参数解释**: 工作流ID。 **取值范围**: 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSqlLimitTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlLimitTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateSqlLimitTaskResponse", string(data)}, " ")
}
