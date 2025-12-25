package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LtslogInfo **参数解释**： LTS日志信息。 **取值范围**： 不涉及。
type LtslogInfo struct {

	// **参数解释**： 配置状态。 **取值范围**： - OPEN：开启。 - CLOSE：关闭。
	Status string `json:"status"`

	// **参数解释**： 日志ID。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 日志类型。 **取值范围**： - messages：系统日志。 - expand：扩容日志。 - roach-controller：roach服务端日志。 - audit：审计日志。 - gtm：gtm日志。 - roach-agent：roach客户端日志。 - cms：cms日志。 - CN：dws-CN节点日志。 - upgrade: 升级日志。 - DN: dws-DN节点日志。
	LogType string `json:"log_type"`

	// **参数解释**： 日志描述。 **取值范围**： 不涉及。
	LogDesc string `json:"log_desc"`

	// **参数解释**： LTS日志访问URL。 **取值范围**： 不涉及。
	AccessUrl string `json:"access_url"`
}

func (o LtslogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtslogInfo struct{}"
	}

	return strings.Join([]string{"LtslogInfo", string(data)}, " ")
}
