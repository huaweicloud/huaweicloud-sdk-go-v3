package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteVolumeResponse Response Object
type DeleteVolumeResponse struct {

	// API版本。
	ApiVersion *DeleteVolumeResponseApiVersion `json:"api_version,omitempty"`

	// 资源种类。
	Kind *string `json:"kind,omitempty"`

	// 挂载组件列表
	Items          *[]MountComponent `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVolumeResponse struct{}"
	}

	return strings.Join([]string{"DeleteVolumeResponse", string(data)}, " ")
}

type DeleteVolumeResponseApiVersion struct {
	value string
}

type DeleteVolumeResponseApiVersionEnum struct {
	COMPONENT DeleteVolumeResponseApiVersion
}

func GetDeleteVolumeResponseApiVersionEnum() DeleteVolumeResponseApiVersionEnum {
	return DeleteVolumeResponseApiVersionEnum{
		COMPONENT: DeleteVolumeResponseApiVersion{
			value: "Component",
		},
	}
}

func (c DeleteVolumeResponseApiVersion) Value() string {
	return c.value
}

func (c DeleteVolumeResponseApiVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteVolumeResponseApiVersion) UnmarshalJSON(b []byte) error {
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
