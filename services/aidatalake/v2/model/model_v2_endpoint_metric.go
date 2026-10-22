package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2EndpointMetric **参数解释**：端点指标信息。
type V2EndpointMetric struct {

	// **参数解释**：指标名称。 **取值范围**：监控指标与端点引擎类型相关，具体支持的指标如下：   - aura_job类型端点支持的指标参见：[[端点监控指标](https://support.huaweicloud.com/usermanual-aidatalake/aidatalake_01_0075.html)](tag:HEC)[[端点监控指标](https://support.huaweicloud.com/intl/zh-cn/usermanual-aidatalake/aidatalake_01_0075.html)](tag:hws_hk)。   - spark（SPARK_JOB / SPARK_SQL）类型端点支持的指标参见：[[端点监控指标](https://support.huaweicloud.com/usermanual-aidatalake/aidatalake_01_0067.html)](tag:HEC)[[端点监控指标](https://support.huaweicloud.com/intl/zh-cn/usermanual-aidatalake/aidatalake_01_0067.html)](tag:hws_hk)。
	MetricName *string `json:"metric_name,omitempty"`

	// **参数解释**：指标值。 **取值范围**：不涉及。
	MetricValue float32 `json:"metric_value,omitempty"`
}

func (o V2EndpointMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2EndpointMetric struct{}"
	}

	return strings.Join([]string{"V2EndpointMetric", string(data)}, " ")
}
