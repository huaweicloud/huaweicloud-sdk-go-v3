package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CloneServer 克隆服务器类
type CloneServer struct {

	// 克隆服务器ID
	VmId *string `json:"vm_id,omitempty"`

	// 克隆虚拟机的名称
	Name *string `json:"name,omitempty"`

	// 克隆错误信息
	CloneError *string `json:"clone_error,omitempty"`

	// 克隆状态 NOT_READY: 未准备好 READY: 就绪 PREPARING: 准备创建中 CREATING: 创建中 ERROR: 克隆错误 FINISHED: 克隆结束
	CloneState *CloneServerCloneState `json:"clone_state,omitempty"`

	// 克隆错误信息描述
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o CloneServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneServer struct{}"
	}

	return strings.Join([]string{"CloneServer", string(data)}, " ")
}

type CloneServerCloneState struct {
	value string
}

type CloneServerCloneStateEnum struct {
	NOT_READY CloneServerCloneState
	READY     CloneServerCloneState
	PREPARING CloneServerCloneState
	CREATING  CloneServerCloneState
	ERROR     CloneServerCloneState
	FINISHED  CloneServerCloneState
}

func GetCloneServerCloneStateEnum() CloneServerCloneStateEnum {
	return CloneServerCloneStateEnum{
		NOT_READY: CloneServerCloneState{
			value: "NOT_READY",
		},
		READY: CloneServerCloneState{
			value: "READY",
		},
		PREPARING: CloneServerCloneState{
			value: "PREPARING",
		},
		CREATING: CloneServerCloneState{
			value: "CREATING",
		},
		ERROR: CloneServerCloneState{
			value: "ERROR",
		},
		FINISHED: CloneServerCloneState{
			value: "FINISHED",
		},
	}
}

func (c CloneServerCloneState) Value() string {
	return c.value
}

func (c CloneServerCloneState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloneServerCloneState) UnmarshalJSON(b []byte) error {
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
