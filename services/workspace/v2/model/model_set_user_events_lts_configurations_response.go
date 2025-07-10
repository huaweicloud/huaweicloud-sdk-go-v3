package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetUserEventsLtsConfigurationsResponse Response Object
type SetUserEventsLtsConfigurationsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetUserEventsLtsConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetUserEventsLtsConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"SetUserEventsLtsConfigurationsResponse", string(data)}, " ")
}
