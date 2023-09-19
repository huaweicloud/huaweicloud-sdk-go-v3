package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogConvergeConfigResponse Response Object
type UpdateLogConvergeConfigResponse struct {
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLogConvergeConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogConvergeConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateLogConvergeConfigResponse", string(data)}, " ")
}
