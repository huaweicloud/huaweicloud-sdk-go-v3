package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDisasterRecoveryResponse Response Object
type DeleteDisasterRecoveryResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDisasterRecoveryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDisasterRecoveryResponse struct{}"
	}

	return strings.Join([]string{"DeleteDisasterRecoveryResponse", string(data)}, " ")
}
