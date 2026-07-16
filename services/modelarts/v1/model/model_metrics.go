package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Metrics 指标采集配置。
type Metrics struct {

	// **参数解释：** 指标采集地址，支持IP地址、域名或localhost。 **取值范围：** 不涉及。
	Endpoint string `json:"endpoint"`

	// **参数解释：** 指标采集路径。 **取值范围：** 不涉及
	Path *string `json:"path,omitempty"`

	// **参数解释：** 指标采集端口。 **取值范围：** 1~65535。
	Port string `json:"port"`

	// **参数解释：** 指标采集协议。 **取值范围：** - HTTP。 - HTTPS。
	Scheme string `json:"scheme"`

	// **参数解释：** 指标来源类型。 **取值范围：** - CONTAINER表示容器内。 - OTHERS表示外部其他地址。 **约束限制：** 不涉及。 **默认取值：** CONTAINER。
	MetricsSource *string `json:"metrics_source,omitempty"`
}

func (o Metrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metrics struct{}"
	}

	return strings.Join([]string{"Metrics", string(data)}, " ")
}
