package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListBcsEventsRequest struct {

	// 区块链服务id
	BlockchainId string `json:"blockchain_id"`

	// 查询类型。type=active_alert代表查询活动告警，type=history_alert代表查询历史告警。不传或者传其他值则返回指定查询条件的所有信息。
	Type *ListBcsEventsRequestType `json:"type,omitempty"`

	Body *ListBcsEventRequestBody `json:"body,omitempty"`
}

func (o ListBcsEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBcsEventsRequest struct{}"
	}

	return strings.Join([]string{"ListBcsEventsRequest", string(data)}, " ")
}

type ListBcsEventsRequestType struct {
	value string
}

type ListBcsEventsRequestTypeEnum struct {
	HISTORY_ALERT ListBcsEventsRequestType
	ACTIVE_ALERT  ListBcsEventsRequestType
}

func GetListBcsEventsRequestTypeEnum() ListBcsEventsRequestTypeEnum {
	return ListBcsEventsRequestTypeEnum{
		HISTORY_ALERT: ListBcsEventsRequestType{
			value: "history_alert",
		},
		ACTIVE_ALERT: ListBcsEventsRequestType{
			value: "active_alert",
		},
	}
}

func (c ListBcsEventsRequestType) Value() string {
	return c.value
}

func (c ListBcsEventsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBcsEventsRequestType) UnmarshalJSON(b []byte) error {
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
