package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SmnSubscription 租户smn订阅信息
type SmnSubscription struct {

	// 租户project_id
	Owner string `json:"owner"`

	// 订阅终端
	Endpoint string `json:"endpoint"`

	// 终端协议，比如HTTPS协议，SMS协议，EMAIL协议，HTTP协议
	Protocol SmnSubscriptionProtocol `json:"protocol"`

	// smn订阅的urn
	SubscriptionUrn string `json:"subscription_urn"`

	// 订阅topic对应的urn
	TopicUrn string `json:"topic_urn"`

	// 订阅状态 0：未确认 1：已确认 2：不需要确认 3：已取消确认 4：已经删除。
	Status SmnSubscriptionStatus `json:"status"`
}

func (o SmnSubscription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnSubscription struct{}"
	}

	return strings.Join([]string{"SmnSubscription", string(data)}, " ")
}

type SmnSubscriptionProtocol struct {
	value string
}

type SmnSubscriptionProtocolEnum struct {
	HTTP  SmnSubscriptionProtocol
	HTTPS SmnSubscriptionProtocol
	SMS   SmnSubscriptionProtocol
	EMAIL SmnSubscriptionProtocol
}

func GetSmnSubscriptionProtocolEnum() SmnSubscriptionProtocolEnum {
	return SmnSubscriptionProtocolEnum{
		HTTP: SmnSubscriptionProtocol{
			value: "HTTP",
		},
		HTTPS: SmnSubscriptionProtocol{
			value: "HTTPS",
		},
		SMS: SmnSubscriptionProtocol{
			value: "SMS",
		},
		EMAIL: SmnSubscriptionProtocol{
			value: "EMAIL",
		},
	}
}

func (c SmnSubscriptionProtocol) Value() string {
	return c.value
}

func (c SmnSubscriptionProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmnSubscriptionProtocol) UnmarshalJSON(b []byte) error {
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

type SmnSubscriptionStatus struct {
	value int32
}

type SmnSubscriptionStatusEnum struct {
	E_0 SmnSubscriptionStatus
	E_1 SmnSubscriptionStatus
	E_2 SmnSubscriptionStatus
	E_3 SmnSubscriptionStatus
	E_4 SmnSubscriptionStatus
	E_5 SmnSubscriptionStatus
}

func GetSmnSubscriptionStatusEnum() SmnSubscriptionStatusEnum {
	return SmnSubscriptionStatusEnum{
		E_0: SmnSubscriptionStatus{
			value: 0,
		}, E_1: SmnSubscriptionStatus{
			value: 1,
		}, E_2: SmnSubscriptionStatus{
			value: 2,
		}, E_3: SmnSubscriptionStatus{
			value: 3,
		}, E_4: SmnSubscriptionStatus{
			value: 4,
		}, E_5: SmnSubscriptionStatus{
			value: 5,
		},
	}
}

func (c SmnSubscriptionStatus) Value() int32 {
	return c.value
}

func (c SmnSubscriptionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmnSubscriptionStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
