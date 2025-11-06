package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteProtectedBranchesResponse Response Object
type BatchDeleteProtectedBranchesResponse struct {

	// **参数解释：** 状态码。 - success，表示接口请求成功。 - fail，表示接口请求失败。
	Status         *BatchDeleteProtectedBranchesResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o BatchDeleteProtectedBranchesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteProtectedBranchesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteProtectedBranchesResponse", string(data)}, " ")
}

type BatchDeleteProtectedBranchesResponseStatus struct {
	value string
}

type BatchDeleteProtectedBranchesResponseStatusEnum struct {
	SUCCESS BatchDeleteProtectedBranchesResponseStatus
	FAIL    BatchDeleteProtectedBranchesResponseStatus
}

func GetBatchDeleteProtectedBranchesResponseStatusEnum() BatchDeleteProtectedBranchesResponseStatusEnum {
	return BatchDeleteProtectedBranchesResponseStatusEnum{
		SUCCESS: BatchDeleteProtectedBranchesResponseStatus{
			value: "success",
		},
		FAIL: BatchDeleteProtectedBranchesResponseStatus{
			value: "fail",
		},
	}
}

func (c BatchDeleteProtectedBranchesResponseStatus) Value() string {
	return c.value
}

func (c BatchDeleteProtectedBranchesResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteProtectedBranchesResponseStatus) UnmarshalJSON(b []byte) error {
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
