package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAuthmodeDetailResponse Response Object
type ShowAuthmodeDetailResponse struct {

	// **参数解释**：账号domainId。 **取值范围**：不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**：授权模式。 **取值范围**： - strict：严格模式。 - loose：非严格模式。
	Mode           *ShowAuthmodeDetailResponseMode `json:"mode,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowAuthmodeDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuthmodeDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowAuthmodeDetailResponse", string(data)}, " ")
}

type ShowAuthmodeDetailResponseMode struct {
	value string
}

type ShowAuthmodeDetailResponseModeEnum struct {
	LOOSE  ShowAuthmodeDetailResponseMode
	STRICT ShowAuthmodeDetailResponseMode
}

func GetShowAuthmodeDetailResponseModeEnum() ShowAuthmodeDetailResponseModeEnum {
	return ShowAuthmodeDetailResponseModeEnum{
		LOOSE: ShowAuthmodeDetailResponseMode{
			value: "loose",
		},
		STRICT: ShowAuthmodeDetailResponseMode{
			value: "strict",
		},
	}
}

func (c ShowAuthmodeDetailResponseMode) Value() string {
	return c.value
}

func (c ShowAuthmodeDetailResponseMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAuthmodeDetailResponseMode) UnmarshalJSON(b []byte) error {
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
