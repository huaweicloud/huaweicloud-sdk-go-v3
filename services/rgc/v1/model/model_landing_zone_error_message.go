package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LandingZoneErrorMessage Landing Zone错误信息。
type LandingZoneErrorMessage struct {

	// Landing Zone的错误级别。
	Level *string `json:"level,omitempty"`

	// Landing Zone详细的错误信息。
	Message *string `json:"message,omitempty"`
}

func (o LandingZoneErrorMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LandingZoneErrorMessage struct{}"
	}

	return strings.Join([]string{"LandingZoneErrorMessage", string(data)}, " ")
}
