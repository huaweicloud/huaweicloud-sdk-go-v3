package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRestoreTablesResponse Response Object
type CreateRestoreTablesResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRestoreTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRestoreTablesResponse struct{}"
	}

	return strings.Join([]string{"CreateRestoreTablesResponse", string(data)}, " ")
}
