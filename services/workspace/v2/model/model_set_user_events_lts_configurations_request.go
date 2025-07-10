package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetUserEventsLtsConfigurationsRequest Request Object
type SetUserEventsLtsConfigurationsRequest struct {
	Body *SetUserEventsLtsConfigurationsRequestBody `json:"body,omitempty"`
}

func (o SetUserEventsLtsConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetUserEventsLtsConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"SetUserEventsLtsConfigurationsRequest", string(data)}, " ")
}
