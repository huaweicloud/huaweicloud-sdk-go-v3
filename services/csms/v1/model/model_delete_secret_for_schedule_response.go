package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecretForScheduleResponse Response Object
type DeleteSecretForScheduleResponse struct {
	Secret         *Secret `json:"secret,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSecretForScheduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecretForScheduleResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecretForScheduleResponse", string(data)}, " ")
}
