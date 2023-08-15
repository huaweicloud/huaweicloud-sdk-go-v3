package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StackStatusPrimitiveTypeHolder struct {

	// 资源栈的状态    * `CREATION_COMPLETE` - 生成空资源栈完成，并没有任何部署    * `DEPLOYMENT_IN_PROGRESS` - 正在部署，请等待    * `DEPLOYMENT_FAILED` - 部署失败。请从status_message获取错误信息汇总，或者调用ListStackEvents获得事件详情    * `DEPLOYMENT_COMPLETE` - 部署完成    * `ROLLBACK_IN_PROGRESS` - 部署失败，正在回滚，请等待    * `ROLLBACK_FAILED` - 回滚失败。请从status_message获取错误信息汇总，或者调用ListStackEvents获得事件详情    * `ROLLBACK_COMPLETE` - 回滚完成    * `DELETION_IN_PROGRESS` - 正在删除，请等待    * `DELETION_FAILED` - 删除失败。请从status_message获取错误信息汇总，或者调用ListStackEvents获得事件详情
	Status *StackStatusPrimitiveTypeHolderStatus `json:"status,omitempty"`
}

func (o StackStatusPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackStatusPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackStatusPrimitiveTypeHolder", string(data)}, " ")
}

type StackStatusPrimitiveTypeHolderStatus struct {
	value string
}

type StackStatusPrimitiveTypeHolderStatusEnum struct {
	CREATION_COMPLETE      StackStatusPrimitiveTypeHolderStatus
	DEPLOYMENT_IN_PROGRESS StackStatusPrimitiveTypeHolderStatus
	DEPLOYMENT_FAILED      StackStatusPrimitiveTypeHolderStatus
	DEPLOYMENT_COMPLETE    StackStatusPrimitiveTypeHolderStatus
	ROLLBACK_IN_PROGRESS   StackStatusPrimitiveTypeHolderStatus
	ROLLBACK_FAILED        StackStatusPrimitiveTypeHolderStatus
	ROLLBACK_COMPLETE      StackStatusPrimitiveTypeHolderStatus
	DELETION_IN_PROGRESS   StackStatusPrimitiveTypeHolderStatus
	DELETION_FAILED        StackStatusPrimitiveTypeHolderStatus
}

func GetStackStatusPrimitiveTypeHolderStatusEnum() StackStatusPrimitiveTypeHolderStatusEnum {
	return StackStatusPrimitiveTypeHolderStatusEnum{
		CREATION_COMPLETE: StackStatusPrimitiveTypeHolderStatus{
			value: "CREATION_COMPLETE",
		},
		DEPLOYMENT_IN_PROGRESS: StackStatusPrimitiveTypeHolderStatus{
			value: "DEPLOYMENT_IN_PROGRESS",
		},
		DEPLOYMENT_FAILED: StackStatusPrimitiveTypeHolderStatus{
			value: "DEPLOYMENT_FAILED",
		},
		DEPLOYMENT_COMPLETE: StackStatusPrimitiveTypeHolderStatus{
			value: "DEPLOYMENT_COMPLETE",
		},
		ROLLBACK_IN_PROGRESS: StackStatusPrimitiveTypeHolderStatus{
			value: "ROLLBACK_IN_PROGRESS",
		},
		ROLLBACK_FAILED: StackStatusPrimitiveTypeHolderStatus{
			value: "ROLLBACK_FAILED",
		},
		ROLLBACK_COMPLETE: StackStatusPrimitiveTypeHolderStatus{
			value: "ROLLBACK_COMPLETE",
		},
		DELETION_IN_PROGRESS: StackStatusPrimitiveTypeHolderStatus{
			value: "DELETION_IN_PROGRESS",
		},
		DELETION_FAILED: StackStatusPrimitiveTypeHolderStatus{
			value: "DELETION_FAILED",
		},
	}
}

func (c StackStatusPrimitiveTypeHolderStatus) Value() string {
	return c.value
}

func (c StackStatusPrimitiveTypeHolderStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackStatusPrimitiveTypeHolderStatus) UnmarshalJSON(b []byte) error {
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
