package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDisasterRecoveryResponse struct {

	// 搭建容灾关系的工作ID。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 容灾ID。
	DisasterRecoveryId *string `json:"disaster_recovery_id,omitempty" xml:"disaster_recovery_id"`
	HttpStatusCode     int     `json:"-"`
}

func (o CreateDisasterRecoveryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDisasterRecoveryResponse struct{}"
	}

	return strings.Join([]string{"CreateDisasterRecoveryResponse", string(data)}, " ")
}
