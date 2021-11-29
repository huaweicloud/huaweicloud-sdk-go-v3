package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 会话持久性对象。
type UpdatePoolSessionPersistenceOption struct {
	// cookie名称。  格式：仅支持字母、数字、中划线(-)、下划线(_)和点号(.)。  使用说明： - 只有当type为APP_COOKIE时才有效。其他情况下传该字段会报错。

	CookieName *string `json:"cookie_name,omitempty"`
	// 描述：类型，可以为SOURCE_IP、HTTP_COOKIE、APP_COOKIE。  使用说明： - 当pool的protocol为TCP、UDP、QUIC时，只按SOURCE_IP生效； - 当pool的protocol为HTTP、HTTPS时，只按HTTP_COOKIE、APP_COOKIE生效。

	Type *UpdatePoolSessionPersistenceOptionType `json:"type,omitempty"`
	// 会话保持的时间。当type为APP_COOKIE时不生效。 适用范围：如果pool的protocol为TCP、UDP和QUIC则范围为[1,60]（分钟），默认值1；如果pool的protocol为HTTP和HTTPS则范围为[1,1440]（分钟），默认值1440。

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

func (c UpdatePoolSessionPersistenceOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePoolSessionPersistenceOptionType) UnmarshalJSON(b []byte) error {
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
