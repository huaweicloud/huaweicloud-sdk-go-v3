package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAccessSitesRequest Request Object
type ListAccessSitesRequest struct {

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *int32 `json:"offset,omitempty"`

	// 分页起始点
	Marker *string `json:"marker,omitempty"`

	// 翻页方向
	PageReverse *bool `json:"page_reverse,omitempty"`

	Fields *[]ListAccessSitesRequestFields `json:"fields,omitempty"`

	// 按照sort_key指定的字段排序
	SortKey *[]ListAccessSitesRequestSortKey `json:"sort_key,omitempty"`

	// 排序的方向，倒序或者正序
	SortDir *[]ListAccessSitesRequestSortDir `json:"sort_dir,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Code *[]string `json:"code,omitempty"`

	ProxyRegion *[]string `json:"proxy_region,omitempty"`

	IecAzCode *[]string `json:"iec_az_code,omitempty"`
}

func (o ListAccessSitesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessSitesRequest struct{}"
	}

	return strings.Join([]string{"ListAccessSitesRequest", string(data)}, " ")
}

type ListAccessSitesRequestFields struct {
	value string
}

type ListAccessSitesRequestFieldsEnum struct {
	ID           ListAccessSitesRequestFields
	NAME         ListAccessSitesRequestFields
	PROXY_REGION ListAccessSitesRequestFields
	EN_NAME      ListAccessSitesRequestFields
	CN_NAME      ListAccessSitesRequestFields
	CREATED_AT   ListAccessSitesRequestFields
	UPDATED_AT   ListAccessSitesRequestFields
}

func GetListAccessSitesRequestFieldsEnum() ListAccessSitesRequestFieldsEnum {
	return ListAccessSitesRequestFieldsEnum{
		ID: ListAccessSitesRequestFields{
			value: "id",
		},
		NAME: ListAccessSitesRequestFields{
			value: "name",
		},
		PROXY_REGION: ListAccessSitesRequestFields{
			value: "proxy_region",
		},
		EN_NAME: ListAccessSitesRequestFields{
			value: "en_name",
		},
		CN_NAME: ListAccessSitesRequestFields{
			value: "cn_name",
		},
		CREATED_AT: ListAccessSitesRequestFields{
			value: "created_at",
		},
		UPDATED_AT: ListAccessSitesRequestFields{
			value: "updated_at",
		},
	}
}

func (c ListAccessSitesRequestFields) Value() string {
	return c.value
}

func (c ListAccessSitesRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAccessSitesRequestFields) UnmarshalJSON(b []byte) error {
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

type ListAccessSitesRequestSortKey struct {
	value string
}

type ListAccessSitesRequestSortKeyEnum struct {
	ID         ListAccessSitesRequestSortKey
	NAME       ListAccessSitesRequestSortKey
	CREATED_AT ListAccessSitesRequestSortKey
	UPDATED_AT ListAccessSitesRequestSortKey
}

func GetListAccessSitesRequestSortKeyEnum() ListAccessSitesRequestSortKeyEnum {
	return ListAccessSitesRequestSortKeyEnum{
		ID: ListAccessSitesRequestSortKey{
			value: "id",
		},
		NAME: ListAccessSitesRequestSortKey{
			value: "name",
		},
		CREATED_AT: ListAccessSitesRequestSortKey{
			value: "created_at",
		},
		UPDATED_AT: ListAccessSitesRequestSortKey{
			value: "updated_at",
		},
	}
}

func (c ListAccessSitesRequestSortKey) Value() string {
	return c.value
}

func (c ListAccessSitesRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAccessSitesRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListAccessSitesRequestSortDir struct {
	value string
}

type ListAccessSitesRequestSortDirEnum struct {
	ASC  ListAccessSitesRequestSortDir
	DESC ListAccessSitesRequestSortDir
}

func GetListAccessSitesRequestSortDirEnum() ListAccessSitesRequestSortDirEnum {
	return ListAccessSitesRequestSortDirEnum{
		ASC: ListAccessSitesRequestSortDir{
			value: "asc",
		},
		DESC: ListAccessSitesRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListAccessSitesRequestSortDir) Value() string {
	return c.value
}

func (c ListAccessSitesRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAccessSitesRequestSortDir) UnmarshalJSON(b []byte) error {
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
