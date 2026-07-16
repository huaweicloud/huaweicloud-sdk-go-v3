package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAuthModeResponse Response Object
type UpdateAuthModeResponse struct {

	// **参数解释**：账号domainId。 **取值范围**：不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**：授权模式。 **取值范围**： - strict：严格模式。 - loose：非严格模式。
	Mode           *UpdateAuthModeResponseMode `json:"mode,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o UpdateAuthModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthModeResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuthModeResponse", string(data)}, " ")
}

type UpdateAuthModeResponseMode struct {
	value string
}

type UpdateAuthModeResponseModeEnum struct {
	LOOSE  UpdateAuthModeResponseMode
	STRICT UpdateAuthModeResponseMode
}

func GetUpdateAuthModeResponseModeEnum() UpdateAuthModeResponseModeEnum {
	return UpdateAuthModeResponseModeEnum{
		LOOSE: UpdateAuthModeResponseMode{
			value: "loose",
		},
		STRICT: UpdateAuthModeResponseMode{
			value: "strict",
		},
	}
}

func (c UpdateAuthModeResponseMode) Value() string {
	return c.value
}

func (c UpdateAuthModeResponseMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAuthModeResponseMode) UnmarshalJSON(b []byte) error {
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
