package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopJobRequest Request Object
type StopJobRequest struct {
	Body *StopJobRequestBody `json:"body,omitempty"`
}

func (o StopJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobRequest struct{}"
	}

	return strings.Join([]string{"StopJobRequest", string(data)}, " ")
}
