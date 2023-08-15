package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBcsEventsRequest Request Object
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
