package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotasDetailResponse Response Object
type ListQuotasDetailResponse struct {

	// **参数解释**： 包周期配额数 **取值范围**： 0到10000000
	PacketCycleNum *int32 `json:"packet_cycle_num,omitempty"`

	// **参数解释**： 按需配额数 **取值范围**： 0到10000000
	OnDemandNum *int32 `json:"on_demand_num,omitempty"`

	// **参数解释**： 已使用配额数 **取值范围**： 0到10000000
	UsedNum *int32 `json:"used_num,omitempty"`

	// **参数解释**： 空闲配额数 **取值范围**： 0到10000000
	IdleNum *int32 `json:"idle_num,omitempty"`

	// **参数解释**： 正常配额数 **取值范围**： 0到10000000
	NormalNum *int32 `json:"normal_num,omitempty"`

	// **参数解释**： 过期配额数 **取值范围**： 0到10000000
	ExpiredNum *int32 `json:"expired_num,omitempty"`

	// **参数解释**： 创建时间 **取值范围**： 0到9223372036854775807
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 冻结配额数 **取值范围**： 0到10000000
	FreezeNum *int32 `json:"freeze_num,omitempty"`

	// **参数解释**： 配额统计列表 **取值范围**： 不涉及
	QuotaStatisticsList *[]QuotaStatisticsResponseInfo `json:"quota_statistics_list,omitempty"`

	// **参数解释**： 配额总数 **取值范围**： 0到10000000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： 配额列表 **取值范围**： 不涉及
	DataList       *[]QuotaResourcesResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListQuotasDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotasDetailResponse struct{}"
	}

	return strings.Join([]string{"ListQuotasDetailResponse", string(data)}, " ")
}
