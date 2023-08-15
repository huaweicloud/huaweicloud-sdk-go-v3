package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDisasterRecoveryResponse Response Object
type CreateDisasterRecoveryResponse struct {

	// 搭建容灾关系的工作ID。
	JobId *string `json:"job_id,omitempty"`

	// 容灾ID。
	DisasterRecoveryId *string `json:"disaster_recovery_id,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o CreateDisasterRecoveryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDisasterRecoveryResponse struct{}"
	}

	return strings.Join([]string{"CreateDisasterRecoveryResponse", string(data)}, " ")
}
