package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScrumJobCacheRequest Request Object
type ListScrumJobCacheRequest struct {
	Body *ListCacheDatasRequest `json:"body,omitempty"`
}

func (o ListScrumJobCacheRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScrumJobCacheRequest struct{}"
	}

	return strings.Join([]string{"ListScrumJobCacheRequest", string(data)}, " ")
}
