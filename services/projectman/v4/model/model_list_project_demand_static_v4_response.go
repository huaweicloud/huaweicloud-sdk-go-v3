package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectDemandStaticV4Response Response Object
type ListProjectDemandStaticV4Response struct {

	// 需求统计
	DemandStatistics *[]DemandStatisticResponseV4 `json:"demand_statistics,omitempty"`
	HttpStatusCode   int                          `json:"-"`
}

func (o ListProjectDemandStaticV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectDemandStaticV4Response struct{}"
	}

	return strings.Join([]string{"ListProjectDemandStaticV4Response", string(data)}, " ")
}
