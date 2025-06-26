package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IndicatorInfo struct {

	// **参数解释**： 监控指标名称。 **取值范围**： 不涉及。
	IndicatorName *string `json:"indicator_name,omitempty"`

	// **参数解释**： 采集模块名称。 **取值范围**： 不涉及。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 默认采集频率。 **取值范围**： 不涉及。
	DefaultCollectRate *string `json:"default_collect_rate,omitempty"`

	// **参数解释**： 支持的集群版本。 **取值范围**： 不涉及。
	SupportDatastoreVersion *string `json:"support_datastore_version,omitempty"`
}

func (o IndicatorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorInfo struct{}"
	}

	return strings.Join([]string{"IndicatorInfo", string(data)}, " ")
}
