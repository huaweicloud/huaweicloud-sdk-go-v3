package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListComponentSnapshotsResponse Response Object
type ListComponentSnapshotsResponse struct {

	// API版本。
	ApiVersion *ListComponentSnapshotsResponseApiVersion `json:"api_version,omitempty"`

	// 资源种类。
	Kind *ListComponentSnapshotsResponseKind `json:"kind,omitempty"`

	// 快照列表。
	Items          *[]ComponentSnapshotItem `json:"items,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListComponentSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"ListComponentSnapshotsResponse", string(data)}, " ")
}

type ListComponentSnapshotsResponseApiVersion struct {
	value string
}

type ListComponentSnapshotsResponseApiVersionEnum struct {
	V1 ListComponentSnapshotsResponseApiVersion
}

func GetListComponentSnapshotsResponseApiVersionEnum() ListComponentSnapshotsResponseApiVersionEnum {
	return ListComponentSnapshotsResponseApiVersionEnum{
		V1: ListComponentSnapshotsResponseApiVersion{
			value: "v1",
		},
	}
}

func (c ListComponentSnapshotsResponseApiVersion) Value() string {
	return c.value
}

func (c ListComponentSnapshotsResponseApiVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListComponentSnapshotsResponseApiVersion) UnmarshalJSON(b []byte) error {
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

type ListComponentSnapshotsResponseKind struct {
	value string
}

type ListComponentSnapshotsResponseKindEnum struct {
	COMPONENT_SNAPSHOT ListComponentSnapshotsResponseKind
}

func GetListComponentSnapshotsResponseKindEnum() ListComponentSnapshotsResponseKindEnum {
	return ListComponentSnapshotsResponseKindEnum{
		COMPONENT_SNAPSHOT: ListComponentSnapshotsResponseKind{
			value: "ComponentSnapshot",
		},
	}
}

func (c ListComponentSnapshotsResponseKind) Value() string {
	return c.value
}

func (c ListComponentSnapshotsResponseKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListComponentSnapshotsResponseKind) UnmarshalJSON(b []byte) error {
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
