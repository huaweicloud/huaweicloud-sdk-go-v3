package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSyncRegionsRequest Request Object
type ListSyncRegionsRequest struct {
}

func (o ListSyncRegionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSyncRegionsRequest struct{}"
	}

	return strings.Join([]string{"ListSyncRegionsRequest", string(data)}, " ")
}
