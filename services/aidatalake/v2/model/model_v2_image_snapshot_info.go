package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// V2ImageSnapshotInfo **参数解释**：镜像快照信息。
type V2ImageSnapshotInfo struct {

	// **参数解释**：镜像快照记录ID。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：镜像ID。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**：镜像版本ID。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	ImageVersionId *string `json:"image_version_id,omitempty"`

	// **参数解释**：镜像快照名称（用户指定）。 **约束限制**：不涉及。 **取值范围**：以小写字母开头，小写字母、数字、中划线（-）组成，小写字母或数字结尾的1到190字符的字符串。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：镜像完整地址（含registry/namespace/name:version）。 **约束限制**：不涉及。 **取值范围**：支持英文字母、数字、小数点、冒号、斜杠及连字符，长度为0~1024个字符。 **默认取值**：不涉及。
	ImageAddress *string `json:"image_address,omitempty"`

	// **参数解释**：镜像名称。 **约束限制**：不涉及。 **取值范围**：支持英文字母、数字、小数点、斜杠、冒号、下划线和连字符，长度为0~128个字符。 **默认取值**：不涉及。
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**：镜像版本。 **约束限制**：不涉及。 **取值范围**：支持英文字母、数字、小数点、冒号、下划线和连字符，长度为0~1024个字符。 **默认取值**：不涉及。
	ImageVersion *string `json:"image_version,omitempty"`

	Status *ImageSnapshotStatus `json:"status,omitempty"`

	// **参数解释**：快照大小，单位GiB。 **约束限制**：不涉及。 **取值范围**：20~400。 **默认取值**：不涉及。
	ImageSnapshotSize *int32 `json:"image_snapshot_size,omitempty"`

	// **参数解释**：创建方式。 **约束限制**：不涉及。 **取值范围**：   - Created：直接创建。   - Associated：关联CCI已有快照。 **默认取值**：不涉及。
	CreationMode *V2ImageSnapshotInfoCreationMode `json:"creation_mode,omitempty"`

	// **参数解释**：镜像快照保留天数。超过保留天数自动失效，0表示永久保留。 **约束限制**：不涉及。 **取值范围**：0~65536。 **默认取值**：不涉及。
	ReservedDays *int32 `json:"reserved_days,omitempty"`

	// **参数解释**：超期时间，null表示永久有效。使用UTC时间，ISO 8601格式，例如：2025-01-10T03:37:12.000+00:00。 **取值范围**：不涉及。
	ExpireTime *sdktime.SdkTime `json:"expire_time,omitempty"`

	// **参数解释**：镜像快照失效的提示信息。 **约束限制**：不涉及。 **取值范围**：支持中文字符、英文字母、数字及常用符号，长度为0~65535个字符。 **默认取值**：不涉及。
	FailedMessage *string `json:"failed_message,omitempty"`

	// **参数解释**：资源创建时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：资源更新时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o V2ImageSnapshotInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2ImageSnapshotInfo struct{}"
	}

	return strings.Join([]string{"V2ImageSnapshotInfo", string(data)}, " ")
}

type V2ImageSnapshotInfoCreationMode struct {
	value string
}

type V2ImageSnapshotInfoCreationModeEnum struct {
	CREATED    V2ImageSnapshotInfoCreationMode
	ASSOCIATED V2ImageSnapshotInfoCreationMode
}

func GetV2ImageSnapshotInfoCreationModeEnum() V2ImageSnapshotInfoCreationModeEnum {
	return V2ImageSnapshotInfoCreationModeEnum{
		CREATED: V2ImageSnapshotInfoCreationMode{
			value: "Created",
		},
		ASSOCIATED: V2ImageSnapshotInfoCreationMode{
			value: "Associated",
		},
	}
}

func (c V2ImageSnapshotInfoCreationMode) Value() string {
	return c.value
}

func (c V2ImageSnapshotInfoCreationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *V2ImageSnapshotInfoCreationMode) UnmarshalJSON(b []byte) error {
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
