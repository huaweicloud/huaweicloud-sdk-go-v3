package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonitorSystemResponse Response Object
type ShowMonitorSystemResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *MonitorSystemKindObj `json:"kind,omitempty"`

	Spec           *ShowMonitorSystemResponseBodySpec `json:"spec,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ShowMonitorSystemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitorSystemResponse struct{}"
	}

	return strings.Join([]string{"ShowMonitorSystemResponse", string(data)}, " ")
}
