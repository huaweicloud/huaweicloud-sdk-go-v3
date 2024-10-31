package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartMysqlCompatibilityResponse Response Object
type StartMysqlCompatibilityResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartMysqlCompatibilityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartMysqlCompatibilityResponse struct{}"
	}

	return strings.Join([]string{"StartMysqlCompatibilityResponse", string(data)}, " ")
}
