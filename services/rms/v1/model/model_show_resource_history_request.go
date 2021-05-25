package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowResourceHistoryRequest struct {
	// 资源ID

	ResourceId string `json:"resource_id"`
	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页

	Marker *string `json:"marker,omitempty"`
	// 最大的返回数量

	Limit *int32 `json:"limit,omitempty"`
	// 指定查询范围的起始时间点，如果不设置此参数，默认为最早的时间

	EarlierTime *int64 `json:"earlier_time,omitempty"`
	// 指定查询范围的结束时间点，如果不设置此参数，默认为当前时间

	LaterTime *int64 `json:"later_time,omitempty"`
	// 指定返回数据的时间顺序，默认为倒序

	ChronologicalOrder *ShowResourceHistoryRequestChronologicalOrder `json:"chronological_order,omitempty"`
}

func (o ShowResourceHistoryRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowResourceHistoryRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceHistoryRequest", string(data)}, " ")
}

type ShowResourceHistoryRequestChronologicalOrder struct {
	value string
}

type ShowResourceHistoryRequestChronologicalOrderEnum struct {
	FORWARD ShowResourceHistoryRequestChronologicalOrder
	REVERSE ShowResourceHistoryRequestChronologicalOrder
}

func GetShowResourceHistoryRequestChronologicalOrderEnum() ShowResourceHistoryRequestChronologicalOrderEnum {
	return ShowResourceHistoryRequestChronologicalOrderEnum{
		FORWARD: ShowResourceHistoryRequestChronologicalOrder{
			value: "Forward",
		},
		REVERSE: ShowResourceHistoryRequestChronologicalOrder{
			value: "Reverse",
		},
	}
}

func (c ShowResourceHistoryRequestChronologicalOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowResourceHistoryRequestChronologicalOrder) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
