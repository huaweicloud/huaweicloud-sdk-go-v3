package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDisasterRecoveryResponse struct {

	// 解除容灾关系的工作ID。
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDisasterRecoveryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDisasterRecoveryResponse struct{}"
	}

	return strings.Join([]string{"DeleteDisasterRecoveryResponse", string(data)}, " ")
}
