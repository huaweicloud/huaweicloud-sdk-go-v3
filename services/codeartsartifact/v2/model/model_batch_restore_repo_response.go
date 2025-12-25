package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchRestoreRepoResponse Response Object
type BatchRestoreRepoResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *BatchRestoreRepoResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	Result         *TrashcanResultData `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o BatchRestoreRepoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestoreRepoResponse struct{}"
	}

	return strings.Join([]string{"BatchRestoreRepoResponse", string(data)}, " ")
}

type BatchRestoreRepoResponseStatus struct {
	value string
}

type BatchRestoreRepoResponseStatusEnum struct {
	SUCCESS BatchRestoreRepoResponseStatus
	ERROR   BatchRestoreRepoResponseStatus
}

func GetBatchRestoreRepoResponseStatusEnum() BatchRestoreRepoResponseStatusEnum {
	return BatchRestoreRepoResponseStatusEnum{
		SUCCESS: BatchRestoreRepoResponseStatus{
			value: "success",
		},
		ERROR: BatchRestoreRepoResponseStatus{
			value: "error",
		},
	}
}

func (c BatchRestoreRepoResponseStatus) Value() string {
	return c.value
}

func (c BatchRestoreRepoResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchRestoreRepoResponseStatus) UnmarshalJSON(b []byte) error {
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
