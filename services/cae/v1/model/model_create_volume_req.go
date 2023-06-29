package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateVolumeReq struct {

	// API版本。
	ApiVersion string `json:"api_version"`

	// 资源种类。
	Kind CreateVolumeReqKind `json:"kind"`

	Spec *VolumeSpec `json:"spec"`
}

func (o CreateVolumeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeReq struct{}"
	}

	return strings.Join([]string{"CreateVolumeReq", string(data)}, " ")
}

type CreateVolumeReqKind struct {
	value string
}

type CreateVolumeReqKindEnum struct {
	VOLUME CreateVolumeReqKind
}

func GetCreateVolumeReqKindEnum() CreateVolumeReqKindEnum {
	return CreateVolumeReqKindEnum{
		VOLUME: CreateVolumeReqKind{
			value: "Volume",
		},
	}
}

func (c CreateVolumeReqKind) Value() string {
	return c.value
}

func (c CreateVolumeReqKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVolumeReqKind) UnmarshalJSON(b []byte) error {
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
