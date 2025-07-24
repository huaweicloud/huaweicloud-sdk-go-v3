package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeServerPowerStateRequest Request Object
type ChangeServerPowerStateRequest struct {
	Body *PowerAction `json:"body,omitempty"`
}

func (o ChangeServerPowerStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerPowerStateRequest struct{}"
	}

	return strings.Join([]string{"ChangeServerPowerStateRequest", string(data)}, " ")
}
