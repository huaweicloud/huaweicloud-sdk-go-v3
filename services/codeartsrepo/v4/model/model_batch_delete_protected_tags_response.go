package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteProtectedTagsResponse Response Object
type BatchDeleteProtectedTagsResponse struct {

	// **参数解释：** 状态码。 - success，表示接口请求成功。 - fail，表示接口请求失败。
	Status         *BatchDeleteProtectedTagsResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o BatchDeleteProtectedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteProtectedTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteProtectedTagsResponse", string(data)}, " ")
}

type BatchDeleteProtectedTagsResponseStatus struct {
	value string
}

type BatchDeleteProtectedTagsResponseStatusEnum struct {
	SUCCESS BatchDeleteProtectedTagsResponseStatus
	FAIL    BatchDeleteProtectedTagsResponseStatus
}

func GetBatchDeleteProtectedTagsResponseStatusEnum() BatchDeleteProtectedTagsResponseStatusEnum {
	return BatchDeleteProtectedTagsResponseStatusEnum{
		SUCCESS: BatchDeleteProtectedTagsResponseStatus{
			value: "success",
		},
		FAIL: BatchDeleteProtectedTagsResponseStatus{
			value: "fail",
		},
	}
}

func (c BatchDeleteProtectedTagsResponseStatus) Value() string {
	return c.value
}

func (c BatchDeleteProtectedTagsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteProtectedTagsResponseStatus) UnmarshalJSON(b []byte) error {
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
