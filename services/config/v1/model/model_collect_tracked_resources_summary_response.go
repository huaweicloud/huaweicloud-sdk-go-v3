package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectTrackedResourcesSummaryResponse Response Object
type CollectTrackedResourcesSummaryResponse struct {

	// 资源概要信息列表
	Body           *[]ResourceSummaryResponseItem `json:"body,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o CollectTrackedResourcesSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectTrackedResourcesSummaryResponse struct{}"
	}

	return strings.Join([]string{"CollectTrackedResourcesSummaryResponse", string(data)}, " ")
}
