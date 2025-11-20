package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAddonPrecheckTasksRequest Request Object
type ListAddonPrecheckTasksRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`

	// **参数解释：** 根据插件检查类型筛选结果 **约束限制：** 不涉及 **取值范围：** - addonStatic: 运行中插件巡检 - addonUpgrade: 插件升级前检查  **默认取值：** 不涉及
	Type *ListAddonPrecheckTasksRequestType `json:"type,omitempty"`

	// **参数解释：** 根据插件检查任务ID筛选结果，插件检查任务ID可以通过[批量创建插件检查任务](BatchCreateAddonPrecheck.xml)中的status.items[].metadata.taskID字段获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释：** 根据插件实例ID筛选结果，实例ID可以通过[获取AddonInstance列表](cce_02_0326.xml)中的items[].metadata.uid字段获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AddonInstanceId *string `json:"addon_instance_id,omitempty"`
}

func (o ListAddonPrecheckTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddonPrecheckTasksRequest struct{}"
	}

	return strings.Join([]string{"ListAddonPrecheckTasksRequest", string(data)}, " ")
}

type ListAddonPrecheckTasksRequestType struct {
	value string
}

type ListAddonPrecheckTasksRequestTypeEnum struct {
	ADDON_STATIC  ListAddonPrecheckTasksRequestType
	ADDON_UPGRADE ListAddonPrecheckTasksRequestType
}

func GetListAddonPrecheckTasksRequestTypeEnum() ListAddonPrecheckTasksRequestTypeEnum {
	return ListAddonPrecheckTasksRequestTypeEnum{
		ADDON_STATIC: ListAddonPrecheckTasksRequestType{
			value: "addonStatic",
		},
		ADDON_UPGRADE: ListAddonPrecheckTasksRequestType{
			value: "addonUpgrade",
		},
	}
}

func (c ListAddonPrecheckTasksRequestType) Value() string {
	return c.value
}

func (c ListAddonPrecheckTasksRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAddonPrecheckTasksRequestType) UnmarshalJSON(b []byte) error {
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
