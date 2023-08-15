package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHistoryTrafficEventsRequest Request Object
type ShowHistoryTrafficEventsRequest struct {

	// \"**参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和横杠（-）的组合，长度36。\"
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：分页查询时的页码。
	Offset *int32 `json:"offset,omitempty"`

	// **参数说明**：分页查询时每页显示的记录数。
	Limit *int32 `json:"limit,omitempty"`

	// **参数说明**：查询此时间后达到平台的消息，十位Unix时间戳，秒级，例如：1575446506。
	FromDate *int64 `json:"from_date,omitempty"`

	// **参数说明**：查询此时间前达到平台的消息，十位Unix时间戳，秒级，例如：1576396905。
	ToDate *int64 `json:"to_date,omitempty"`

	// **参数说明**：事件id
	EventId *string `json:"event_id,omitempty"`

	// **参数说明**：事件的分类，event_type存在时，event_class必选。  **取值范围**：  - AbnormalTraffic：异常路况  - AbnormalVehicle：异常车况  - AdverseWeather：恶劣天气  - TrafficSign：标志标牌
	EventClass *string `json:"event_class,omitempty"`

	// **参数说明**：事件类型，参照附录[《国标交通事件及标志列表》](https://support.huaweicloud.com/api-v2x/v2x_04_0101.html)文件。
	EventType *string `json:"event_type,omitempty"`

	// **参数说明**：事件来源，默认值v2xServer。  **取值范围**：  - v2xServer：来源为通过Console下发的事件  - detection：来源为边缘edge上报的RSI  - rsu：来源为通过MQTT接入的RSU上报的RSI
	EventSource *string `json:"event_source,omitempty"`
}

func (o ShowHistoryTrafficEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoryTrafficEventsRequest struct{}"
	}

	return strings.Join([]string{"ShowHistoryTrafficEventsRequest", string(data)}, " ")
}
