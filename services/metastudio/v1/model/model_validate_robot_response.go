package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateRobotResponse Response Object
type ValidateRobotResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ValidateRobotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateRobotResponse struct{}"
	}

	return strings.Join([]string{"ValidateRobotResponse", string(data)}, " ")
}
