package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSupportRegionsRequest Request Object
type ListSupportRegionsRequest struct {

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *int32 `json:"offset,omitempty"`

	// 分页起始点
	Marker *string `json:"marker,omitempty"`

	// 翻页方向
	PageReverse *bool `json:"page_reverse,omitempty"`

	Fields *[]ListSupportRegionsRequestFields `json:"fields,omitempty"`

	// 按照sort_key指定的字段排序
	SortKey *[]ListSupportRegionsRequestSortKey `json:"sort_key,omitempty"`

	// 排序的方向，倒序或者正序
	SortDir *[]ListSupportRegionsRequestSortDir `json:"sort_dir,omitempty"`

	Id *[]string `json:"id,omitempty"`

	InstanceType *[]string `json:"instance_type,omitempty"`

	PublicBorderGroup *[]string `json:"public_border_group,omitempty"`

	AccessSite *[]string `json:"access_site,omitempty"`

	RegionId *[]string `json:"region_id,omitempty"`

	RemoteEndpoint *[]string `json:"remote_endpoint,omitempty"`

	Status *[]string `json:"status,omitempty"`
}

func (o ListSupportRegionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportRegionsRequest struct{}"
	}

	return strings.Join([]string{"ListSupportRegionsRequest", string(data)}, " ")
}

type ListSupportRegionsRequestFields struct {
	value string
}

type ListSupportRegionsRequestFieldsEnum struct {
	ID                  ListSupportRegionsRequestFields
	INSTANCE_TYPE       ListSupportRegionsRequestFields
	ACCESS_SITE         ListSupportRegionsRequestFields
	REGION_ID           ListSupportRegionsRequestFields
	PUBLIC_BORDER_GROUP ListSupportRegionsRequestFields
	REMOTE_ENDPOINT     ListSupportRegionsRequestFields
	STATUS              ListSupportRegionsRequestFields
	CREATED_AT          ListSupportRegionsRequestFields
	UPDATED_AT          ListSupportRegionsRequestFields
}

func GetListSupportRegionsRequestFieldsEnum() ListSupportRegionsRequestFieldsEnum {
	return ListSupportRegionsRequestFieldsEnum{
		ID: ListSupportRegionsRequestFields{
			value: "id",
		},
		INSTANCE_TYPE: ListSupportRegionsRequestFields{
			value: "instance_type",
		},
		ACCESS_SITE: ListSupportRegionsRequestFields{
			value: "access_site",
		},
		REGION_ID: ListSupportRegionsRequestFields{
			value: "region_id",
		},
		PUBLIC_BORDER_GROUP: ListSupportRegionsRequestFields{
			value: "public_border_group",
		},
		REMOTE_ENDPOINT: ListSupportRegionsRequestFields{
			value: "remote_endpoint",
		},
		STATUS: ListSupportRegionsRequestFields{
			value: "status",
		},
		CREATED_AT: ListSupportRegionsRequestFields{
			value: "created_at",
		},
		UPDATED_AT: ListSupportRegionsRequestFields{
			value: "updated_at",
		},
	}
}

func (c ListSupportRegionsRequestFields) Value() string {
	return c.value
}

func (c ListSupportRegionsRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportRegionsRequestFields) UnmarshalJSON(b []byte) error {
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

type ListSupportRegionsRequestSortKey struct {
	value string
}

type ListSupportRegionsRequestSortKeyEnum struct {
	ID                  ListSupportRegionsRequestSortKey
	INSTANCE_TYPE       ListSupportRegionsRequestSortKey
	ACCESS_SITE         ListSupportRegionsRequestSortKey
	PUBLIC_BORDER_GROUP ListSupportRegionsRequestSortKey
	REGION_ID           ListSupportRegionsRequestSortKey
	REMOTE_ENDPOINT     ListSupportRegionsRequestSortKey
	STATUS              ListSupportRegionsRequestSortKey
	CREATED_AT          ListSupportRegionsRequestSortKey
	UPDATED_AT          ListSupportRegionsRequestSortKey
}

func GetListSupportRegionsRequestSortKeyEnum() ListSupportRegionsRequestSortKeyEnum {
	return ListSupportRegionsRequestSortKeyEnum{
		ID: ListSupportRegionsRequestSortKey{
			value: "id",
		},
		INSTANCE_TYPE: ListSupportRegionsRequestSortKey{
			value: "instance_type",
		},
		ACCESS_SITE: ListSupportRegionsRequestSortKey{
			value: "access_site",
		},
		PUBLIC_BORDER_GROUP: ListSupportRegionsRequestSortKey{
			value: "public_border_group",
		},
		REGION_ID: ListSupportRegionsRequestSortKey{
			value: "region_id",
		},
		REMOTE_ENDPOINT: ListSupportRegionsRequestSortKey{
			value: "remote_endpoint",
		},
		STATUS: ListSupportRegionsRequestSortKey{
			value: "status",
		},
		CREATED_AT: ListSupportRegionsRequestSortKey{
			value: "created_at",
		},
		UPDATED_AT: ListSupportRegionsRequestSortKey{
			value: "updated_at",
		},
	}
}

func (c ListSupportRegionsRequestSortKey) Value() string {
	return c.value
}

func (c ListSupportRegionsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportRegionsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListSupportRegionsRequestSortDir struct {
	value string
}

type ListSupportRegionsRequestSortDirEnum struct {
	ASC  ListSupportRegionsRequestSortDir
	DESC ListSupportRegionsRequestSortDir
}

func GetListSupportRegionsRequestSortDirEnum() ListSupportRegionsRequestSortDirEnum {
	return ListSupportRegionsRequestSortDirEnum{
		ASC: ListSupportRegionsRequestSortDir{
			value: "asc",
		},
		DESC: ListSupportRegionsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListSupportRegionsRequestSortDir) Value() string {
	return c.value
}

func (c ListSupportRegionsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportRegionsRequestSortDir) UnmarshalJSON(b []byte) error {
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
