package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LtslogInfo **参数解释**： LTS日志信息。 **取值范围**： 不涉及。
type LtslogInfo struct {

	// **参数解释**： 配置状态。 **取值范围**： - OPEN：开启。 - CLOSE：关闭。 - ARCHIVED：历史日志流，不再上报但仍保留LTS访问入口。
	Status string `json:"status"`

	// **参数解释**： 日志ID。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 日志类型。 **取值范围**： - messages：系统日志。 - expand：扩容日志。 - roach-controller：roach服务端日志。 - audit：审计日志。 - gtm：gtm日志。 - roach-agent：roach客户端日志。 - cms：cms日志。 - CN：dws-CN节点日志。 - upgrade: 升级日志。 - DN: dws-DN节点日志。
	LogType string `json:"log_type"`

	// **参数解释**： 日志描述。 **取值范围**： 不涉及。
	LogDesc string `json:"log_desc"`

	// **参数解释**： LTS日志访问URL。 **取值范围**： 不涉及。
	AccessUrl string `json:"access_url"`

	// **参数解释**： 日志上报频率，单位秒。 **约束限制**： 仅 status=OPEN 的日志流返回此字段，ARCHIVED 状态不返回。 **取值范围**： 5~60。
	ReportInterval *int32 `json:"report_interval,omitempty"`

	// **参数解释**： 单次上报大小，单位字节。 **约束限制**： 仅 status=OPEN 的日志流返回此字段，ARCHIVED 状态不返回。 **取值范围**： 51200~1048576。
	MaxReportSize *int32 `json:"max_report_size,omitempty"`
}

func (o LtslogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtslogInfo struct{}"
	}

	return strings.Join([]string{"LtslogInfo", string(data)}, " ")
}
