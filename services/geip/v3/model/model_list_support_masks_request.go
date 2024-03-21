package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSupportMasksRequest Request Object
type ListSupportMasksRequest struct {

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *int32 `json:"offset,omitempty"`

	// 分页起始点
	Marker *string `json:"marker,omitempty"`

	// 翻页方向
	PageReverse *bool `json:"page_reverse,omitempty"`

	Fields *[]ListSupportMasksRequestFields `json:"fields,omitempty"`

	// 按照sort_key指定的字段排序
	SortKey *[]ListSupportMasksRequestSortKey `json:"sort_key,omitempty"`

	// 排序的方向，倒序或者正序
	SortDir *[]ListSupportMasksRequestSortDir `json:"sort_dir,omitempty"`

	Id *[]string `json:"id,omitempty"`

	IpVersion *[]ListSupportMasksRequestIpVersion `json:"ip_version,omitempty"`

	Mask *[]int32 `json:"mask,omitempty"`
}

func (o ListSupportMasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportMasksRequest struct{}"
	}

	return strings.Join([]string{"ListSupportMasksRequest", string(data)}, " ")
}

type ListSupportMasksRequestFields struct {
	value string
}

type ListSupportMasksRequestFieldsEnum struct {
	ID         ListSupportMasksRequestFields
	IP_VERSION ListSupportMasksRequestFields
	MASK       ListSupportMasksRequestFields
	CREATED_AT ListSupportMasksRequestFields
	UPDATED_AT ListSupportMasksRequestFields
}

func GetListSupportMasksRequestFieldsEnum() ListSupportMasksRequestFieldsEnum {
	return ListSupportMasksRequestFieldsEnum{
		ID: ListSupportMasksRequestFields{
			value: "id",
		},
		IP_VERSION: ListSupportMasksRequestFields{
			value: "ip_version",
		},
		MASK: ListSupportMasksRequestFields{
			value: "mask",
		},
		CREATED_AT: ListSupportMasksRequestFields{
			value: "created_at",
		},
		UPDATED_AT: ListSupportMasksRequestFields{
			value: "updated_at",
		},
	}
}

func (c ListSupportMasksRequestFields) Value() string {
	return c.value
}

func (c ListSupportMasksRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportMasksRequestFields) UnmarshalJSON(b []byte) error {
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

type ListSupportMasksRequestSortKey struct {
	value string
}

type ListSupportMasksRequestSortKeyEnum struct {
	ID         ListSupportMasksRequestSortKey
	IP_VERSION ListSupportMasksRequestSortKey
	MASK       ListSupportMasksRequestSortKey
	CREATED_AT ListSupportMasksRequestSortKey
	UPDATED_AT ListSupportMasksRequestSortKey
}

func GetListSupportMasksRequestSortKeyEnum() ListSupportMasksRequestSortKeyEnum {
	return ListSupportMasksRequestSortKeyEnum{
		ID: ListSupportMasksRequestSortKey{
			value: "id",
		},
		IP_VERSION: ListSupportMasksRequestSortKey{
			value: "ip_version",
		},
		MASK: ListSupportMasksRequestSortKey{
			value: "mask",
		},
		CREATED_AT: ListSupportMasksRequestSortKey{
			value: "created_at",
		},
		UPDATED_AT: ListSupportMasksRequestSortKey{
			value: "updated_at",
		},
	}
}

func (c ListSupportMasksRequestSortKey) Value() string {
	return c.value
}

func (c ListSupportMasksRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportMasksRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListSupportMasksRequestSortDir struct {
	value string
}

type ListSupportMasksRequestSortDirEnum struct {
	ASC  ListSupportMasksRequestSortDir
	DESC ListSupportMasksRequestSortDir
}

func GetListSupportMasksRequestSortDirEnum() ListSupportMasksRequestSortDirEnum {
	return ListSupportMasksRequestSortDirEnum{
		ASC: ListSupportMasksRequestSortDir{
			value: "asc",
		},
		DESC: ListSupportMasksRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListSupportMasksRequestSortDir) Value() string {
	return c.value
}

func (c ListSupportMasksRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportMasksRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListSupportMasksRequestIpVersion struct {
	value int32
}

type ListSupportMasksRequestIpVersionEnum struct {
	E_4 ListSupportMasksRequestIpVersion
	E_6 ListSupportMasksRequestIpVersion
}

func GetListSupportMasksRequestIpVersionEnum() ListSupportMasksRequestIpVersionEnum {
	return ListSupportMasksRequestIpVersionEnum{
		E_4: ListSupportMasksRequestIpVersion{
			value: 4,
		}, E_6: ListSupportMasksRequestIpVersion{
			value: 6,
		},
	}
}

func (c ListSupportMasksRequestIpVersion) Value() int32 {
	return c.value
}

func (c ListSupportMasksRequestIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportMasksRequestIpVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
