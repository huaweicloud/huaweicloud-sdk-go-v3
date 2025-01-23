package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceDistributionResponse Response Object
type ListInstanceDistributionResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 引擎分布
	EngineDistribution *[]InstanceEngineDistributionListEngineDistribution `json:"engine_distribution,omitempty"`
	HttpStatusCode     int                                                 `json:"-"`
}

func (o ListInstanceDistributionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceDistributionResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceDistributionResponse", string(data)}, " ")
}
