package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceDistributionResponse Response Object
type ShowInstanceDistributionResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 引擎分布
	EngineDistribution *[]EngineDistributionInfo `json:"engine_distribution,omitempty"`
	HttpStatusCode     int                       `json:"-"`
}

func (o ShowInstanceDistributionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceDistributionResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceDistributionResponse", string(data)}, " ")
}
