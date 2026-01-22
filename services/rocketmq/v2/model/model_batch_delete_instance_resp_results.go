package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchDeleteInstanceRespResults struct {

	// **参数解释**： 操作结果。 **约束限制**： 不涉及。 **取值范围**： - success - failed **默认取值**： 不涉及。
	Result *BatchDeleteInstanceRespResultsResult `json:"result,omitempty"`

	// 实例ID。
	Instance *string `json:"instance,omitempty"`
}

func (o BatchDeleteInstanceRespResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceRespResults struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceRespResults", string(data)}, " ")
}

type BatchDeleteInstanceRespResultsResult struct {
	value string
}

type BatchDeleteInstanceRespResultsResultEnum struct {
	SUCCESS BatchDeleteInstanceRespResultsResult
	FAILED  BatchDeleteInstanceRespResultsResult
}

func GetBatchDeleteInstanceRespResultsResultEnum() BatchDeleteInstanceRespResultsResultEnum {
	return BatchDeleteInstanceRespResultsResultEnum{
		SUCCESS: BatchDeleteInstanceRespResultsResult{
			value: "success",
		},
		FAILED: BatchDeleteInstanceRespResultsResult{
			value: "failed",
		},
	}
}

func (c BatchDeleteInstanceRespResultsResult) Value() string {
	return c.value
}

func (c BatchDeleteInstanceRespResultsResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteInstanceRespResultsResult) UnmarshalJSON(b []byte) error {
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
