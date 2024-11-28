package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyDatabaseResponse Response Object
type CopyDatabaseResponse struct {

	// 任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyDatabaseResponse struct{}"
	}

	return strings.Join([]string{"CopyDatabaseResponse", string(data)}, " ")
}
