package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSubscriptionListsRequest Request Object
type ShowSubscriptionListsRequest struct {

	// 请求语言类型。
	XLanguage *ShowSubscriptionListsRequestXLanguage `json:"X-Language,omitempty"`

	// 查询返回记录的数量限制，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	Body *QuerySubscriptionsReq `json:"body,omitempty"`
}

func (o ShowSubscriptionListsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubscriptionListsRequest struct{}"
	}

	return strings.Join([]string{"ShowSubscriptionListsRequest", string(data)}, " ")
}

type ShowSubscriptionListsRequestXLanguage struct {
	value string
}

type ShowSubscriptionListsRequestXLanguageEnum struct {
	EN_US ShowSubscriptionListsRequestXLanguage
	ZH_CN ShowSubscriptionListsRequestXLanguage
}

func GetShowSubscriptionListsRequestXLanguageEnum() ShowSubscriptionListsRequestXLanguageEnum {
	return ShowSubscriptionListsRequestXLanguageEnum{
		EN_US: ShowSubscriptionListsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowSubscriptionListsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowSubscriptionListsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSubscriptionListsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSubscriptionListsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
