package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ListVolumesResponse struct {

	// API版本。
	ApiVersion *string `json:"api_version,omitempty"`

	// 资源种类。
	Kind *ListVolumesResponseKind `json:"kind,omitempty"`

	// 数据卷列表。
	Items          *[]Volume `json:"items,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesResponse struct{}"
	}

	return strings.Join([]string{"ListVolumesResponse", string(data)}, " ")
}

type ListVolumesResponseKind struct {
	value string
}

type ListVolumesResponseKindEnum struct {
	VOLUME ListVolumesResponseKind
}

func GetListVolumesResponseKindEnum() ListVolumesResponseKindEnum {
	return ListVolumesResponseKindEnum{
		VOLUME: ListVolumesResponseKind{
			value: "Volume",
		},
	}
}

func (c ListVolumesResponseKind) Value() string {
	return c.value
}

func (c ListVolumesResponseKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVolumesResponseKind) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
