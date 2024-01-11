package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateListenerReq 更新监听器的请求体。
type UpdateListenerReq struct {

	// 监听器名称。
	Name *string `json:"name,omitempty"`

	// 监听器的描述信息
	Description *string `json:"description,omitempty"`

	// 监听器的最大连接数。该字段为预留字段，暂未启用。默认为-1。
	ConnectionLimit *int32 `json:"connection_limit,omitempty"`

	// HTTP2功能的开启状态。该字段只有当监听器的协议是TERMINATED_HTTPS时才有意义。
	Http2Enable *bool `json:"http2_enable,omitempty"`

	// 监听器的默认后端云服务器组ID。当请求没有匹配的转发策略时，转发到默认后端云服务器上处理。当该字段为null时，表示监听器无默认的后端云服务器组。
	DefaultPoolId *string `json:"default_pool_id,omitempty"`

	// 监听器使用的服务器证书ID。当protocol参数为TERMINATED_HTTPS时，为必选字段
	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`

	// 监听器使用的CA证书ID。
	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`

	// 监听器使用的SNI证书（带域名的服务器证书）ID的列表。
	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`

	InsertHeaders *InsertHeader `json:"insert_headers,omitempty"`

	// 监听器使用的安全策略，仅对TERMINATED_HTTPS协议类型的监听器有效。  取值包括：tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict多种安全策略。  加密套件的排序为国密套件、ecc套件、rsa套件、tls1.3协议的套件（即支持ecc又支持rsa）
	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty"`

	// 监听器的管理状态。  该字段为预留字段，暂未启动。只支持设定为true
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 修改保护状态, 取值： - nonProtection: 不保护 - consoleProtection: 控制台修改保护
	ProtectionStatus *UpdateListenerReqProtectionStatus `json:"protection_status,omitempty"`

	// 设置保护的原因 >仅当protection_status为consoleProtection时有效。
	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o UpdateListenerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerReq struct{}"
	}

	return strings.Join([]string{"UpdateListenerReq", string(data)}, " ")
}

type UpdateListenerReqProtectionStatus struct {
	value string
}

type UpdateListenerReqProtectionStatusEnum struct {
	NON_PROTECTION     UpdateListenerReqProtectionStatus
	CONSOLE_PROTECTION UpdateListenerReqProtectionStatus
}

func GetUpdateListenerReqProtectionStatusEnum() UpdateListenerReqProtectionStatusEnum {
	return UpdateListenerReqProtectionStatusEnum{
		NON_PROTECTION: UpdateListenerReqProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: UpdateListenerReqProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c UpdateListenerReqProtectionStatus) Value() string {
	return c.value
}

func (c UpdateListenerReqProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateListenerReqProtectionStatus) UnmarshalJSON(b []byte) error {
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
