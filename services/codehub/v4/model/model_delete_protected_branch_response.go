package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteProtectedBranchResponse Response Object
type DeleteProtectedBranchResponse struct {

	// **参数解释：** 状态码。 - success，表示接口请求成功。 - fail，表示接口请求失败。
	Status         *DeleteProtectedBranchResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o DeleteProtectedBranchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtectedBranchResponse struct{}"
	}

	return strings.Join([]string{"DeleteProtectedBranchResponse", string(data)}, " ")
}

type DeleteProtectedBranchResponseStatus struct {
	value string
}

type DeleteProtectedBranchResponseStatusEnum struct {
	SUCCESS DeleteProtectedBranchResponseStatus
	FAIL    DeleteProtectedBranchResponseStatus
}

func GetDeleteProtectedBranchResponseStatusEnum() DeleteProtectedBranchResponseStatusEnum {
	return DeleteProtectedBranchResponseStatusEnum{
		SUCCESS: DeleteProtectedBranchResponseStatus{
			value: "success",
		},
		FAIL: DeleteProtectedBranchResponseStatus{
			value: "fail",
		},
	}
}

func (c DeleteProtectedBranchResponseStatus) Value() string {
	return c.value
}

func (c DeleteProtectedBranchResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteProtectedBranchResponseStatus) UnmarshalJSON(b []byte) error {
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
