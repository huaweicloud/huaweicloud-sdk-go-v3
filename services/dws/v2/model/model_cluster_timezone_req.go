package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterTimezoneReq **参数解释**： 修改集群时区请求参数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ClusterTimezoneReq struct {

	// **参数解释**：   时区。示例值：UTC。   **约束限制**：   不涉及。    **取值范围**：  ^(Etc/GMT\\+11|US/Hawaii|Etc/GMT\\+9|US/Alaska|US/Pacific|Etc/GMT\\+8|Canada/Mountain|US/Arizona|Canada/Saskatchewan|Etc/GMT\\+6|US/Central|EST|America/Bogota|Etc/GMT\\+5|Canada/Atlantic|America/Cuiaba|America/Buenos_Aires|Etc/GMT\\+3|Canada/Newfoundland|America/Santiago|Etc/GMT\\+2|Atlantic/Cape_Verde|Europe/London|Africa/Monrovia|UTC|Europe/Belgrade|CET|MET|Europe/Amsterdam|EET|Europe/Athens|Asia/Amman|Asia/Beirut|Europe/Minsk|Africa/Nairobi|Europe/Moscow|Etc/GMT-4|Asia/Tbilisi|Asia/Kabul|Etc/GMT-5|Asia/Calcutta|Etc/GMT-6|Etc/GMT-7|PRC|Asia/Shanghai|Etc/GMT-8|Australia/Perth|Asia/Seoul|Asia/Tokyo|Australia/Darwin|Australia/Adelaide|Australia/Sydney|Australia/Brisbane|Etc/GMT-11|Pacific/Auckland|Etc/GMT-12)$   **默认取值**：    不涉及。
	ClusterTimezone *string `json:"cluster_timezone,omitempty"`
}

func (o ClusterTimezoneReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterTimezoneReq struct{}"
	}

	return strings.Join([]string{"ClusterTimezoneReq", string(data)}, " ")
}
