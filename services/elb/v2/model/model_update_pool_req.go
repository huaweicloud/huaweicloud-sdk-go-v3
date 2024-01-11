package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePoolReq 更新后端云服务器组的请求体
type UpdatePoolReq struct {

	// 后端云服务器组的负载均衡算法，取值：ROUND_ROBIN：加权轮询算法；LEAST_CONNECTIONS：加权最少连接算法；SOURCE_IP：源IP算法；当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。
	LbAlgorithm *string `json:"lb_algorithm,omitempty"`

	// 后端云服务器组的名称。
	Name *string `json:"name,omitempty"`

	// 后端云服务器组的描述信息
	Description *string `json:"description,omitempty"`

	// 后端云服务器组的管理状态；该字段为预留字段，暂未启用。只支持更新为true。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	SessionPersistence *SessionPersistence `json:"session_persistence,omitempty"`

	// 修改保护状态, 取值： - nonProtection: 不保护 - consoleProtection: 控制台修改保护
	ProtectionStatus *UpdatePoolReqProtectionStatus `json:"protection_status,omitempty"`

	// 设置保护的原因 >仅当protection_status为consoleProtection时有效。
	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o UpdatePoolReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolReq struct{}"
	}

	return strings.Join([]string{"UpdatePoolReq", string(data)}, " ")
}

type UpdatePoolReqProtectionStatus struct {
	value string
}

type UpdatePoolReqProtectionStatusEnum struct {
	NON_PROTECTION     UpdatePoolReqProtectionStatus
	CONSOLE_PROTECTION UpdatePoolReqProtectionStatus
}

func GetUpdatePoolReqProtectionStatusEnum() UpdatePoolReqProtectionStatusEnum {
	return UpdatePoolReqProtectionStatusEnum{
		NON_PROTECTION: UpdatePoolReqProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: UpdatePoolReqProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c UpdatePoolReqProtectionStatus) Value() string {
	return c.value
}

func (c UpdatePoolReqProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePoolReqProtectionStatus) UnmarshalJSON(b []byte) error {
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
