package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaHostStatistics 配额主机统计信息
type QuotaHostStatistics struct {

	// **参数解释**: 配额版本 **取值范围**: - hss.version.basic：基础版 - hss.version.advanced：高级版 - hss.version.enterprise：企业版 - hss.version.premium：旗舰版 - hss.version.container：容器版 - hss.version.wtp：网页防篡改版 - hss.version.container_wtp：容器网页防篡改版
	Version *string `json:"version,omitempty"`

	// **参数解释**: 该版本配额总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 即将过期的配额数量 **取值范围**: 最小值0，最大值2147483647
	AboutToExpireNum *int32 `json:"about_to_expire_num,omitempty"`

	// **参数解释**: 空闲配额数量 **取值范围**: 最小值0，最大值2147483647
	IdleNum *int32 `json:"idle_num,omitempty"`

	// **参数解释**: 已使用该版本配额的主机数量 **取值范围**: 最小值0，最大值2147483647
	HostNum *int32 `json:"host_num,omitempty"`
}

func (o QuotaHostStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaHostStatistics struct{}"
	}

	return strings.Join([]string{"QuotaHostStatistics", string(data)}, " ")
}
