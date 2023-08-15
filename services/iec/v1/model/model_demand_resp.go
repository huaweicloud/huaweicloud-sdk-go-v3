package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DemandResp 租户需求
type DemandResp struct {

	// 站点需要发放的资源(组)总数。  > 实际发放实例数量为count*demand_count。
	DemandCount int32 `json:"demand_count"`
}

func (o DemandResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DemandResp struct{}"
	}

	return strings.Join([]string{"DemandResp", string(data)}, " ")
}
