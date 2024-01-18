package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRobotResponse Response Object
type CreateRobotResponse struct {

	// 应用ID。
	RobotId *string `json:"robot_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRobotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRobotResponse struct{}"
	}

	return strings.Join([]string{"CreateRobotResponse", string(data)}, " ")
}
