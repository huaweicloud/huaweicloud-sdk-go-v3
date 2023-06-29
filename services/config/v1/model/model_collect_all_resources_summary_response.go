package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectAllResourcesSummaryResponse Response Object
type CollectAllResourcesSummaryResponse struct {

	// 资源概要信息列表
	Body           *[]ResourceSummaryResponseItem `json:"body,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o CollectAllResourcesSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectAllResourcesSummaryResponse struct{}"
	}

	return strings.Join([]string{"CollectAllResourcesSummaryResponse", string(data)}, " ")
}
