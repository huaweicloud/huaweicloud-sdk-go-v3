package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InferFlavor struct {

	// **参数解释：** 架构类型。 **取值范围：** - X86_64 - arm64
	Arch *InferFlavorArch `json:"arch,omitempty"`

	Ascend *AscendResource `json:"ascend,omitempty"`

	Billing *BillingResource `json:"billing,omitempty"`

	// **参数解释：** 规格处理器类型。 **取值范围：** - CPU - GPU - [ASCEND](tag:hws,hws_hk,hk,fcs_super)
	Category *InferFlavorCategory `json:"category,omitempty"`

	// **参数解释：** 规格描述信息。 **取值范围：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 规格类别。 **取值范围：** - DEFAULT：CodeLab规格。 - NOTEBOOK：Notebook规格。
	Feature *InferFlavorFeature `json:"feature,omitempty"`

	// **参数解释：** 是否为免费规格。 **取值范围：** - true: 免费规格。 - false: 付费规格。
	Free *bool `json:"free,omitempty"`

	Gpu *GpuResource `json:"gpu,omitempty"`

	// **参数解释：** 规格ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 内存大小。 **取值范围：** 不涉及。
	Memory *int64 `json:"memory,omitempty"`

	// **参数解释：** 规格名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 资源是否充足。 **取值范围：** - true 资源不足。 - false 资源充足。
	SoldOut *bool `json:"sold_out,omitempty"`

	// **参数解释：** 规格支持的存储类型。
	Storages *[]InferFlavorStorages `json:"storages,omitempty"`

	// **参数解释：** CPU核数。 **取值范围：** 不涉及。
	Vcpus *int32 `json:"vcpus,omitempty"`

	// **参数解释：** EVS磁盘最大容量。 **取值范围：** 不涉及。
	EvsMaxSize *string `json:"evs_max_size,omitempty"`
}

func (o InferFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InferFlavor struct{}"
	}

	return strings.Join([]string{"InferFlavor", string(data)}, " ")
}

type InferFlavorArch struct {
	value string
}

type InferFlavorArchEnum struct {
	ARM64  InferFlavorArch
	X86_64 InferFlavorArch
}

func GetInferFlavorArchEnum() InferFlavorArchEnum {
	return InferFlavorArchEnum{
		ARM64: InferFlavorArch{
			value: "arm64",
		},
		X86_64: InferFlavorArch{
			value: "X86_64",
		},
	}
}

func (c InferFlavorArch) Value() string {
	return c.value
}

func (c InferFlavorArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InferFlavorArch) UnmarshalJSON(b []byte) error {
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

type InferFlavorCategory struct {
	value string
}

type InferFlavorCategoryEnum struct {
	ASCEND InferFlavorCategory
	CPU    InferFlavorCategory
	GPU    InferFlavorCategory
}

func GetInferFlavorCategoryEnum() InferFlavorCategoryEnum {
	return InferFlavorCategoryEnum{
		ASCEND: InferFlavorCategory{
			value: "ASCEND",
		},
		CPU: InferFlavorCategory{
			value: "CPU",
		},
		GPU: InferFlavorCategory{
			value: "GPU",
		},
	}
}

func (c InferFlavorCategory) Value() string {
	return c.value
}

func (c InferFlavorCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InferFlavorCategory) UnmarshalJSON(b []byte) error {
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

type InferFlavorFeature struct {
	value string
}

type InferFlavorFeatureEnum struct {
	DEFAULT  InferFlavorFeature
	NOTEBOOK InferFlavorFeature
}

func GetInferFlavorFeatureEnum() InferFlavorFeatureEnum {
	return InferFlavorFeatureEnum{
		DEFAULT: InferFlavorFeature{
			value: "DEFAULT",
		},
		NOTEBOOK: InferFlavorFeature{
			value: "NOTEBOOK",
		},
	}
}

func (c InferFlavorFeature) Value() string {
	return c.value
}

func (c InferFlavorFeature) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InferFlavorFeature) UnmarshalJSON(b []byte) error {
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

type InferFlavorStorages struct {
	value string
}

type InferFlavorStoragesEnum struct {
	EFS           InferFlavorStorages
	EFS_DEDICATED InferFlavorStorages
	EVS           InferFlavorStorages
	OBS           InferFlavorStorages
	OBSFS         InferFlavorStorages
}

func GetInferFlavorStoragesEnum() InferFlavorStoragesEnum {
	return InferFlavorStoragesEnum{
		EFS: InferFlavorStorages{
			value: "EFS",
		},
		EFS_DEDICATED: InferFlavorStorages{
			value: "EFS_DEDICATED",
		},
		EVS: InferFlavorStorages{
			value: "EVS",
		},
		OBS: InferFlavorStorages{
			value: "OBS",
		},
		OBSFS: InferFlavorStorages{
			value: "OBSFS",
		},
	}
}

func (c InferFlavorStorages) Value() string {
	return c.value
}

func (c InferFlavorStorages) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InferFlavorStorages) UnmarshalJSON(b []byte) error {
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
