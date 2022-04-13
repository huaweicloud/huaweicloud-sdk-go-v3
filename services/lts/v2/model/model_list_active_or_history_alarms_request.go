package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListActiveOrHistoryAlarmsRequest struct {
	// domainId

	DomainId string `json:"domain_id"`
	// 是活动告警还是历史告警

	Type ListActiveOrHistoryAlarmsRequestType `json:"type"`
	// 取值为上一页数据的最后一条记录的id(填写上一页数据返回得previous_marker或者next_marker值。)

	Marker *string `json:"marker,omitempty"`
	// 每页数据量

	Limit *int32 `json:"limit,omitempty"`

	Body *ListActiveOrHistoryAlarmsRequestBody `json:"body,omitempty"`
}

func (o ListActiveOrHistoryAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActiveOrHistoryAlarmsRequest struct{}"
	}

	return strings.Join([]string{"ListActiveOrHistoryAlarmsRequest", string(data)}, " ")
}

type ListActiveOrHistoryAlarmsRequestType struct {
	value string
}

type ListActiveOrHistoryAlarmsRequestTypeEnum struct {
	ACTIVE_ALERT  ListActiveOrHistoryAlarmsRequestType
	HISTORY_ALERT ListActiveOrHistoryAlarmsRequestType
}

func GetListActiveOrHistoryAlarmsRequestTypeEnum() ListActiveOrHistoryAlarmsRequestTypeEnum {
	return ListActiveOrHistoryAlarmsRequestTypeEnum{
		ACTIVE_ALERT: ListActiveOrHistoryAlarmsRequestType{
			value: "active_alert",
		},
		HISTORY_ALERT: ListActiveOrHistoryAlarmsRequestType{
			value: "history_alert",
		},
	}
}

func (c ListActiveOrHistoryAlarmsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListActiveOrHistoryAlarmsRequestType) UnmarshalJSON(b []byte) error {
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
