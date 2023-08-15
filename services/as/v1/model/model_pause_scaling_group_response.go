package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PauseScalingGroupResponse Response Object
type PauseScalingGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PauseScalingGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseScalingGroupResponse struct{}"
	}

	return strings.Join([]string{"PauseScalingGroupResponse", string(data)}, " ")
}
