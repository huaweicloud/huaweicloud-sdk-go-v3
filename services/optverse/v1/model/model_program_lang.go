package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ProgramLang **参数解释**： 编程语言，包含python，c++，java **约束限制**： 不涉及 **取值范围**： * python： python编程语言。 * c++:   c++编程语言。 * java:  java编程语言。 **默认取值**： 不涉及
type ProgramLang struct {
	value string
}

type ProgramLangEnum struct {
	C      ProgramLang
	PYTHON ProgramLang
	JAVA   ProgramLang
}

func GetProgramLangEnum() ProgramLangEnum {
	return ProgramLangEnum{
		C: ProgramLang{
			value: "c++",
		},
		PYTHON: ProgramLang{
			value: "python",
		},
		JAVA: ProgramLang{
			value: "java",
		},
	}
}

func (c ProgramLang) Value() string {
	return c.value
}

func (c ProgramLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProgramLang) UnmarshalJSON(b []byte) error {
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
