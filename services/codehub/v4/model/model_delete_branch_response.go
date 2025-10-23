package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteBranchResponse Response Object
type DeleteBranchResponse struct {

	// **参数解释：** 状态码。 - success，表示接口请求成功。 - fail，表示接口请求失败。
	Status         *DeleteBranchResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o DeleteBranchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBranchResponse struct{}"
	}

	return strings.Join([]string{"DeleteBranchResponse", string(data)}, " ")
}

type DeleteBranchResponseStatus struct {
	value string
}

type DeleteBranchResponseStatusEnum struct {
	SUCCESS DeleteBranchResponseStatus
	FAIL    DeleteBranchResponseStatus
}

func GetDeleteBranchResponseStatusEnum() DeleteBranchResponseStatusEnum {
	return DeleteBranchResponseStatusEnum{
		SUCCESS: DeleteBranchResponseStatus{
			value: "success",
		},
		FAIL: DeleteBranchResponseStatus{
			value: "fail",
		},
	}
}

func (c DeleteBranchResponseStatus) Value() string {
	return c.value
}

func (c DeleteBranchResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteBranchResponseStatus) UnmarshalJSON(b []byte) error {
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
