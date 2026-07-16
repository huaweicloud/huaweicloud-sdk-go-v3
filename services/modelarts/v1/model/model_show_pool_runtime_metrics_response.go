package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPoolRuntimeMetricsResponse Response Object
type ShowPoolRuntimeMetricsResponse struct {

	// **参数解释**：资源版本。 **取值范围**：可选值如下： - os.modelarts.xxxxx/v2
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**：资源类型。 **取值范围**：可选值如下： - PoolMetricsList
	Kind *string `json:"kind,omitempty"`

	// **参数解释**：指标列表。
	Items          *[]MetricsItem `json:"items,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowPoolRuntimeMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolRuntimeMetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowPoolRuntimeMetricsResponse", string(data)}, " ")
}
