package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePoolSessionPersistenceOption 参数解释：会话持久性对象。
type CreatePoolSessionPersistenceOption struct {

	// 参数解释：cookie名称。  约束限制： - 只有当type为APP_COOKIE时才有效。其他情况下传该字段会报错。 [- 网关型LB，不支持该特性，请勿使用。](tag:hws_eu)  [取值范围： - 共享型LB，支持字母、数字、中划线(-)和下划线(_)，最大长度64个字符。 - 独享型LB，支持字母、数字、中划线(-)、下划线(_)和点号(.)，最大长度255个字符。](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt,hk_tm)  [取值范围：支持字母、数字、中划线(-)、下划线(_)和点号(.)，最大长度255个字符。](tag:hws_eu,hcso_dt)  [不支持该字段，请勿使用。](tag:hws_eu,hcso_dt)
	CookieName *string `json:"cookie_name,omitempty"`

	// 参数解释：会话保持类型。  [约束限制： - 当pool的protocol为TCP、UDP，无论type取值如何，都会被忽略，会话保持只按SOURCE_IP生效； - 当pool的protocol为HTTP、HTTPS时。type可以为HTTP_COOKIE和APP_COOKIE，其他取值会话保持失效。 - 若pool的protocol为QUIC，则必须开启session_persistence且type为SOURCE_IP。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt)  [约束限制： - 当pool的protocol为TCP、UDP，无论type取值如何，都会被忽略，会话保持只按SOURCE_IP生效； - 当pool的protocol为HTTP、HTTPS时。type只能为HTTP_COOKIE，其他取值会话保持失效。](tag:hws_eu,hcso_dt)  取值范围：SOURCE_IP、HTTP_COOKIE、APP_COOKIE。  [荷兰region不支持QUIC。](tag:dt,dt_test) [不支持QUIC。](tag:tm)
	Type CreatePoolSessionPersistenceOptionType `json:"type"`

	// 参数解释：会话保持的时间。当type为APP_COOKIE时不生效。  取值范围： - 若pool的protocol为TCP、UDP则范围为[1,60]（分钟），默认值1； - 若pool的protocol为HTTP和HTTPS则范围为[1,1440]（分钟），默认值1440。
	PersistenceTimeout *int32 `json:"persistence_timeout,omitempty"`
}

func (o CreatePoolSessionPersistenceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolSessionPersistenceOption struct{}"
	}

	return strings.Join([]string{"CreatePoolSessionPersistenceOption", string(data)}, " ")
}

type CreatePoolSessionPersistenceOptionType struct {
	value string
}

type CreatePoolSessionPersistenceOptionTypeEnum struct {
	SOURCE_IP   CreatePoolSessionPersistenceOptionType
	HTTP_COOKIE CreatePoolSessionPersistenceOptionType
	APP_COOKIE  CreatePoolSessionPersistenceOptionType
}

func GetCreatePoolSessionPersistenceOptionTypeEnum() CreatePoolSessionPersistenceOptionTypeEnum {
	return CreatePoolSessionPersistenceOptionTypeEnum{
		SOURCE_IP: CreatePoolSessionPersistenceOptionType{
			value: "SOURCE_IP",
		},
		HTTP_COOKIE: CreatePoolSessionPersistenceOptionType{
			value: "HTTP_COOKIE",
		},
		APP_COOKIE: CreatePoolSessionPersistenceOptionType{
			value: "APP_COOKIE",
		},
	}
}

func (c CreatePoolSessionPersistenceOptionType) Value() string {
	return c.value
}

func (c CreatePoolSessionPersistenceOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePoolSessionPersistenceOptionType) UnmarshalJSON(b []byte) error {
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
