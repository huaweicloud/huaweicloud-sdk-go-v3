package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaStatisticsResponseInfo struct {

	// **参数解释**： 资源规格编码 **取值范围**： 包含如下6种。 - hss.version.basic ：基础版。 - hss.version.advanced ：专业版。 - hss.version.enterprise ：企业版。 - hss.version.premium ：旗舰版。 - hss.version.wtp ：网页防篡改版。 - hss.version.container.enterprise：容器版。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 配额总数 **取值范围**： 0到10000000
	TotalNum *int32 `json:"total_num,omitempty"`
}

func (o QuotaStatisticsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaStatisticsResponseInfo struct{}"
	}

	return strings.Join([]string{"QuotaStatisticsResponseInfo", string(data)}, " ")
}
