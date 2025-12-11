package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSyncRegionsResponse Response Object
type ListSyncRegionsResponse struct {
	Body           *[]RegionInfo `json:"body,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListSyncRegionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSyncRegionsResponse struct{}"
	}

	return strings.Join([]string{"ListSyncRegionsResponse", string(data)}, " ")
}
