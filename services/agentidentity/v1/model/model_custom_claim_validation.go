package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CustomClaimValidation struct {
	AuthorizingClaimMatchValue *AuthorizingClaimMatchValue `json:"authorizing_claim_match_value"`

	// The name of the custom claim field to check.
	InboundTokenClaimName string `json:"inbound_token_claim_name"`

	// The data type of the claim value to check for.
	InboundTokenClaimValueType CustomClaimValidationInboundTokenClaimValueType `json:"inbound_token_claim_value_type"`
}

func (o CustomClaimValidation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomClaimValidation struct{}"
	}

	return strings.Join([]string{"CustomClaimValidation", string(data)}, " ")
}

type CustomClaimValidationInboundTokenClaimValueType struct {
	value string
}

type CustomClaimValidationInboundTokenClaimValueTypeEnum struct {
	STRING       CustomClaimValidationInboundTokenClaimValueType
	STRING_ARRAY CustomClaimValidationInboundTokenClaimValueType
}

func GetCustomClaimValidationInboundTokenClaimValueTypeEnum() CustomClaimValidationInboundTokenClaimValueTypeEnum {
	return CustomClaimValidationInboundTokenClaimValueTypeEnum{
		STRING: CustomClaimValidationInboundTokenClaimValueType{
			value: "STRING",
		},
		STRING_ARRAY: CustomClaimValidationInboundTokenClaimValueType{
			value: "STRING_ARRAY",
		},
	}
}

func (c CustomClaimValidationInboundTokenClaimValueType) Value() string {
	return c.value
}

func (c CustomClaimValidationInboundTokenClaimValueType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomClaimValidationInboundTokenClaimValueType) UnmarshalJSON(b []byte) error {
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
