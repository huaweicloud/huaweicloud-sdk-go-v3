package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDisasterRecoveryOperationResponse Response Object
type CheckDisasterRecoveryOperationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckDisasterRecoveryOperationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDisasterRecoveryOperationResponse struct{}"
	}

	return strings.Join([]string{"CheckDisasterRecoveryOperationResponse", string(data)}, " ")
}
