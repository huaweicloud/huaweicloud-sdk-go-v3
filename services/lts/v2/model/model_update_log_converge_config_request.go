package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogConvergeConfigRequest Request Object
type UpdateLogConvergeConfigRequest struct {
	Body *UpdatelogConvergeConfig `json:"body,omitempty"`
}

func (o UpdateLogConvergeConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogConvergeConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateLogConvergeConfigRequest", string(data)}, " ")
}
