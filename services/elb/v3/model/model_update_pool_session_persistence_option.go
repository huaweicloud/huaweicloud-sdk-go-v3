package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePoolSessionPersistenceOption **参数解释**：会话持久性对象。  **约束限制**：慢启动与会话保持不能同时开启。若都开启则会导致会话保持失效。 [荷兰region不支持该字段，请勿使用。](tag:dt)
type UpdatePoolSessionPersistenceOption struct {

	// **参数解释**：cookie名称。  **约束限制**： - 只有当type为APP_COOKIE时才有效。其他情况下传该字段会报错。 - 网关型LB，不支持该特性，请勿使用。  **取值范围**：最大长度1024个字符。  **默认取值**：不涉及  [不支持该字段，请勿使用。](tag:hws_eu,hcso_dt)
	CookieName *string `json:"cookie_name,omitempty"`

	// **参数解释**：会话保持类型。  [**约束限制**： - 当pool的protocol为TCP、UDP，无论type取值如何，都会被忽略，会话保持只按SOURCE_IP生效。 - 当pool的protocol为HTTP、HTTPS时。type可以为HTTP_COOKIE和APP_COOKIE，其他取值会话保持失效。 - 若pool的protocol为QUIC，则必须开启session_persistence且type为SOURCE_IP。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,srg,fcs,dt)  [**约束限制**： - 当pool的protocol为TCP、UDP，无论type取值如何，都会被忽略，会话保持只按SOURCE_IP生效。 - 当pool的protocol为HTTP、HTTPS时。type只能为HTTP_COOKIE，其他取值会话保持失效。](tag:hws_eu,hcso_dt)  **取值范围**： - SOURCE_IP：基于源IP地址的会话保持，也就是根据客户端的源IP地址将流量持续转发到同一后端服务器（member）。 - HTTP_COOKIE：通过在HTTP响应中插入一个特殊的Cookie，实现会话保持。负载均衡会在后续请求中识别该Cookie，并将请求转发到对应的后端服务器。 - APP_COOKIE：基于应用层Cookie的会话保持。与HTTP_COOKIE不同的是，APP_COOKIE要求后端服务在响应中插入一个指定名称的Cookie，负载均衡器识别后，将后续请求转发到相同后端member。  **默认取值**：不涉及  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt)
	Type *UpdatePoolSessionPersistenceOptionType `json:"type,omitempty"`

	// **参数解释**：会话保持的时间。  **约束限制**：当type为APP_COOKIE时不生效。  **取值范围**： - 如果pool的protocol为TCP、UDP和QUIC，则范围为[1,60]（分钟）。 - 如果pool的protocol为HTTP和HTTPS，则范围为[1,1440]（分钟）。  **默认取值**：不涉及 [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt)
	PersistenceTimeout *int32 `json:"persistence_timeout,omitempty"`
}

func (o UpdatePoolSessionPersistenceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolSessionPersistenceOption struct{}"
	}

	return strings.Join([]string{"UpdatePoolSessionPersistenceOption", string(data)}, " ")
}

type UpdatePoolSessionPersistenceOptionType struct {
	value string
}

type UpdatePoolSessionPersistenceOptionTypeEnum struct {
	SOURCE_IP   UpdatePoolSessionPersistenceOptionType
	HTTP_COOKIE UpdatePoolSessionPersistenceOptionType
	APP_COOKIE  UpdatePoolSessionPersistenceOptionType
}

func GetUpdatePoolSessionPersistenceOptionTypeEnum() UpdatePoolSessionPersistenceOptionTypeEnum {
	return UpdatePoolSessionPersistenceOptionTypeEnum{
		SOURCE_IP: UpdatePoolSessionPersistenceOptionType{
			value: "SOURCE_IP",
		},
		HTTP_COOKIE: UpdatePoolSessionPersistenceOptionType{
			value: "HTTP_COOKIE",
		},
		APP_COOKIE: UpdatePoolSessionPersistenceOptionType{
			value: "APP_COOKIE",
		},
	}
}

func (c UpdatePoolSessionPersistenceOptionType) Value() string {
	return c.value
}

func (c UpdatePoolSessionPersistenceOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePoolSessionPersistenceOptionType) UnmarshalJSON(b []byte) error {
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
