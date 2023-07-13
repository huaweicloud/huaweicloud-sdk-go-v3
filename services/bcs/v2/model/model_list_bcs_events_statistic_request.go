package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBcsEventsStatisticRequest Request Object
type ListBcsEventsStatisticRequest struct {

	// 区块链服务id
	BlockchainId string `json:"blockchain_id"`

	// 查询类型。type=active_alert代表查询活动告警，type=history_alert代表查询历史告警。不传或者传其他值则返回指定查询条件的所有信息
	Type *ListBcsEventsStatisticRequestType `json:"type,omitempty"`

	Body *ListBcsEventRequestBody `json:"body,omitempty"`
}

func (o ListBcsEventsStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBcsEventsStatisticRequest struct{}"
	}

	return strings.Join([]string{"ListBcsEventsStatisticRequest", string(data)}, " ")
}

type ListBcsEventsStatisticRequestType struct {
	value string
}

type ListBcsEventsStatisticRequestTypeEnum struct {
	HISTORY_ALERT ListBcsEventsStatisticRequestType
	ACTIVE_ALERT  ListBcsEventsStatisticRequestType
}

func GetListBcsEventsStatisticRequestTypeEnum() ListBcsEventsStatisticRequestTypeEnum {
	return ListBcsEventsStatisticRequestTypeEnum{
		HISTORY_ALERT: ListBcsEventsStatisticRequestType{
			value: "history_alert",
		},
		ACTIVE_ALERT: ListBcsEventsStatisticRequestType{
			value: "active_alert",
		},
	}
}

func (c ListBcsEventsStatisticRequestType) Value() string {
	return c.value
}

func (c ListBcsEventsStatisticRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBcsEventsStatisticRequestType) UnmarshalJSON(b []byte) error {
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
