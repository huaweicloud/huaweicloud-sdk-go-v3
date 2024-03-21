package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFuncSnapshotStateRequest Request Object
type ShowFuncSnapshotStateRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 查询快照制作开关状态
	Action ShowFuncSnapshotStateRequestAction `json:"action"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ShowFuncSnapshotStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFuncSnapshotStateRequest struct{}"
	}

	return strings.Join([]string{"ShowFuncSnapshotStateRequest", string(data)}, " ")
}

type ShowFuncSnapshotStateRequestAction struct {
	value string
}

type ShowFuncSnapshotStateRequestActionEnum struct {
	STATE           ShowFuncSnapshotStateRequestAction
	ENABLE_SNAPSHOT ShowFuncSnapshotStateRequestAction
}

func GetShowFuncSnapshotStateRequestActionEnum() ShowFuncSnapshotStateRequestActionEnum {
	return ShowFuncSnapshotStateRequestActionEnum{
		STATE: ShowFuncSnapshotStateRequestAction{
			value: "state",
		},
		ENABLE_SNAPSHOT: ShowFuncSnapshotStateRequestAction{
			value: "enableSnapshot",
		},
	}
}

func (c ShowFuncSnapshotStateRequestAction) Value() string {
	return c.value
}

func (c ShowFuncSnapshotStateRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFuncSnapshotStateRequestAction) UnmarshalJSON(b []byte) error {
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
