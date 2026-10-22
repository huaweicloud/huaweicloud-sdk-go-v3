package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RegisterV2ImageRequestInput struct {

	// **参数解释**：自定义镜像的URL地址，由镜像仓库地址、镜像命名空间、镜像名称和镜像版本组成。 **约束限制**：不涉及。 **取值范围**：由大小写英文字母、数字、小数点(.)、斜杠(/)、冒号(:)、下划线(_)和连字符(-)组成的容器镜像地址。长度为1~2000个字符。 **默认取值**：不涉及。
	ImageUrl string `json:"image_url"`

	Category *V2ImageCategory `json:"category"`

	// **参数解释**：镜像支持的架构列表，取值包括：x86_64（x86架构）、aarch64（ARM架构），未传入该参数时默认为x86_64。 **约束限制**：不涉及。
	Architectures *[]RegisterV2ImageRequestInputArchitectures `json:"architectures,omitempty"`

	Type *V2ImageVersionType `json:"type"`

	// **参数解释**：自定义镜像包含的软件列表，每个元素为一个ContainerSoftware对象，包含软件名称和版本信息。 **约束限制**：软件列表最小个数为0，最大个数为100。
	ContainerSoftwares *[]ContainerSoftware `json:"container_softwares,omitempty"`

	// **参数解释**：自定义镜像的描述。 **约束限制**：不涉及。 **取值范围**：不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号、艾特符、井号、百分号、分号、左括号、右括号、竖线、星号、反斜杠。长度为0~2000个字符。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：基础镜像ID，可通过[查询自定义镜像列表](ListV2TenantImages.xml)获取镜像ID。 **约束限制**：Spark类型镜像必填。 **取值范围**：支持英文字母、数字及中划线，长度0~64个字符。 **默认取值**：不涉及。
	BaseImageId *string `json:"base_image_id,omitempty"`

	// **参数解释**：基础镜像版本ID，可通过[查询自定义镜像版本列表](ListV2TenantImageVersions.xml)获取镜像版本ID。 **约束限制**：当镜像分类为Spark类型时此字段必填，其余情况可选。 **取值范围**：支持英文字母、数字及中划线，长度0~64个字符。 **默认取值**：不涉及。
	BaseImageVersionId *string `json:"base_image_version_id,omitempty"`
}

func (o RegisterV2ImageRequestInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterV2ImageRequestInput struct{}"
	}

	return strings.Join([]string{"RegisterV2ImageRequestInput", string(data)}, " ")
}

type RegisterV2ImageRequestInputArchitectures struct {
	value string
}

type RegisterV2ImageRequestInputArchitecturesEnum struct {
	X86_64  RegisterV2ImageRequestInputArchitectures
	AARCH64 RegisterV2ImageRequestInputArchitectures
}

func GetRegisterV2ImageRequestInputArchitecturesEnum() RegisterV2ImageRequestInputArchitecturesEnum {
	return RegisterV2ImageRequestInputArchitecturesEnum{
		X86_64: RegisterV2ImageRequestInputArchitectures{
			value: "x86_64",
		},
		AARCH64: RegisterV2ImageRequestInputArchitectures{
			value: "aarch64",
		},
	}
}

func (c RegisterV2ImageRequestInputArchitectures) Value() string {
	return c.value
}

func (c RegisterV2ImageRequestInputArchitectures) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterV2ImageRequestInputArchitectures) UnmarshalJSON(b []byte) error {
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
