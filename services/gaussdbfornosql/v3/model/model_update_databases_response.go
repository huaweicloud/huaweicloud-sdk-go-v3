package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDatabasesResponse Response Object
type UpdateDatabasesResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabasesResponse struct{}"
	}

	return strings.Join([]string{"UpdateDatabasesResponse", string(data)}, " ")
}
