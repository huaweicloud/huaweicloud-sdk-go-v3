package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceMetricsMetadata struct {

	// **参数解释**：资源指标的名称。 **取值范围**：不涉及。
	Name string `json:"name"`

	// **参数解释**：资源指标的标签信息
	Labels *interface{} `json:"labels,omitempty"`
}

func (o ResourceMetricsMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceMetricsMetadata struct{}"
	}

	return strings.Join([]string{"ResourceMetricsMetadata", string(data)}, " ")
}
