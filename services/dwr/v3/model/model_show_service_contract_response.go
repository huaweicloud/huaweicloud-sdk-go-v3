package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowServiceContractResponse Response Object
type ShowServiceContractResponse struct {

	// 服务协议类型,默认为use_public_action_privacy_statement。
	AgreementType *ShowServiceContractResponseAgreementType `json:"agreement_type,omitempty"`

	// 同意时间。
	CreateTime *string `json:"create_time,omitempty"`

	XRequestId *string `json:"x-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowServiceContractResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServiceContractResponse struct{}"
	}

	return strings.Join([]string{"ShowServiceContractResponse", string(data)}, " ")
}

type ShowServiceContractResponseAgreementType struct {
	value string
}

type ShowServiceContractResponseAgreementTypeEnum struct {
	USE_PUBLIC_ACTION_PRIVACY_STATEMENT ShowServiceContractResponseAgreementType
}

func GetShowServiceContractResponseAgreementTypeEnum() ShowServiceContractResponseAgreementTypeEnum {
	return ShowServiceContractResponseAgreementTypeEnum{
		USE_PUBLIC_ACTION_PRIVACY_STATEMENT: ShowServiceContractResponseAgreementType{
			value: "use_public_action_privacy_statement",
		},
	}
}

func (c ShowServiceContractResponseAgreementType) Value() string {
	return c.value
}

func (c ShowServiceContractResponseAgreementType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowServiceContractResponseAgreementType) UnmarshalJSON(b []byte) error {
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
