package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddTrafficEventDto struct {

	//  **参数说明**：事件来源类型列表,支持事件来源。  **取值范围**：  - unknown：未知数据  - police：警方数据  - government：政府数据  - meteorological：气象数据  - internet：互联网数据  - detection：检测器检测到的数据  - v2xServer：平台上报数据  - rsu：RSU上报数据  - obu：车载终端上报数据
	EventSourceType string `json:"event_source_type"`

	// **参数说明**：事件来源的ID，由用户自定义。
	EventSourceId *string `json:"event_source_id,omitempty"`

	// **参数说明**：事件类型，参照附录[《国标交通事件及标志列表》](https://support.huaweicloud.com/api-v2x/v2x_04_0101.html)文件。当填写event_type时，event_class为必选。
	EventClass string `json:"event_class"`

	// **参数说明**：事件类型，参照附录[《国标交通事件及标志列表》](https://support.huaweicloud.com/api-v2x/v2x_04_0101.html)文件。
	EventType int32 `json:"event_type"`

	// **参数说明**：区域码，参考[区域码查询](http://xzqh.mca.gov.cn/map)。
	AreaCode int32 `json:"area_code"`

	// **参数说明**：事件优先级，0-7越大优先级越高。
	EventLevel *int32 `json:"event_level,omitempty"`

	//  **参数说明**：事件附加信息。  事件类型为如下数据时生效：  - 道路最高限：必选，设置最高限速（整数值）km/h  - 道路最低限速：必选，设置最低限速（整数值）km/h  - 建议速度：必选，建议速度（整数值）km/h  - 急弯路：可选，建议最高限速（整数值）km/h  - 雨：可选，请输入1~4：1-细雨，2-小雨，3-中雨，4-大雨  - 雪：可选，请输入1~4：1-小雪，2-中雪，3-大雪，4-暴雪  - 风：可选，设置风速值（整数值）km/h  - 雾：可选，请输入1或2：1-薄雾，2-浓雾  - 路面湿滑：可选，设置湿滑系数（0~1）  - 路面结冰：可选，请设置冰层厚度（整数值）mm  建议填写方式为：user_defined_param1: \"xx\"
	EventParams map[string]string `json:"event_params,omitempty"`

	EventPosition *EventLocation `json:"event_position"`

	// **参数说明**：事件描述。支持英文字母、数字、下划线、斜杠、中文及中文常用字符：。 ？ ！ ， 、 ； ： “ ”
	EventDescription *string `json:"event_description,omitempty"`

	// **参数说明**：事件生效的关联路径，至少需写入起始和终止位置的两个坐标点。
	ReferencePaths []EventReferencePath `json:"reference_paths"`

	// **参数说明**：事件所在位置。
	EventPositionName string `json:"event_position_name"`

	// **参数说明**：开始时间。  格式：yyyy-MM-dd''T''HH:mm:ss''Z''。  例如 2020-09-01T01:37:01Z。
	StartTime string `json:"start_time"`

	// **参数说明**：结束时间。  格式：yyyy-MM-dd''T''HH:mm:ss''Z''。  例如 2020-09-01T01:37:01Z。
	EndTime string `json:"end_time"`

	// **参数说明**：备注。
	Note *string `json:"note,omitempty"`

	// **参数说明**：事件可信度。
	EventConfidence *int32 `json:"event_confidence,omitempty"`
}

func (o AddTrafficEventDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTrafficEventDto struct{}"
	}

	return strings.Join([]string{"AddTrafficEventDto", string(data)}, " ")
}
