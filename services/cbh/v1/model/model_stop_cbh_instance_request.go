package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopCbhInstanceRequest Request Object
type StopCbhInstanceRequest struct {
	Body *StopCbhRequestBody `json:"body,omitempty"`
}

func (o StopCbhInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopCbhInstanceRequest struct{}"
	}

	return strings.Join([]string{"StopCbhInstanceRequest", string(data)}, " ")
}
