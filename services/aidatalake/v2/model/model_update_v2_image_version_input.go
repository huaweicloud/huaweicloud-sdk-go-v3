package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateV2ImageVersionInput **参数解释**：更新自定义镜像请求体。
type UpdateV2ImageVersionInput struct {

	// **参数解释**：自定义镜像的描述。 **约束限制**：不涉及。 **取值范围**：不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号、艾特符、井号、百分号、分号、左括号、右括号、竖线、星号、反斜杠。长度为0~2000个字符。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：支持镜像架构列表，取值包括：x86_64（x86架构）、aarch64（ARM架构），未传入该参数时默认为x86_64。 **约束限制**：不涉及。
	Architectures *[]UpdateV2ImageVersionInputArchitectures `json:"architectures,omitempty"`

	Type *V2ImageVersionType `json:"type,omitempty"`

	// **参数解释**：自定义镜像包含的软件列表，每个元素为一个ContainerSoftware对象，包含软件名称和版本信息，软件列表最小个数为0，最大个数为100。 **约束限制**：不涉及。
	ContainerSoftwares *[]ContainerSoftware `json:"container_softwares,omitempty"`

	Category *V2ImageCategory `json:"category,omitempty"`
}

func (o UpdateV2ImageVersionInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateV2ImageVersionInput struct{}"
	}

	return strings.Join([]string{"UpdateV2ImageVersionInput", string(data)}, " ")
}

type UpdateV2ImageVersionInputArchitectures struct {
	value string
}

type UpdateV2ImageVersionInputArchitecturesEnum struct {
	X86_64  UpdateV2ImageVersionInputArchitectures
	AARCH64 UpdateV2ImageVersionInputArchitectures
}

func GetUpdateV2ImageVersionInputArchitecturesEnum() UpdateV2ImageVersionInputArchitecturesEnum {
	return UpdateV2ImageVersionInputArchitecturesEnum{
		X86_64: UpdateV2ImageVersionInputArchitectures{
			value: "x86_64",
		},
		AARCH64: UpdateV2ImageVersionInputArchitectures{
			value: "aarch64",
		},
	}
}

func (c UpdateV2ImageVersionInputArchitectures) Value() string {
	return c.value
}

func (c UpdateV2ImageVersionInputArchitectures) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateV2ImageVersionInputArchitectures) UnmarshalJSON(b []byte) error {
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
