package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListEventsResponsePageInfo struct {
	NextMarker *string `json:"next_marker,omitempty"`
}

func (o ListEventsResponsePageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsResponsePageInfo struct{}"
	}

	return strings.Join([]string{"ListEventsResponsePageInfo", string(data)}, " ")
}
