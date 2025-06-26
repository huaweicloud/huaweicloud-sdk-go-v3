package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Accessory The accessory of the artifact
type Accessory struct {

	// 附件ID
	Id *int64 `json:"id,omitempty"`

	// 附件制品ID
	ArtifactId *int64 `json:"artifact_id,omitempty"`

	// 附件所属制品ID.
	SubjectArtifactId *int64 `json:"subject_artifact_id,omitempty"`

	// 附件的大小
	Size *int64 `json:"size,omitempty"`

	// 附件的sha256值
	Digest *string `json:"digest,omitempty"`

	// 附件的类型
	Type *AccessoryType `json:"type,omitempty"`

	// 附件的创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
}

func (o Accessory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Accessory struct{}"
	}

	return strings.Join([]string{"Accessory", string(data)}, " ")
}

type AccessoryType struct {
	value string
}

type AccessoryTypeEnum struct {
	SIGNATURE_COSIGN AccessoryType
}

func GetAccessoryTypeEnum() AccessoryTypeEnum {
	return AccessoryTypeEnum{
		SIGNATURE_COSIGN: AccessoryType{
			value: "signature.cosign",
		},
	}
}

func (c AccessoryType) Value() string {
	return c.value
}

func (c AccessoryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessoryType) UnmarshalJSON(b []byte) error {
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
