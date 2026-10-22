package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrometheusConfig **参数解释**：Prometheus监控配置。
type PrometheusConfig struct {

	// **参数解释**：是否接入AOM。 **约束限制**：不涉及。 **取值范围**：   - true：接入AOM。   - false：不接入AOM。 **默认取值**：false。
	Enable *bool `json:"enable,omitempty"`

	// **参数解释**：AOM Prometheus实例ID，可通过AOM管理控制台或查询AOM Prometheus实例列表接口获取。 **约束限制**：不涉及。 **取值范围**：长度为0~256个字符，支持英文字符、数字、中划线和下划线的组合。 **默认取值**：不涉及。
	AomPrometheusId *string `json:"aom_prometheus_id,omitempty"`
}

func (o PrometheusConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrometheusConfig struct{}"
	}

	return strings.Join([]string{"PrometheusConfig", string(data)}, " ")
}
