package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMysqlCompatibilityResponse Response Object
type UpdateMysqlCompatibilityResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMysqlCompatibilityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMysqlCompatibilityResponse struct{}"
	}

	return strings.Join([]string{"UpdateMysqlCompatibilityResponse", string(data)}, " ")
}
