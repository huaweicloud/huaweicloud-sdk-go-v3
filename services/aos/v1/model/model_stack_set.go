package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StackSet struct {

	// 资源栈集（stack_set）的唯一ID。  此ID由资源编排服务在生成资源栈集的时候生成，为UUID。  由于资源栈集名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈集，删除，在重新创建一个同名资源栈集。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈集就是我以为的那个，而不是又其他队友删除后创建的同名资源栈集。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈集所对应的ID都不相同，更新不会影响ID。如果给与的stack_set_id和当前资源栈集的ID不一致，则返回400
	StackSetId *string `json:"stack_set_id,omitempty"`

	// 资源栈集（stack_set）的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackSetName string `json:"stack_set_name"`

	// 资源栈集的描述。可用于客户识别自己的资源栈集。
	StackSetDescription *string `json:"stack_set_description,omitempty"`

	// 权限模型，定义了RFS操作资源栈集时所需委托的创建方式，枚举值    * `SELF_MANAGED` - 基于部署需求，用户需要提前手动创建委托，既包含管理账号给RFS的委托，也包含成员账号创建给管理账号的委托。如果委托不存在或错误，创建资源栈集不会失败，部署资源栈集或部署资源栈实例的时候才会报错。
	PermissionModel *StackSetPermissionModel `json:"permission_model,omitempty"`

	// 资源栈集的状态     * `IDLE` - 资源栈集空闲 * `OPERATION_IN_PROGRESS` - 资源栈集操作中 * `DEACTIVATED` - 资源栈集禁用
	Status *StackSetStatus `json:"status,omitempty"`

	// 资源栈集的创建时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	CreateTime *string `json:"create_time,omitempty"`

	// 资源栈集的更新时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o StackSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSet struct{}"
	}

	return strings.Join([]string{"StackSet", string(data)}, " ")
}

type StackSetPermissionModel struct {
	value string
}

type StackSetPermissionModelEnum struct {
	SELF_MANAGED StackSetPermissionModel
}

func GetStackSetPermissionModelEnum() StackSetPermissionModelEnum {
	return StackSetPermissionModelEnum{
		SELF_MANAGED: StackSetPermissionModel{
			value: "SELF_MANAGED",
		},
	}
}

func (c StackSetPermissionModel) Value() string {
	return c.value
}

func (c StackSetPermissionModel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackSetPermissionModel) UnmarshalJSON(b []byte) error {
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

type StackSetStatus struct {
	value string
}

type StackSetStatusEnum struct {
	IDLE                  StackSetStatus
	OPERATION_IN_PROGRESS StackSetStatus
	DEACTIVATED           StackSetStatus
}

func GetStackSetStatusEnum() StackSetStatusEnum {
	return StackSetStatusEnum{
		IDLE: StackSetStatus{
			value: "IDLE",
		},
		OPERATION_IN_PROGRESS: StackSetStatus{
			value: "OPERATION_IN_PROGRESS",
		},
		DEACTIVATED: StackSetStatus{
			value: "DEACTIVATED",
		},
	}
}

func (c StackSetStatus) Value() string {
	return c.value
}

func (c StackSetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackSetStatus) UnmarshalJSON(b []byte) error {
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
