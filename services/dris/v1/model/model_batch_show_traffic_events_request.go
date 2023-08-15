package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchShowTrafficEventsRequest Request Object
type BatchShowTrafficEventsRequest struct {

	// \"**参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和横杠（-）的组合，长度36。\"
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：查询事件列表的页码。
	Offset *int32 `json:"offset,omitempty"`

	// **参数说明**：查询时每页显示的记录数。
	Limit *int32 `json:"limit,omitempty"`

	// **参数说明**：区域码，参考[区域码查询](http://xzqh.mca.gov.cn/map)。
	AreaCode *int32 `json:"area_code,omitempty"`

	//  **参数说明**：事件状态。  **取值范围**：  - Invalid：为过期事件，事件结束时间（end_time）在当前时间之前。  - Active：为活动事件，事件开始时间（start_time）在当前时间之前，并且事件结束时间（end_time）在当前时间之后。  - Future：为未来事件，事件开始时间（start_time）在当前时间之前。
	Status *string `json:"status,omitempty"`

	// **参数说明**：事件类型，参照附录[《国标交通事件及标志列表》](https://support.huaweicloud.com/api-v2x/v2x_04_0101.html)文件。当填写event_type时，event_class为必选。
	EventType *int32 `json:"event_type,omitempty"`

	//  **参数说明**：事件来源类型列表,支持事件来源。  **取值范围**：  - unknown：未知数据  - police：警方数据  - government：政府数据  - meteorological：气象数据  - internet：互联网数据  - detection：检测器检测到的数据  - v2xServer：平台上报数据  - rsu：RSU上报数据  - obu：车载终端上报数据
	EventSourceType *[]string `json:"event_source_type,omitempty"`

	// **参数说明**：事件类型，参照附录[《国标交通事件及标志列表》](https://support.huaweicloud.com/api-v2x/v2x_04_0101.html)文件。当填写event_type时，event_class为必选。
	EventClass *string `json:"event_class,omitempty"`

	// **参数说明**：事件ID，创建事件后获得。方法参见 [新增交通事件](https://support.huaweicloud.com/api-v2x/v2x_04_0048.html)。
	EventId *string `json:"event_id,omitempty"`

	// **参数说明**：查询事件开始时间段的起始时间。  格式：yyyy-MM-dd'T'HH:mm:ss'Z'。  例如 2020-09-01T01:37:01Z。
	FromTime *string `json:"from_time,omitempty"`

	// **参数说明**：查询事件开始时间段的结束时间。  格式：yyyy-MM-dd'T'HH:mm:ss'Z'。  例如 2020-09-01T01:37:01Z。
	ToTime *string `json:"to_time,omitempty"`

	// **参数说明**：按照哪一个字段排序,默认按事件开始时间。
	SortKey *BatchShowTrafficEventsRequestSortKey `json:"sort_key,omitempty"`

	// **参数说明**：升序或降序，升序为ASC, 降序为DESC，默认降序。
	SortDir *BatchShowTrafficEventsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o BatchShowTrafficEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowTrafficEventsRequest struct{}"
	}

	return strings.Join([]string{"BatchShowTrafficEventsRequest", string(data)}, " ")
}

type BatchShowTrafficEventsRequestSortKey struct {
	value string
}

type BatchShowTrafficEventsRequestSortKeyEnum struct {
	START_TIME         BatchShowTrafficEventsRequestSortKey
	END_TIME           BatchShowTrafficEventsRequestSortKey
	CREATED_TIME       BatchShowTrafficEventsRequestSortKey
	LAST_MODIFIED_TIME BatchShowTrafficEventsRequestSortKey
}

func GetBatchShowTrafficEventsRequestSortKeyEnum() BatchShowTrafficEventsRequestSortKeyEnum {
	return BatchShowTrafficEventsRequestSortKeyEnum{
		START_TIME: BatchShowTrafficEventsRequestSortKey{
			value: "start_time",
		},
		END_TIME: BatchShowTrafficEventsRequestSortKey{
			value: "end_time",
		},
		CREATED_TIME: BatchShowTrafficEventsRequestSortKey{
			value: "created_time",
		},
		LAST_MODIFIED_TIME: BatchShowTrafficEventsRequestSortKey{
			value: "last_modified_time",
		},
	}
}

func (c BatchShowTrafficEventsRequestSortKey) Value() string {
	return c.value
}

func (c BatchShowTrafficEventsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchShowTrafficEventsRequestSortKey) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type BatchShowTrafficEventsRequestSortDir struct {
	value string
}

type BatchShowTrafficEventsRequestSortDirEnum struct {
	ASC  BatchShowTrafficEventsRequestSortDir
	DESC BatchShowTrafficEventsRequestSortDir
}

func GetBatchShowTrafficEventsRequestSortDirEnum() BatchShowTrafficEventsRequestSortDirEnum {
	return BatchShowTrafficEventsRequestSortDirEnum{
		ASC: BatchShowTrafficEventsRequestSortDir{
			value: "ASC",
		},
		DESC: BatchShowTrafficEventsRequestSortDir{
			value: "DESC",
		},
	}
}

func (c BatchShowTrafficEventsRequestSortDir) Value() string {
	return c.value
}

func (c BatchShowTrafficEventsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchShowTrafficEventsRequestSortDir) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
