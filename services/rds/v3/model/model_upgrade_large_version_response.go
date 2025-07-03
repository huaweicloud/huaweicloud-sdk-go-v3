package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeLargeVersionResponse Response Object
type UpgradeLargeVersionResponse struct {

	// 任务流ID
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeLargeVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeLargeVersionResponse struct{}"
	}

	return strings.Join([]string{"UpgradeLargeVersionResponse", string(data)}, " ")
}
