package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 转发策略对象，用于状态树
type L7policiesInStatusResp struct {
	// 转发策略ID

	Id string `json:"id"`
	// 转发策略名称

	Name string `json:"name"`
	// 转发策略关联的转发规则列表

	Rules []L7rulesInStatusResp `json:"rules"`
	// 转发策略的转发动作；取值：REDIRECT_TO_POOL：转发到后端云服务器组；REDIRECT_TO_LISTENER：重定向到监听器

	Action L7policiesInStatusRespAction `json:"action"`
	// 健康检查的配置状态；该字段为预留字段，暂未启用。默认为ACTIVE。

	ProvisioningStatus string `json:"provisioning_status"`
}

func (o L7policiesInStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L7policiesInStatusResp struct{}"
	}

	return strings.Join([]string{"L7policiesInStatusResp", string(data)}, " ")
}

type L7policiesInStatusRespAction struct {
	value string
}

type L7policiesInStatusRespActionEnum struct {
	REDIRECT_TO_POOL     L7policiesInStatusRespAction
	REDIRECT_TO_LISTENER L7policiesInStatusRespAction
}

func GetL7policiesInStatusRespActionEnum() L7policiesInStatusRespActionEnum {
	return L7policiesInStatusRespActionEnum{
		REDIRECT_TO_POOL: L7policiesInStatusRespAction{
			value: "REDIRECT_TO_POOL",
		},
		REDIRECT_TO_LISTENER: L7policiesInStatusRespAction{
			value: "REDIRECT_TO_LISTENER",
		},
	}
}

func (c L7policiesInStatusRespAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *L7policiesInStatusRespAction) UnmarshalJSON(b []byte) error {
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
