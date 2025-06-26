package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LtslogInfo **参数解释**： LTS日志信息。 **取值范围**： 不涉及。
type LtslogInfo struct {

	// **参数解释**： 配置状态。 **取值范围**： 不涉及。
	Status string `json:"status"`

	// **参数解释**： 日志ID。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 日志类型。 **取值范围**： 不涉及。
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
