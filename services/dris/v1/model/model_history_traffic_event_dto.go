package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HistoryTrafficEventDto struct {

	// **参数说明**：事件ID，由平台生成。
	EventId *string `json:"event_id,omitempty"`

	//  **参数说明**：事件来源类型列表,支持事件来源。  **取值范围**：  - unknown：未知数据  - police：警方数据  - government：政府数据  - meteorological：气象数据  - internet：互联网数据  - detection：检测器检测到的数据  - v2xServer：平台上报数据  - rsu：RSU上报数据  - obu：车载终端上报数据
	EventSourceType *string `json:"event_source_type,omitempty"`

	// **参数说明**：事件来源ID。
	EventSourceId *string `json:"event_source_id,omitempty"`

	// **参数说明**：区域码。
	AreaCode *int32 `json:"area_code,omitempty"`

	// **参数说明**：事件的分类。  **取值范围**：  - AbnormalTraffic：异常路况  - AbnormalVehicle：异常车况  - AdverseWeather：恶劣天气  - TrafficSign：标志标牌
	EventClass *string `json:"event_class,omitempty"`

	// **参数说明**：事件类型，参照附录[《国标交通事件及标志列表》](https://support.huaweicloud.com/api-v2x/v2x_04_0101.html)文件。
	EventType *int32 `json:"event_type,omitempty"`

	// **参数说明**：路口id，对应到一组雷视拟合设备，检测一个特定的路口或者路段。
	CrossId *string `json:"cross_id,omitempty"`

	// **参数说明**：交通事件描述。
	EventDescription *string `json:"event_description,omitempty"`

	// **参数说明**：事件级别(1-5) 由低到高的事件严重程度。
	EventLevel *int32 `json:"event_level,omitempty"`

	// **参数说明**：事件参数，用户自定义。
	EventParams map[string]string `json:"event_params,omitempty"`

	EventPosition *Position3D `json:"event_position,omitempty"`

	// **参数说明**：事件位置名称。
	EventPositionName *string `json:"event_position_name,omitempty"`

	// **参数说明**：事件的关联路径。
	ReferencePaths *[]ReferencePath `json:"reference_paths,omitempty"`

	// **参数说明**：用户备注信息。
	Note *string `json:"note,omitempty"`

	// **参数说明**：事件的状态。  **取值范围**：  - Invalid：过期事件，事件发生的时间段在当前时间之前。  - Active：活动事件，事件正在发生，当前时间处于事件发生时间段内。  - Future：未来事件，在当前时间之后才会发生的事件。
	EventStatus *string `json:"event_status,omitempty"`

	// **参数说明**：事件可信度。
	EventConfidence *int32 `json:"event_confidence,omitempty"`

	EventExInfo *EventExInfo `json:"event_ex_info,omitempty"`

	// **参数说明**：事件关联的rsuID。
	RsuId *[]string `json:"rsu_id,omitempty"`

	// **参数说明**：事件的开始时间。  格式：yyyy-MM-dd''T''HH:mm:ss''Z''。  例如 2020-09-01T01:37:01Z。
	StartTime *string `json:"start_time,omitempty"`

	// **参数说明**：事件的结束时间。  格式：yyyy-MM-dd''T''HH:mm:ss''Z''。  例如 2020-09-01T01:37:01Z。
	EndTime *string `json:"end_time,omitempty"`

	// **参数说明**：事件的创建时间。  格式：yyyy-MM-dd''T''HH:mm:ss''Z''。  例如 2020-09-01T01:37:01Z。
	CreatedTime *string `json:"created_time,omitempty"`
}

func (o HistoryTrafficEventDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryTrafficEventDto struct{}"
	}

	return strings.Join([]string{"HistoryTrafficEventDto", string(data)}, " ")
}
