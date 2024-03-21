package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateFuncSnapshotRequest Request Object
type UpdateFuncSnapshotRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 禁用/启用
	Action UpdateFuncSnapshotRequestAction `json:"action"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o UpdateFuncSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFuncSnapshotRequest struct{}"
	}

	return strings.Join([]string{"UpdateFuncSnapshotRequest", string(data)}, " ")
}

type UpdateFuncSnapshotRequestAction struct {
	value string
}

type UpdateFuncSnapshotRequestActionEnum struct {
	ENABLE  UpdateFuncSnapshotRequestAction
	DISABLE UpdateFuncSnapshotRequestAction
}

func GetUpdateFuncSnapshotRequestActionEnum() UpdateFuncSnapshotRequestActionEnum {
	return UpdateFuncSnapshotRequestActionEnum{
		ENABLE: UpdateFuncSnapshotRequestAction{
			value: "enable",
		},
		DISABLE: UpdateFuncSnapshotRequestAction{
			value: "disable",
		},
	}
}

func (c UpdateFuncSnapshotRequestAction) Value() string {
	return c.value
}

func (c UpdateFuncSnapshotRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFuncSnapshotRequestAction) UnmarshalJSON(b []byte) error {
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
