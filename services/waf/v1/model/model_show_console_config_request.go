package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsoleConfigRequest Request Object
type ShowConsoleConfigRequest struct {
}

func (o ShowConsoleConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsoleConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowConsoleConfigRequest", string(data)}, " ")
}
