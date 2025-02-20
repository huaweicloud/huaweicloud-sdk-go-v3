package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSubscriptionUserRequest Request Object
type ListSubscriptionUserRequest struct {

	// 订阅用户名称。
	Name *string `json:"name,omitempty"`

	// 协议。 http：HTTP终端 https：HTTPS终端 sms：短信 email：邮件
	Protocol *ListSubscriptionUserRequestProtocol `json:"protocol,omitempty"`

	// 订阅用户状态。 UNCONFIRMED：未确认 CONFIRMED：已确认 CANCELLED：已取消
	Status *ListSubscriptionUserRequestStatus `json:"status,omitempty"`

	// 订阅用户分组。
	Group *string `json:"group,omitempty"`

	// 偏移量。偏移量为一个大于0小于资源总个数的整数，表示查询该偏移量后面的所有的资源，默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量限制。取值范围：1~100，取值一般为10，20，50。功能说明：每页返回的资源个数。默认值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSubscriptionUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionUserRequest struct{}"
	}

	return strings.Join([]string{"ListSubscriptionUserRequest", string(data)}, " ")
}

type ListSubscriptionUserRequestProtocol struct {
	value string
}

type ListSubscriptionUserRequestProtocolEnum struct {
	HTTP  ListSubscriptionUserRequestProtocol
	HTTPS ListSubscriptionUserRequestProtocol
	SMS   ListSubscriptionUserRequestProtocol
	EMAIL ListSubscriptionUserRequestProtocol
}

func GetListSubscriptionUserRequestProtocolEnum() ListSubscriptionUserRequestProtocolEnum {
	return ListSubscriptionUserRequestProtocolEnum{
		HTTP: ListSubscriptionUserRequestProtocol{
			value: "http",
		},
		HTTPS: ListSubscriptionUserRequestProtocol{
			value: "https",
		},
		SMS: ListSubscriptionUserRequestProtocol{
			value: "sms",
		},
		EMAIL: ListSubscriptionUserRequestProtocol{
			value: "email",
		},
	}
}

func (c ListSubscriptionUserRequestProtocol) Value() string {
	return c.value
}

func (c ListSubscriptionUserRequestProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSubscriptionUserRequestProtocol) UnmarshalJSON(b []byte) error {
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

type ListSubscriptionUserRequestStatus struct {
	value string
}

type ListSubscriptionUserRequestStatusEnum struct {
	UNCONFIRMED ListSubscriptionUserRequestStatus
	CONFIRMED   ListSubscriptionUserRequestStatus
	CANCELLED   ListSubscriptionUserRequestStatus
}

func GetListSubscriptionUserRequestStatusEnum() ListSubscriptionUserRequestStatusEnum {
	return ListSubscriptionUserRequestStatusEnum{
		UNCONFIRMED: ListSubscriptionUserRequestStatus{
			value: "UNCONFIRMED",
		},
		CONFIRMED: ListSubscriptionUserRequestStatus{
			value: "CONFIRMED",
		},
		CANCELLED: ListSubscriptionUserRequestStatus{
			value: "CANCELLED",
		},
	}
}

func (c ListSubscriptionUserRequestStatus) Value() string {
	return c.value
}

func (c ListSubscriptionUserRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSubscriptionUserRequestStatus) UnmarshalJSON(b []byte) error {
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
