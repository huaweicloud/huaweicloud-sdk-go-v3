package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CheckTextLanguageResponse Response Object
type CheckTextLanguageResponse struct {

	// 语言检测状态。 * MATCHED: 匹配 * UNMATCHED: 不匹配
	Result         *CheckTextLanguageResponseResult `json:"result,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o CheckTextLanguageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTextLanguageResponse struct{}"
	}

	return strings.Join([]string{"CheckTextLanguageResponse", string(data)}, " ")
}

type CheckTextLanguageResponseResult struct {
	value string
}

type CheckTextLanguageResponseResultEnum struct {
	MATCHED   CheckTextLanguageResponseResult
	UNMATCHED CheckTextLanguageResponseResult
}

func GetCheckTextLanguageResponseResultEnum() CheckTextLanguageResponseResultEnum {
	return CheckTextLanguageResponseResultEnum{
		MATCHED: CheckTextLanguageResponseResult{
			value: "MATCHED",
		},
		UNMATCHED: CheckTextLanguageResponseResult{
			value: "UNMATCHED",
		},
	}
}

func (c CheckTextLanguageResponseResult) Value() string {
	return c.value
}

func (c CheckTextLanguageResponseResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckTextLanguageResponseResult) UnmarshalJSON(b []byte) error {
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
