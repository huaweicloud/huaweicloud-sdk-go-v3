package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// stack resource api model
type StackResource struct {

	// 资源的物理id，由资源提供服务的provider在资源部署的时候生成
	PhysicalResourceId *string `json:"physical_resource_id,omitempty"`

	// 资源的物理名称，资源提供服务在资源部署的时候给予
	PhysicalResourceName *string `json:"physical_resource_name,omitempty"`

	// 资源名，是用户在模板中定义的
	LogicalResourceName *string `json:"logical_resource_name,omitempty"`

	// 资源的类型，是用户在模板中定义的
	LogicalResourceType *string `json:"logical_resource_type,omitempty"`

	// 此次事件的类型 * `CREATION_IN_PROGRESS` - 正在生成 * `CREATION_FAILED`      - 生成失败 * `CREATION_COMPLETE`    - 生成完成 * `DELETION_IN_PROGRESS` - 正在删除 * `DELETION_FAILED`      - 删除失败 * `DELETION_COMPLETE`    - 已经删除 * `DELETION_SKIPPED`     - 跳过删除。未来我们将支持，用户可以从资源编排服务中删除，但是不真的删除资源本身 * `UPDATE_IN_PROGRESS`   - 正在更新。此处的更新特指非替换式更新，如果是替换式更新，则使用CREATION后DELETION * `UPDATE_FAILED`        - 更新失败。此处的更新特指非替换式更新，如果是替换式更新，则使用CREATION后DELETION * `UPDATE_COMPLETE`      - 更新完成。此处的更新特指非替换式更新，如果是替换式更新，则使用CREATION后DELETION
	ResourceStatus *StackResourceResourceStatus `json:"resource_status,omitempty"`

	// 如果是成功状态或执行中状态，则没有信息
	StatusMessage *string `json:"status_message,omitempty"`
}

func (o StackResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackResource struct{}"
	}

	return strings.Join([]string{"StackResource", string(data)}, " ")
}

type StackResourceResourceStatus struct {
	value string
}

type StackResourceResourceStatusEnum struct {
	CREATION_IN_PROGRESS StackResourceResourceStatus
	CREATION_FAILED      StackResourceResourceStatus
	CREATION_COMPLETE    StackResourceResourceStatus
	DELETION_IN_PROGRESS StackResourceResourceStatus
	DELETION_FAILED      StackResourceResourceStatus
	DELETION_COMPLETE    StackResourceResourceStatus
	DELETION_SKIPPED     StackResourceResourceStatus
	UPDATE_IN_PROGRESS   StackResourceResourceStatus
	UPDATE_FAILED        StackResourceResourceStatus
	UPDATE_COMPLETE      StackResourceResourceStatus
}

func GetStackResourceResourceStatusEnum() StackResourceResourceStatusEnum {
	return StackResourceResourceStatusEnum{
		CREATION_IN_PROGRESS: StackResourceResourceStatus{
			value: "CREATION_IN_PROGRESS",
		},
		CREATION_FAILED: StackResourceResourceStatus{
			value: "CREATION_FAILED",
		},
		CREATION_COMPLETE: StackResourceResourceStatus{
			value: "CREATION_COMPLETE",
		},
		DELETION_IN_PROGRESS: StackResourceResourceStatus{
			value: "DELETION_IN_PROGRESS",
		},
		DELETION_FAILED: StackResourceResourceStatus{
			value: "DELETION_FAILED",
		},
		DELETION_COMPLETE: StackResourceResourceStatus{
			value: "DELETION_COMPLETE",
		},
		DELETION_SKIPPED: StackResourceResourceStatus{
			value: "DELETION_SKIPPED",
		},
		UPDATE_IN_PROGRESS: StackResourceResourceStatus{
			value: "UPDATE_IN_PROGRESS",
		},
		UPDATE_FAILED: StackResourceResourceStatus{
			value: "UPDATE_FAILED",
		},
		UPDATE_COMPLETE: StackResourceResourceStatus{
			value: "UPDATE_COMPLETE",
		},
	}
}

func (c StackResourceResourceStatus) Value() string {
	return c.value
}

func (c StackResourceResourceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackResourceResourceStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
