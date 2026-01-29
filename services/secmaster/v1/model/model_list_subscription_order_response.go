package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSubscriptionOrderResponse Response Object
type ListSubscriptionOrderResponse struct {

	// 租户当前的版本信息 BASIC(基础版)，STANDARD（标准版），PROFESSIONAL（专业版），NA（无版本），大小写不敏感
	CsbVersion *ListSubscriptionOrderResponseCsbVersion `json:"csb_version,omitempty"`

	// ecs个数，当请求参数purchase=true时才会返回该值，否则为0
	EcsCount *int32 `json:"ecs_count,omitempty"`

	// 资源列表
	Resources *[]SubscriptionResource `json:"resources,omitempty"`

	// topic订阅条数，当请求参数为smn=true，返回该字段
	SubscriptionCount *int32 `json:"subscription_count,omitempty"`

	// 租户订阅信息，当请求参数为smn=true，会返回租户名下可订阅的smn topic列表
	Subscriptions  *[]SmnSubscription `json:"subscriptions,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListSubscriptionOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionOrderResponse struct{}"
	}

	return strings.Join([]string{"ListSubscriptionOrderResponse", string(data)}, " ")
}

type ListSubscriptionOrderResponseCsbVersion struct {
	value string
}

type ListSubscriptionOrderResponseCsbVersionEnum struct {
	BASIC        ListSubscriptionOrderResponseCsbVersion
	STANDARD     ListSubscriptionOrderResponseCsbVersion
	PROFESSIONAL ListSubscriptionOrderResponseCsbVersion
	NA           ListSubscriptionOrderResponseCsbVersion
}

func GetListSubscriptionOrderResponseCsbVersionEnum() ListSubscriptionOrderResponseCsbVersionEnum {
	return ListSubscriptionOrderResponseCsbVersionEnum{
		BASIC: ListSubscriptionOrderResponseCsbVersion{
			value: "BASIC",
		},
		STANDARD: ListSubscriptionOrderResponseCsbVersion{
			value: "STANDARD",
		},
		PROFESSIONAL: ListSubscriptionOrderResponseCsbVersion{
			value: "PROFESSIONAL",
		},
		NA: ListSubscriptionOrderResponseCsbVersion{
			value: "NA",
		},
	}
}

func (c ListSubscriptionOrderResponseCsbVersion) Value() string {
	return c.value
}

func (c ListSubscriptionOrderResponseCsbVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSubscriptionOrderResponseCsbVersion) UnmarshalJSON(b []byte) error {
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
