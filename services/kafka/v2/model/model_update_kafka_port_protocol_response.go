package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateKafkaPortProtocolResponse Response Object
type UpdateKafkaPortProtocolResponse struct {

	// 后台任务id。
	JobId *string `json:"job_id,omitempty"`

	// 开启或者关闭的Kafka接入方式。
	Protocol *UpdateKafkaPortProtocolResponseProtocol `json:"protocol,omitempty"`

	// 开启动作或者关闭动作。
	Enable         *bool `json:"enable,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateKafkaPortProtocolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKafkaPortProtocolResponse struct{}"
	}

	return strings.Join([]string{"UpdateKafkaPortProtocolResponse", string(data)}, " ")
}

type UpdateKafkaPortProtocolResponseProtocol struct {
	value string
}

type UpdateKafkaPortProtocolResponseProtocolEnum struct {
	PRIVATE_PLAIN_ENABLE          UpdateKafkaPortProtocolResponseProtocol
	PRIVATE_SASL_SSL_ENABLE       UpdateKafkaPortProtocolResponseProtocol
	PRIVATE_SASL_PLAINTEXT_ENABLE UpdateKafkaPortProtocolResponseProtocol
	PUBLIC_PLAIN_ENABLE           UpdateKafkaPortProtocolResponseProtocol
	PUBLIC_SASL_SSL_ENABLE        UpdateKafkaPortProtocolResponseProtocol
	PUBLIC_SASL_PLAINTEXT_ENABLE  UpdateKafkaPortProtocolResponseProtocol
}

func GetUpdateKafkaPortProtocolResponseProtocolEnum() UpdateKafkaPortProtocolResponseProtocolEnum {
	return UpdateKafkaPortProtocolResponseProtocolEnum{
		PRIVATE_PLAIN_ENABLE: UpdateKafkaPortProtocolResponseProtocol{
			value: "private_plain_enable",
		},
		PRIVATE_SASL_SSL_ENABLE: UpdateKafkaPortProtocolResponseProtocol{
			value: "private_sasl_ssl_enable",
		},
		PRIVATE_SASL_PLAINTEXT_ENABLE: UpdateKafkaPortProtocolResponseProtocol{
			value: "private_sasl_plaintext_enable",
		},
		PUBLIC_PLAIN_ENABLE: UpdateKafkaPortProtocolResponseProtocol{
			value: "public_plain_enable",
		},
		PUBLIC_SASL_SSL_ENABLE: UpdateKafkaPortProtocolResponseProtocol{
			value: "public_sasl_ssl_enable",
		},
		PUBLIC_SASL_PLAINTEXT_ENABLE: UpdateKafkaPortProtocolResponseProtocol{
			value: "public_sasl_plaintext_enable",
		},
	}
}

func (c UpdateKafkaPortProtocolResponseProtocol) Value() string {
	return c.value
}

func (c UpdateKafkaPortProtocolResponseProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateKafkaPortProtocolResponseProtocol) UnmarshalJSON(b []byte) error {
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
