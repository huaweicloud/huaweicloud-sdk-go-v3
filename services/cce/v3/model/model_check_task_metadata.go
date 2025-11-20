package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CheckTaskMetadata struct {

	// **参数解释：** 插件检查类型 **取值范围：** - addonStatic: 运行中插件巡检 - addonUpgrade: 插件升级前检查
	Type CheckTaskMetadataType `json:"type"`

	// **参数解释：** 插件检查任务ID，用于任务检查结果查询 **取值范围：** 不涉及
	TaskID string `json:"taskID"`

	// **参数解释：** 插件模板名称 **取值范围：** cce服务提供的插件模板，可以通过[查询AddonTemplates列表](cce_02_0321.xml)中的items[].metadata.name字段获取
	AddonTemplateName string `json:"addonTemplateName"`

	// **参数解释：** 插件实例名称 **取值范围：** 不涉及
	AddonInstanceName *string `json:"addonInstanceName,omitempty"`

	// **参数解释：** 插件实例ID **取值范围：** 不涉及
	AddonInstanceID *string `json:"addonInstanceID,omitempty"`

	// **参数解释：** 插件检查任务创建时间 **取值范围：** 不涉及
	CreateTimeStamp string `json:"createTimeStamp"`

	// **参数解释：** 插件检查任务超时时间，仅终态（Failed/Success）任务存在此字段 **取值范围：** 不涉及
	ExpireTimeStamp *string `json:"expireTimeStamp,omitempty"`
}

func (o CheckTaskMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTaskMetadata struct{}"
	}

	return strings.Join([]string{"CheckTaskMetadata", string(data)}, " ")
}

type CheckTaskMetadataType struct {
	value string
}

type CheckTaskMetadataTypeEnum struct {
	ADDON_STATIC  CheckTaskMetadataType
	ADDON_UPGRADE CheckTaskMetadataType
}

func GetCheckTaskMetadataTypeEnum() CheckTaskMetadataTypeEnum {
	return CheckTaskMetadataTypeEnum{
		ADDON_STATIC: CheckTaskMetadataType{
			value: "addonStatic",
		},
		ADDON_UPGRADE: CheckTaskMetadataType{
			value: "addonUpgrade",
		},
	}
}

func (c CheckTaskMetadataType) Value() string {
	return c.value
}

func (c CheckTaskMetadataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckTaskMetadataType) UnmarshalJSON(b []byte) error {
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
