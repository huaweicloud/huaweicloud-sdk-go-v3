package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceQuotasInfo struct {

	// **参数解释**： 主机开通的版本 **取值范围**： 包含如下7种输入。   - hss.version.null ：无。   - hss.version.basic ：基础版。   - hss.version.advanced ：专业版。   - hss.version.enterprise ：企业版。   - hss.version.premium ：旗舰版。   - hss.version.wtp ：网页防篡改版。   - hss.version.container.enterprise：容器版。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 总配额数 **取值范围**： 0-2000000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： 已使用配额数 **取值范围**： 0-2000000
	UsedNum *int32 `json:"used_num,omitempty"`

	// **参数解释**： 可用总配额数 **取值范围**： 0-2000000
	AvailableNum *int32 `json:"available_num,omitempty"`

	// **参数解释**： 可用资源列表 **取值范围**： 不涉及
	AvailableResourcesList *[]AvailableResourceIdsInfo `json:"available_resources_list,omitempty"`
}

func (o ResourceQuotasInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceQuotasInfo struct{}"
	}

	return strings.Join([]string{"ResourceQuotasInfo", string(data)}, " ")
}
