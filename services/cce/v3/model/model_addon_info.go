package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AddonInfo struct {

	// **参数解释：** 插件模板名称 **约束限制：** 不涉及 **取值范围：** cce服务支持的插件模板 **默认取值：** 不涉及
	AddonTemplateName string `json:"addonTemplateName"`

	// **参数解释：** 插件实例ID，可以通过[获取AddonInstance列表](cce_02_0326.xml)中的items[].metadata.uid字段获取 **约束限制：** 此参数必填 **取值范围：** 不涉及 **默认取值：** 不涉及
	AddonInstanceID string `json:"addonInstanceID"`

	// **参数解释：** 插件升级的目标版本 **约束限制：** 插件升级场景下，此参数必填。 **取值范围：** cce服务提供的插件版本，可以通过[查询AddonTemplates列表](cce_02_0321.xml)中的items[].spec.versions.version字段获取 **默认取值：** 不涉及
	TargetVersion *string `json:"targetVersion,omitempty"`

	// **参数解释：** 插件检查类型 **约束限制：** 此参数必填。 **取值范围：** - addonStatic: 运行中插件巡检 - addonUpgrade: 插件升级前检查  **默认取值：** 不涉及
	Type AddonInfoType `json:"type"`

	// **参数解释：** 插件模板编辑/升级参数（各插件不同），请根据具体插件模板信息填写参数 **约束限制：** 不涉及
	Values map[string]interface{} `json:"values,omitempty"`
}

func (o AddonInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonInfo struct{}"
	}

	return strings.Join([]string{"AddonInfo", string(data)}, " ")
}

type AddonInfoType struct {
	value string
}

type AddonInfoTypeEnum struct {
	ADDON_STATIC  AddonInfoType
	ADDON_UPGRADE AddonInfoType
}

func GetAddonInfoTypeEnum() AddonInfoTypeEnum {
	return AddonInfoTypeEnum{
		ADDON_STATIC: AddonInfoType{
			value: "addonStatic",
		},
		ADDON_UPGRADE: AddonInfoType{
			value: "addonUpgrade",
		},
	}
}

func (c AddonInfoType) Value() string {
	return c.value
}

func (c AddonInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddonInfoType) UnmarshalJSON(b []byte) error {
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
