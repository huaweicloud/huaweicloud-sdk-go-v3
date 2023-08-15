package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RecognizeReceptorPocketMode 口袋识别的模式：自动、全局、配体、残基
type RecognizeReceptorPocketMode struct {
	value string
}

type RecognizeReceptorPocketModeEnum struct {
	AUTO     RecognizeReceptorPocketMode
	GLOBAL   RecognizeReceptorPocketMode
	LIGAND   RecognizeReceptorPocketMode
	RESIDUES RecognizeReceptorPocketMode
}

func GetRecognizeReceptorPocketModeEnum() RecognizeReceptorPocketModeEnum {
	return RecognizeReceptorPocketModeEnum{
		AUTO: RecognizeReceptorPocketMode{
			value: "AUTO",
		},
		GLOBAL: RecognizeReceptorPocketMode{
			value: "GLOBAL",
		},
		LIGAND: RecognizeReceptorPocketMode{
			value: "LIGAND",
		},
		RESIDUES: RecognizeReceptorPocketMode{
			value: "RESIDUES",
		},
	}
}

func (c RecognizeReceptorPocketMode) Value() string {
	return c.value
}

func (c RecognizeReceptorPocketMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecognizeReceptorPocketMode) UnmarshalJSON(b []byte) error {
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
