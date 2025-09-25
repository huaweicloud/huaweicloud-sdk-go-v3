package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartFullSqlResponse Response Object
type StartFullSqlResponse struct {

	// **参数解释**: 任务ID。 **约束限制**: 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartFullSqlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartFullSqlResponse struct{}"
	}

	return strings.Join([]string{"StartFullSqlResponse", string(data)}, " ")
}
