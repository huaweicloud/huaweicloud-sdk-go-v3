package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchoverTestRequest Request Object
type SwitchoverTestRequest struct {
	Body *CreateSwitchoverTestRequestBody `json:"body,omitempty"`
}

func (o SwitchoverTestRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchoverTestRequest struct{}"
	}

	return strings.Join([]string{"SwitchoverTestRequest", string(data)}, " ")
}
