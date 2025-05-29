package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SerialConsoleOptions struct {
	Enabled *bool `json:"enabled,omitempty"`
}

func (o SerialConsoleOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SerialConsoleOptions struct{}"
	}

	return strings.Join([]string{"SerialConsoleOptions", string(data)}, " ")
}
