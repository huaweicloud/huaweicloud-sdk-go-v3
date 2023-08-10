package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LigandSimilarityGraphMode 创建配体相似度图的模式：中心模式
type LigandSimilarityGraphMode struct {
	value string
}

type LigandSimilarityGraphModeEnum struct {
	CENTER LigandSimilarityGraphMode
}

func GetLigandSimilarityGraphModeEnum() LigandSimilarityGraphModeEnum {
	return LigandSimilarityGraphModeEnum{
		CENTER: LigandSimilarityGraphMode{
			value: "CENTER",
		},
	}
}

func (c LigandSimilarityGraphMode) Value() string {
	return c.value
}

func (c LigandSimilarityGraphMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LigandSimilarityGraphMode) UnmarshalJSON(b []byte) error {
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
