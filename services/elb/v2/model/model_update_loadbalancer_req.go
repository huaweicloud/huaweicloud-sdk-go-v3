package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateLoadbalancerReq 更新负载均衡器的请求体
type UpdateLoadbalancerReq struct {

	// 负载均衡器名称。
	Name *string `json:"name,omitempty"`

	// 负载均衡器的描述信息
	Description *string `json:"description,omitempty"`

	// 负载均衡器的管理状态。只支持设定为true，该字段的值无实际意义。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 修改保护状态, 取值： - nonProtection: 不保护 - consoleProtection: 控制台修改保护
	ProtectionStatus *UpdateLoadbalancerReqProtectionStatus `json:"protection_status,omitempty"`

	// 设置保护的原因 >仅当protection_status为consoleProtection时有效。
	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o UpdateLoadbalancerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadbalancerReq struct{}"
	}

	return strings.Join([]string{"UpdateLoadbalancerReq", string(data)}, " ")
}

type UpdateLoadbalancerReqProtectionStatus struct {
	value string
}

type UpdateLoadbalancerReqProtectionStatusEnum struct {
	NON_PROTECTION     UpdateLoadbalancerReqProtectionStatus
	CONSOLE_PROTECTION UpdateLoadbalancerReqProtectionStatus
}

func GetUpdateLoadbalancerReqProtectionStatusEnum() UpdateLoadbalancerReqProtectionStatusEnum {
	return UpdateLoadbalancerReqProtectionStatusEnum{
		NON_PROTECTION: UpdateLoadbalancerReqProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: UpdateLoadbalancerReqProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c UpdateLoadbalancerReqProtectionStatus) Value() string {
	return c.value
}

func (c UpdateLoadbalancerReqProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLoadbalancerReqProtectionStatus) UnmarshalJSON(b []byte) error {
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
