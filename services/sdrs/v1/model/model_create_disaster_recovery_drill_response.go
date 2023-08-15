package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDisasterRecoveryDrillResponse Response Object
type CreateDisasterRecoveryDrillResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDisasterRecoveryDrillResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDisasterRecoveryDrillResponse struct{}"
	}

	return strings.Join([]string{"CreateDisasterRecoveryDrillResponse", string(data)}, " ")
}
