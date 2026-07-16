package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NotebookFlavor struct {

	// **参数解释**：架构类型。 **取值范围**：枚举类型，取值如下： - x86_64 - aarch64
	Arch *NotebookFlavorArch `json:"arch,omitempty"`

	Ascend *AscendInfo `json:"ascend,omitempty"`

	Billing *BillingInfo `json:"billing,omitempty"`

	// **参数解释**：规格处理器类型。 **取值范围**：枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)
	Category *NotebookFlavorCategory `json:"category,omitempty"`

	// **参数解释**：规格描述信息。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：实例类别。 **取值范围**：枚举类型，取值如下： - DEFAULT：CodeLab免费规格实例，每个用户最多只能创建一个。 - NOTEBOOK：计费规格实例。
	Feature *NotebookFlavorFeature `json:"feature,omitempty"`

	// **参数解释**：是否为免费规格。 **取值范围**：布尔类型： - true：免费规格。 - false：不是免费规格。
	Free *bool `json:"free,omitempty"`

	Gpu *GpuInfo `json:"gpu,omitempty"`

	// **参数解释**：规格ID。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：内存大小。 **取值范围**：不涉及。
	Memory *int64 `json:"memory,omitempty"`

	// **参数解释**：规格名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：资源是否充足。 **取值范围**：布尔类型： - true 资源不足 - false 资源充足
	SoldOut *bool `json:"sold_out,omitempty"`

	// **参数解释**：规格支持的存储类型。枚举类型，取值如下： - EFS - EVS
	Storages *[]NotebookFlavorStorages `json:"storages,omitempty"`

	// **参数解释**：CPU核数。 **取值范围**：不涉及。
	Vcpus *int32 `json:"vcpus,omitempty"`

	// **参数解释**：规格包含EVS时，EVS存储创建的最大上限(单位：GB)。 **取值范围**：不涉及。
	EvsMaxSize *string `json:"evs_max_size,omitempty"`

	// **参数解释**：规格包含EVS时，EVS存储的sku编码。 **取值范围**：不涉及。
	EvsSkuCode *string `json:"evs_sku_code,omitempty"`

	// **参数解释**：支持站点类型。 **取值范围**：枚举类型，取值如下： - COMMON：国内与国际站都支持。 - NATIONAL：仅在国内站支持。 - INTERNATIONAL：仅在国际站支持。 - NONE：国内与国际站都不支持。
	GrowSupportType *string `json:"grow_support_type,omitempty"`
}

func (o NotebookFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotebookFlavor struct{}"
	}

	return strings.Join([]string{"NotebookFlavor", string(data)}, " ")
}

type NotebookFlavorArch struct {
	value string
}

type NotebookFlavorArchEnum struct {
	X86_64  NotebookFlavorArch
	AARCH64 NotebookFlavorArch
}

func GetNotebookFlavorArchEnum() NotebookFlavorArchEnum {
	return NotebookFlavorArchEnum{
		X86_64: NotebookFlavorArch{
			value: "x86_64",
		},
		AARCH64: NotebookFlavorArch{
			value: "aarch64",
		},
	}
}

func (c NotebookFlavorArch) Value() string {
	return c.value
}

func (c NotebookFlavorArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotebookFlavorArch) UnmarshalJSON(b []byte) error {
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

type NotebookFlavorCategory struct {
	value string
}

type NotebookFlavorCategoryEnum struct {
	ASCEND NotebookFlavorCategory
	CPU    NotebookFlavorCategory
	GPU    NotebookFlavorCategory
}

func GetNotebookFlavorCategoryEnum() NotebookFlavorCategoryEnum {
	return NotebookFlavorCategoryEnum{
		ASCEND: NotebookFlavorCategory{
			value: "ASCEND",
		},
		CPU: NotebookFlavorCategory{
			value: "CPU",
		},
		GPU: NotebookFlavorCategory{
			value: "GPU",
		},
	}
}

func (c NotebookFlavorCategory) Value() string {
	return c.value
}

func (c NotebookFlavorCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotebookFlavorCategory) UnmarshalJSON(b []byte) error {
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

type NotebookFlavorFeature struct {
	value string
}

type NotebookFlavorFeatureEnum struct {
	DEFAULT  NotebookFlavorFeature
	NOTEBOOK NotebookFlavorFeature
}

func GetNotebookFlavorFeatureEnum() NotebookFlavorFeatureEnum {
	return NotebookFlavorFeatureEnum{
		DEFAULT: NotebookFlavorFeature{
			value: "DEFAULT",
		},
		NOTEBOOK: NotebookFlavorFeature{
			value: "NOTEBOOK",
		},
	}
}

func (c NotebookFlavorFeature) Value() string {
	return c.value
}

func (c NotebookFlavorFeature) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotebookFlavorFeature) UnmarshalJSON(b []byte) error {
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

type NotebookFlavorStorages struct {
	value string
}

type NotebookFlavorStoragesEnum struct {
	EFS           NotebookFlavorStorages
	EFS_DEDICATED NotebookFlavorStorages
	EVS           NotebookFlavorStorages
	OBS           NotebookFlavorStorages
	OBSFS         NotebookFlavorStorages
}

func GetNotebookFlavorStoragesEnum() NotebookFlavorStoragesEnum {
	return NotebookFlavorStoragesEnum{
		EFS: NotebookFlavorStorages{
			value: "EFS",
		},
		EFS_DEDICATED: NotebookFlavorStorages{
			value: "EFS_DEDICATED",
		},
		EVS: NotebookFlavorStorages{
			value: "EVS",
		},
		OBS: NotebookFlavorStorages{
			value: "OBS",
		},
		OBSFS: NotebookFlavorStorages{
			value: "OBSFS",
		},
	}
}

func (c NotebookFlavorStorages) Value() string {
	return c.value
}

func (c NotebookFlavorStorages) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotebookFlavorStorages) UnmarshalJSON(b []byte) error {
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
