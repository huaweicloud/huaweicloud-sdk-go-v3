package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 转发策略对象
type CreateL7policyReq struct {
	// 转发策略名称

	Name *string `json:"name,omitempty"`
	// 转发策略的转发动作；取值：REDIRECT_TO_POOL：转发到后端云服务器组；REDIRECT_TO_LISTENER：重定向到监听器

	Action CreateL7policyReqAction `json:"action"`
	// 转发策略所在的项目ID。

	TenantId *string `json:"tenant_id,omitempty"`
	// 转发策略的管理状态；该字段为预留字段，暂未启用。默认为true。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 转发策略额描述信息

	Description *string `json:"description,omitempty"`
	// 转发策略对应的监听器ID。当action为REDIRECT_TO_POOL时，只支持创建在PROTOCOL为HTTP或TERMINATED_HTTPS的listener上。 当action为REDIRECT_TO_LISTENER时，只支持创建在PROTOCOL为HTTP的listener上。

	ListenerId string `json:"listener_id"`
	// 转发到pool的ID。转发到pool的ID。当action为REDIRECT_TO_POOL时生效。当action为REDIRECT_TO_POOL时必选

	RedirectPoolId *string `json:"redirect_pool_id,omitempty"`
	// 转发到的listener的ID，当action为REDIRECT_TO_LISTENER时生效。当action为REDIRECT_TO_LISTENER时必选

	RedirectListenerId *string `json:"redirect_listener_id,omitempty"`
	// 转发到的url。该字段未启用。

	RedirectUrl *string `json:"redirect_url,omitempty"`
	// 转发策略的优先级，从1递增，最高100。该字段为预留字段，暂未启用。

	Position *int32 `json:"position,omitempty"`
	// 指定L7rule的参数，可以在创建L7policy的同时创建L7rule

	Rules *[]CreateL7ruleReqInPolicy `json:"rules,omitempty"`
}

func (o CreateL7policyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7policyReq struct{}"
	}

	return strings.Join([]string{"CreateL7policyReq", string(data)}, " ")
}

type CreateL7policyReqAction struct {
	value string
}

type CreateL7policyReqActionEnum struct {
	REDIRECT_TO_POOL     CreateL7policyReqAction
	REDIRECT_TO_LISTENER CreateL7policyReqAction
}

func GetCreateL7policyReqActionEnum() CreateL7policyReqActionEnum {
	return CreateL7policyReqActionEnum{
		REDIRECT_TO_POOL: CreateL7policyReqAction{
			value: "REDIRECT_TO_POOL",
		},
		REDIRECT_TO_LISTENER: CreateL7policyReqAction{
			value: "REDIRECT_TO_LISTENER",
		},
	}
}

func (c CreateL7policyReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateL7policyReqAction) UnmarshalJSON(b []byte) error {
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
