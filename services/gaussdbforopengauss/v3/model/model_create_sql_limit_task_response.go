package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSqlLimitTaskResponse Response Object
type CreateSqlLimitTaskResponse struct {

	// **参数解释**: 工作流ID。 **取值范围**: 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSqlLimitTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlLimitTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateSqlLimitTaskResponse", string(data)}, " ")
}
