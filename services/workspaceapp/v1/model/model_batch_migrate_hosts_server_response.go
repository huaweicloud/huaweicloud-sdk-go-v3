package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchMigrateHostsServerResponse Response Object
type BatchMigrateHostsServerResponse struct {

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchMigrateHostsServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMigrateHostsServerResponse struct{}"
	}

	return strings.Join([]string{"BatchMigrateHostsServerResponse", string(data)}, " ")
}
