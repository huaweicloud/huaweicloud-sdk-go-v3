package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NotebookCustomSpecRep **参数描述**：CPU、GPU专属池下，用户指定自定义规格的响应体信息
type NotebookCustomSpecRep struct {

	// **参数描述**：实例申请的GPU卡数。 **取值范围**：不涉及。
	Gpu *float32 `json:"gpu,omitempty"`

	// **参数描述**：实例申请的GPU加速卡类型。 **取值范围**：不涉及。
	GpuType *string `json:"gpu_type,omitempty"`

	// **参数描述**：实例申请的CPU核数大小。 **取值范围**：整数部分最多10位，小数部分最多2位，且数值不得小于0.4。
	Cpu float32 `json:"cpu"`

	// **参数描述**：实例申请的内存大小。 **取值范围**：必须是整数，整数部分最多10位，且数值不得小于513。
	Memory float32 `json:"memory"`

	// **参数描述**：实例申请的CPU架构。 **取值范围**：枚举类型，取值如下：  - X86_64：x86架构 - AARCH64：ARM架构
	Arch NotebookCustomSpecRepArch `json:"arch"`

	// **参数描述**：实例申请的规格类型。 **取值范围**：枚举类型，取值如下：  - CPU：CPU规格。 - GPU：GPU规格。
	Category NotebookCustomSpecRepCategory `json:"category"`

	// **参数解释**：实例选择的目标资源池节点实例规格。 **取值范围**：不涉及。
	ResourceFlavor *string `json:"resource_flavor,omitempty"`
}

func (o NotebookCustomSpecRep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotebookCustomSpecRep struct{}"
	}

	return strings.Join([]string{"NotebookCustomSpecRep", string(data)}, " ")
}

type NotebookCustomSpecRepArch struct {
	value string
}

type NotebookCustomSpecRepArchEnum struct {
	AARCH64 NotebookCustomSpecRepArch
	X86_64  NotebookCustomSpecRepArch
}

func GetNotebookCustomSpecRepArchEnum() NotebookCustomSpecRepArchEnum {
	return NotebookCustomSpecRepArchEnum{
		AARCH64: NotebookCustomSpecRepArch{
			value: "AARCH64",
		},
		X86_64: NotebookCustomSpecRepArch{
			value: "X86_64",
		},
	}
}

func (c NotebookCustomSpecRepArch) Value() string {
	return c.value
}

func (c NotebookCustomSpecRepArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotebookCustomSpecRepArch) UnmarshalJSON(b []byte) error {
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

type NotebookCustomSpecRepCategory struct {
	value string
}

type NotebookCustomSpecRepCategoryEnum struct {
	CPU NotebookCustomSpecRepCategory
	GPU NotebookCustomSpecRepCategory
}

func GetNotebookCustomSpecRepCategoryEnum() NotebookCustomSpecRepCategoryEnum {
	return NotebookCustomSpecRepCategoryEnum{
		CPU: NotebookCustomSpecRepCategory{
			value: "CPU",
		},
		GPU: NotebookCustomSpecRepCategory{
			value: "GPU",
		},
	}
}

func (c NotebookCustomSpecRepCategory) Value() string {
	return c.value
}

func (c NotebookCustomSpecRepCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotebookCustomSpecRepCategory) UnmarshalJSON(b []byte) error {
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
