package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserEventsLtsConfigurationsRequest Request Object
type ListUserEventsLtsConfigurationsRequest struct {
}

func (o ListUserEventsLtsConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserEventsLtsConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListUserEventsLtsConfigurationsRequest", string(data)}, " ")
}
