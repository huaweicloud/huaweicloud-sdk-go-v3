package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowV2EndpointsMetricsRequest Request Object
type ShowV2EndpointsMetricsRequest struct {

	// **参数解释**：**参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。  **取值范围**：长度为32-36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：查询时间步长。 **约束限制**：不涉及。 **取值范围**：单位秒。   - 60：每分钟一个点。   - 3600：每1小时一个点。   - 86400：每天一个点。 **默认取值**：60。
	Step ShowV2EndpointsMetricsRequestStep `json:"step"`

	// **参数解释**：指标名称，可以取值多个，多个值用逗号分隔。 **约束限制**：   - aura_job：[[端点监控指标](https://support.huaweicloud.com/usermanual-aidatalake/aidatalake_01_0075.html)](tag:HEC)[[端点监控指标](https://support.huaweicloud.com/intl/zh-cn/usermanual-aidatalake/aidatalake_01_0075.html)](tag:hws_hk)。   - spark（SPARK_JOB / SPARK_SQL）：[[端点监控指标](https://support.huaweicloud.com/usermanual-aidatalake/aidatalake_01_0067.html)](tag:HEC)[[端点监控指标](https://support.huaweicloud.com/intl/zh-cn/usermanual-aidatalake/aidatalake_01_0067.html)](tag:hws_hk)。 **取值范围**：不涉及。 **默认取值**：不涉及。
	MetricName string `json:"metric_name"`

	// **参数解释**：开始时间。 **约束限制**：不涉及。 **取值范围**：时间戳，单位秒。 **默认取值**：不涉及。
	StartTime int64 `json:"start_time"`

	// **参数解释**：结束时间。 **约束限制**：不涉及。 **取值范围**：时间戳，单位秒。 **默认取值**：不涉及。
	EndTime int64 `json:"end_time"`

	// **参数解释**：端点ID，可通过[查询端点列表](ListV2Endpoints.xml)获取端点ID。 **约束限制**：不涉及。  **取值范围**：长度为32-36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	EndpointId string `json:"endpoint_id"`
}

func (o ShowV2EndpointsMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2EndpointsMetricsRequest struct{}"
	}

	return strings.Join([]string{"ShowV2EndpointsMetricsRequest", string(data)}, " ")
}

type ShowV2EndpointsMetricsRequestStep struct {
	value int32
}

type ShowV2EndpointsMetricsRequestStepEnum struct {
	E_60    ShowV2EndpointsMetricsRequestStep
	E_3600  ShowV2EndpointsMetricsRequestStep
	E_86400 ShowV2EndpointsMetricsRequestStep
}

func GetShowV2EndpointsMetricsRequestStepEnum() ShowV2EndpointsMetricsRequestStepEnum {
	return ShowV2EndpointsMetricsRequestStepEnum{
		E_60: ShowV2EndpointsMetricsRequestStep{
			value: 60,
		}, E_3600: ShowV2EndpointsMetricsRequestStep{
			value: 3600,
		}, E_86400: ShowV2EndpointsMetricsRequestStep{
			value: 86400,
		},
	}
}

func (c ShowV2EndpointsMetricsRequestStep) Value() int32 {
	return c.value
}

func (c ShowV2EndpointsMetricsRequestStep) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowV2EndpointsMetricsRequestStep) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
