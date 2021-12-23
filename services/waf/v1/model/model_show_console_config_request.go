package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowConsoleConfigRequest struct {
}

func (o ShowConsoleConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsoleConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowConsoleConfigRequest", string(data)}, " ")
}
