package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type V2ImageVersionInfo struct {

	// **参数解释**：镜像版本ID。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。
	VersionId *string `json:"version_id,omitempty"`

	// **参数解释**：镜像版本。 **取值范围**：支持英文字母、数字、小数点、斜杠、冒号、下划线和连字符，长度2~32个字符。
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**：镜像分类。 **取值范围**：   - RAY_BASIC：基于Ray构建的自定义镜像。   - SPARK_JOB：基于SparkJob构建的自定义镜像。   - SPARK_SQL：基于SparkSql构建的自定义镜像。   - FLINK_JOB：基于FlinkJob构建的自定义镜像。   - AURA_JOB：基于AuraJobV2构建的自定义镜像。
	Category *string `json:"category,omitempty"`

	// **参数解释**：镜像域名。 **取值范围**：支持英文字母、数字、小数点及连字符，长度0~255个字符。
	ImageRegistry *string `json:"image_registry,omitempty"`

	// **参数解释**：镜像命名空间。 **取值范围**：由英文字母、数字、中划线和下划线组成，长度0~64个字符。
	ImageNamespace *string `json:"image_namespace,omitempty"`

	// **参数解释**：镜像名。 **取值范围**：由中文字符、英文字母、数字、中划线和下划线组成，长度0~128个字符。
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**：镜像支持的架构列表，取值包括：x86_64（x86架构）、aarch64（ARM架构）。
	Architectures *[]V2ImageVersionInfoArchitectures `json:"architectures,omitempty"`

	// **参数解释**：镜像软件列表，每个元素为一个ContainerSoftware对象，包含软件名称和版本信息。 **约束限制**：最小个数为0，最大个数为100。
	ContainerSoftwares *[]ContainerSoftware `json:"container_softwares,omitempty"`

	Status *V2ImageVersionStatus `json:"status,omitempty"`

	Type *V2ImageVersionType `json:"type,omitempty"`

	// **参数解释**：镜像版本可见性。 **取值范围**：   - PRIVATE：私有，仅工作空间内可见。   - PUBLIC：公共，所有工作空间可见。
	Visibility *string `json:"visibility,omitempty"`

	// **参数解释**：资源创建时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：资源更新时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// **参数解释**：镜像版本描述。 **取值范围**：支持中文字符、英文字母、数字及常用符号（不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号），长度0~2000个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**：关联的预置基础镜像ID。 **约束限制**：Spark类型自定义镜像时必填。 **取值范围**：支持英文字母、数字及中划线，长度0~64个字符。 **默认取值**：不涉及。
	BaseImageId *string `json:"base_image_id,omitempty"`

	// **参数解释**：关联的预置基础镜像版本ID。 **约束限制**：Spark类型自定义镜像时必填。 **取值范围**：支持英文字母、数字及中划线，长度0~64个字符。 **默认取值**：不涉及。
	BaseImageVersionId *string `json:"base_image_version_id,omitempty"`

	ImageSnapshot *V2ImageSnapshotInfo `json:"image_snapshot,omitempty"`
}

func (o V2ImageVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2ImageVersionInfo struct{}"
	}

	return strings.Join([]string{"V2ImageVersionInfo", string(data)}, " ")
}

type V2ImageVersionInfoArchitectures struct {
	value string
}

type V2ImageVersionInfoArchitecturesEnum struct {
	X86_64  V2ImageVersionInfoArchitectures
	AARCH64 V2ImageVersionInfoArchitectures
}

func GetV2ImageVersionInfoArchitecturesEnum() V2ImageVersionInfoArchitecturesEnum {
	return V2ImageVersionInfoArchitecturesEnum{
		X86_64: V2ImageVersionInfoArchitectures{
			value: "x86_64",
		},
		AARCH64: V2ImageVersionInfoArchitectures{
			value: "aarch64",
		},
	}
}

func (c V2ImageVersionInfoArchitectures) Value() string {
	return c.value
}

func (c V2ImageVersionInfoArchitectures) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *V2ImageVersionInfoArchitectures) UnmarshalJSON(b []byte) error {
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
