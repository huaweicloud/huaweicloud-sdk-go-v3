package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSubscriptionOrderRequest Request Object
type ListSubscriptionOrderRequest struct {

	// 用户当前语言环境
	XLanguage string `json:"X-Language"`

	// smn订阅偏移量
	Offset *int32 `json:"offset,omitempty"`

	// smn订阅返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 订单资源详情信息枚举，DEFAULT:默认缺省值，获取开通的资源列表，不包含套餐包；PURCHASE：在DEFAULT基础上返回租户名下ECS数量；RESOURCE_LIST在DEFAULT基础上返回套餐包列表；USAGE：返回资源用量信息；SMN：返回已订阅的smn topic列表
	Page *ListSubscriptionOrderRequestPage `json:"page,omitempty"`
}

func (o ListSubscriptionOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionOrderRequest struct{}"
	}

	return strings.Join([]string{"ListSubscriptionOrderRequest", string(data)}, " ")
}

type ListSubscriptionOrderRequestPage struct {
	value string
}

type ListSubscriptionOrderRequestPageEnum struct {
	DEFAULT       ListSubscriptionOrderRequestPage
	PURCHASE      ListSubscriptionOrderRequestPage
	SMN           ListSubscriptionOrderRequestPage
	USAGE         ListSubscriptionOrderRequestPage
	RESOURCE_LIST ListSubscriptionOrderRequestPage
}

func GetListSubscriptionOrderRequestPageEnum() ListSubscriptionOrderRequestPageEnum {
	return ListSubscriptionOrderRequestPageEnum{
		DEFAULT: ListSubscriptionOrderRequestPage{
			value: "DEFAULT",
		},
		PURCHASE: ListSubscriptionOrderRequestPage{
			value: "PURCHASE",
		},
		SMN: ListSubscriptionOrderRequestPage{
			value: "SMN",
		},
		USAGE: ListSubscriptionOrderRequestPage{
			value: "USAGE",
		},
		RESOURCE_LIST: ListSubscriptionOrderRequestPage{
			value: "RESOURCE_LIST",
		},
	}
}

func (c ListSubscriptionOrderRequestPage) Value() string {
	return c.value
}

func (c ListSubscriptionOrderRequestPage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSubscriptionOrderRequestPage) UnmarshalJSON(b []byte) error {
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
