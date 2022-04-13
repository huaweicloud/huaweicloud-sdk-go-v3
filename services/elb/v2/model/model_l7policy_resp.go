package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 转发策略对象
type L7policyResp struct {
	// 转发策略ID

	Id string `json:"id"`
	// 转发策略名称

	Name string `json:"name"`
	// 转发策略关联的转发规则列表

	Rules []ResourceList `json:"rules"`
	// 转发策略的转发动作；取值：REDIRECT_TO_POOL：转发到后端云服务器组；REDIRECT_TO_LISTENER：重定向到监听器

	Action L7policyRespAction `json:"action"`
	// 健康检查的配置状态；该字段为预留字段，暂未启用。默认为ACTIVE。

	ProvisioningStatus string `json:"provisioning_status"`
	// 转发策略所在的项目ID。

	TenantId string `json:"tenant_id"`
	// 转发策略所在的项目ID。

	ProjectId string `json:"project_id"`
	// 转发策略的管理状态；该字段为预留字段，暂未启用。默认为true。

	AdminStateUp bool `json:"admin_state_up"`
	// 转发策略额描述信息

	Description string `json:"description"`
	// 转发策略对应的监听器ID

	ListenerId string `json:"listener_id"`
	// 转发到pool的ID。转发到pool的ID。当action为REDIRECT_TO_POOL时生效。

	RedirectPoolId string `json:"redirect_pool_id"`
	// 转发到的listener的ID，当action为REDIRECT_TO_LISTENER时生效。

	RedirectListenerId string `json:"redirect_listener_id"`
	// 转发到的url。该字段未启用。

	RedirectUrl string `json:"redirect_url"`
	// 转发策略的优先级，从1递增，最高100。该字段为预留字段，暂未启用。

	Position int32 `json:"position"`
}

func (o L7policyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L7policyResp struct{}"
	}

	return strings.Join([]string{"L7policyResp", string(data)}, " ")
}

type L7policyRespAction struct {
	value string
}

type L7policyRespActionEnum struct {
	REDIRECT_TO_POOL     L7policyRespAction
	REDIRECT_TO_LISTENER L7policyRespAction
}

func GetL7policyRespActionEnum() L7policyRespActionEnum {
	return L7policyRespActionEnum{
		REDIRECT_TO_POOL: L7policyRespAction{
			value: "REDIRECT_TO_POOL",
		},
		REDIRECT_TO_LISTENER: L7policyRespAction{
			value: "REDIRECT_TO_LISTENER",
		},
	}
}

func (c L7policyRespAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *L7policyRespAction) UnmarshalJSON(b []byte) error {
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
