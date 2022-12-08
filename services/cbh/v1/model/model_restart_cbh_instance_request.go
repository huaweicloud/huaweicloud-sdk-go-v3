package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RestartCbhInstanceRequest struct {
	Body *RebootCbhRequestBody `json:"body,omitempty"`
}

func (o RestartCbhInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartCbhInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestartCbhInstanceRequest", string(data)}, " ")
}
