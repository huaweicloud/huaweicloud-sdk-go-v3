package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePoolSessionPersistenceOption 会话持久性对象。
type UpdatePoolSessionPersistenceOption struct {

	// cookie名称。  [共享型LB，支持字母、数字、中划线(-)和下划线(_)，最大长度64个字符。 独享型LB，支持字母、数字、中划线(-)、下划线(_)和点号(.)，最大长度255个字符。](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm) [支持字母、数字、中划线(-)、下划线(_)和点号(.)，最大长度255个字符。](tag:hws_eu,hcso_dt) 使用说明： - 只有当type为APP_COOKIE时才有效。其他情况下传该字段会报错。 [网关型LB，不支持该特性，请勿使用。](tag:hws_eu) [不支持该字段，请勿使用。](tag:hws_eu,hcso_dt)
	CookieName *string `json:"cookie_name,omitempty"`

	// 会话保持类型。  取值范围：SOURCE_IP、HTTP_COOKIE、APP_COOKIE。  [使用说明： - 当pool的protocol为TCP、UDP，无论type取值如何，都会被忽略，会话保持只按SOURCE_IP生效。 - 当pool的protocol为HTTP、HTTPS时。type可以为HTTP_COOKIE和APP_COOKIE，其他取值会话保持失效。 - 若pool的protocol为QUIC，则必须开启session_persistence且type为SOURCE_IP。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt)  [使用说明： - 当pool的protocol为TCP、UDP，无论type取值如何，都会被忽略，会话保持只按SOURCE_IP生效。 - 当pool的protocol为HTTP、HTTPS时。type只能为HTTP_COOKIE， 其他取值会话保持失效。](tag:hws_eu,hcso_dt)  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt,dt_test)
	Type *UpdatePoolSessionPersistenceOptionType `json:"type,omitempty"`

	// 会话保持的时间。当type为APP_COOKIE时不生效。  适用范围：如果pool的protocol为TCP、UDP和QUIC则范围为[1,60]（分钟），默认值1； 如果pool的protocol为HTTP和HTTPS则范围为[1,1440]（分钟），默认值1440。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt,dt_test)
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
