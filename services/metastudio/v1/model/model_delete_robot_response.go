package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRobotResponse Response Object
type DeleteRobotResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRobotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRobotResponse struct{}"
	}

	return strings.Join([]string{"DeleteRobotResponse", string(data)}, " ")
}
