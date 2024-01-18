package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableUpdatesRequest Request Object
type ShowAvailableUpdatesRequest struct {
}

func (o ShowAvailableUpdatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableUpdatesRequest struct{}"
	}

	return strings.Join([]string{"ShowAvailableUpdatesRequest", string(data)}, " ")
}
