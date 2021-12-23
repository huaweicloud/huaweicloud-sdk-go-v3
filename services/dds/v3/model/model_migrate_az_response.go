package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type MigrateAzResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o MigrateAzResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateAzResponse struct{}"
	}

	return strings.Join([]string{"MigrateAzResponse", string(data)}, " ")
}
