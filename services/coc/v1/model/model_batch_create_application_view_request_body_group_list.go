package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchCreateApplicationViewRequestBodyGroupList struct {

	// **参数解释：** 分组名称。 **约束限制：** 不涉及。 **取值范围：** 由中文、英文字母、数字、中划线、下划线组成，长度在3~50个字符之间。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 区域id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释：** 关联的资源id列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	CmdbResourceIdList *[]string `json:"cmdb_resource_id_list,omitempty"`

	// **参数解释：** 父节点名称。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度3到50个字符。 **默认取值：** 不涉及。
	ParentName *string `json:"parent_name,omitempty"`

	// **参数解释：** 同步模式。 **约束限制：** 不涉及。 **取值范围：** - MANUAL：表示手动关联：用户在对应分组下，手动将对应资源数据关联至分组内进行管理。 - AUTO：表示智能关联：用户通过企业项目和标签的形式，将企业项目下的相同标签资源创建至同一资源分组。 **默认取值：** 不涉及。
	SyncMode *BatchCreateApplicationViewRequestBodyGroupListSyncMode `json:"sync_mode,omitempty"`

	// **参数解释：** 智能关联规则。 **约束限制：** 不涉及。 **取值范围：** 智能关联已选择的企业项目和对应标签的现存及未来创建的资源。 **默认取值：** 不涉及。
	SyncRules *[]BatchCreateApplicationViewRequestBodySyncRules `json:"sync_rules,omitempty"`

	// **参数解释：** 分组关联的应用名称。 **约束限制：** 不涉及。 **取值范围：**  由中文、英文字母、数字、中划线、下划线组成，长度在3~50个字符之间。 **默认取值：** 不涉及。
	ApplicationName *string `json:"application_name,omitempty"`

	// **参数解释：** 分组关联的组件名称。 **约束限制：** 不涉及。 **取值范围：**  由中文、英文字母、数字、中划线、下划线组成，长度在3~50个字符之间。 **默认取值：** 不涉及。
	ComponentName *string `json:"component_name,omitempty"`

	// **参数解释：** 云广商信息。 **约束限制：** 不涉及。 **取值范围：** - RMS： 华为云。 - AWS：亚马逊。 - AZURE：微软。 - ALI：阿里云。 - VMWARE：VMware。 - OPENSTACK：openstack云平台。 - HCS：Huawei Cloud Stack。 - OTHER：其他云广商。 **默认取值：** 不涉及。
	Vendor *string `json:"vendor,omitempty"`

	// **参数解释：** 分组配置信息。 **约束限制：** 不涉及。 **取值范围：** 分组的关联配置信息，比如对应的APM的配置信息。 **默认取值：** 不涉及。
	RelationConfigurations *[]GroupRelationConfiguration `json:"relation_configurations,omitempty"`

	// **参数解释：** 关联的租户id。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度32个字符。 **默认取值：** 不涉及。
	RelatedDomainId *string `json:"related_domain_id,omitempty"`
}

func (o BatchCreateApplicationViewRequestBodyGroupList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewRequestBodyGroupList struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewRequestBodyGroupList", string(data)}, " ")
}

type BatchCreateApplicationViewRequestBodyGroupListSyncMode struct {
	value string
}

type BatchCreateApplicationViewRequestBodyGroupListSyncModeEnum struct {
	MANUAL BatchCreateApplicationViewRequestBodyGroupListSyncMode
	AUTO   BatchCreateApplicationViewRequestBodyGroupListSyncMode
}

func GetBatchCreateApplicationViewRequestBodyGroupListSyncModeEnum() BatchCreateApplicationViewRequestBodyGroupListSyncModeEnum {
	return BatchCreateApplicationViewRequestBodyGroupListSyncModeEnum{
		MANUAL: BatchCreateApplicationViewRequestBodyGroupListSyncMode{
			value: "MANUAL",
		},
		AUTO: BatchCreateApplicationViewRequestBodyGroupListSyncMode{
			value: "AUTO",
		},
	}
}

func (c BatchCreateApplicationViewRequestBodyGroupListSyncMode) Value() string {
	return c.value
}

func (c BatchCreateApplicationViewRequestBodyGroupListSyncMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateApplicationViewRequestBodyGroupListSyncMode) UnmarshalJSON(b []byte) error {
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
