package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopInstanceRequest Request Object
type StopInstanceRequest struct {
	Body *CommonCbhRequestBody `json:"body,omitempty"`
}

func (o StopInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopInstanceRequest struct{}"
	}

	return strings.Join([]string{"StopInstanceRequest", string(data)}, " ")
}
