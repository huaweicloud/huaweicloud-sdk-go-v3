package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 会话持久性对象。
type CreatePoolSessionPersistenceOption struct {

	// cookie名称。  格式：仅支持字母、数字、中划线(-)、下划线(_)和点号(.)。  使用说明： - 只有当type为APP_COOKIE时才有效。其他情况下传该字段会报错。
	CookieName *string `json:"cookie_name,omitempty"`

	// 会话保持类型。 取值范围：SOURCE_IP、HTTP_COOKIE、APP_COOKIE。 [使用说明：  - 当pool的protocol为TCP、UDP，无论type取值如何，都会被忽略，会话保持只按SOURCE_IP生效；  - 当pool的protocol为HTTP、HTTPS时。如果是独享型负载均衡器的pool，则type只能为HTTP_COOKIE，其他取值会话保持失效。如果是共享型负载均衡器的pool，则type可以为HTTP_COOKIE和APP_COOKIE，其他取值会话保持失效。](tag:hws,hws_hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42,dt,dt_test) [使用说明：  - 当pool的protocol为TCP、UDP，无论type取值如何，都会被忽略，会话保持只按SOURCE_IP生效；  - 当pool的protocol为HTTP、HTTPS时。type只能为HTTP_COOKIE，其他取值会话保持失效。](tag:hcso_dt)
	Type CreatePoolSessionPersistenceOptionType `json:"type"`

	// 会话保持的时间。当type为APP_COOKIE时不生效。  适用范围：如果pool的protocol为TCP、UDP则范围为[1,60]（分钟），默认值1；如果pool的protocol为HTTP和HTTPS则范围为[1,1440]（分钟），默认值1440。
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
