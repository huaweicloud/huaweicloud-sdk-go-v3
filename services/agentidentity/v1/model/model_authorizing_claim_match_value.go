package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AuthorizingClaimMatchValue struct {

	// Defines the relationship between the claim field value and the value or values you're matching for.
	ClaimMatchOperator AuthorizingClaimMatchValueClaimMatchOperator `json:"claim_match_operator"`

	ClaimMatchValue *ClaimMatchValue `json:"claim_match_value"`
}

func (o AuthorizingClaimMatchValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizingClaimMatchValue struct{}"
	}

	return strings.Join([]string{"AuthorizingClaimMatchValue", string(data)}, " ")
}

type AuthorizingClaimMatchValueClaimMatchOperator struct {
	value string
}

type AuthorizingClaimMatchValueClaimMatchOperatorEnum struct {
	EQUALS       AuthorizingClaimMatchValueClaimMatchOperator
	CONTAINS     AuthorizingClaimMatchValueClaimMatchOperator
	CONTAINS_ANY AuthorizingClaimMatchValueClaimMatchOperator
}

func GetAuthorizingClaimMatchValueClaimMatchOperatorEnum() AuthorizingClaimMatchValueClaimMatchOperatorEnum {
	return AuthorizingClaimMatchValueClaimMatchOperatorEnum{
		EQUALS: AuthorizingClaimMatchValueClaimMatchOperator{
			value: "EQUALS",
		},
		CONTAINS: AuthorizingClaimMatchValueClaimMatchOperator{
			value: "CONTAINS",
		},
		CONTAINS_ANY: AuthorizingClaimMatchValueClaimMatchOperator{
			value: "CONTAINS_ANY",
		},
	}
}

func (c AuthorizingClaimMatchValueClaimMatchOperator) Value() string {
	return c.value
}

func (c AuthorizingClaimMatchValueClaimMatchOperator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizingClaimMatchValueClaimMatchOperator) UnmarshalJSON(b []byte) error {
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
