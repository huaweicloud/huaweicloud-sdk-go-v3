package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInternetBandwidthsRequest Request Object
type ListInternetBandwidthsRequest struct {

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *int32 `json:"offset,omitempty"`

	// 分页起始点
	Marker *string `json:"marker,omitempty"`

	// 翻页方向
	PageReverse *bool `json:"page_reverse,omitempty"`

	Fields *[]ListInternetBandwidthsRequestFields `json:"fields,omitempty"`

	ExtFields *[]ListInternetBandwidthsRequestExtFields `json:"ext-fields,omitempty"`

	// 按照sort_key指定的字段排序
	SortKey *[]ListInternetBandwidthsRequestSortKey `json:"sort_key,omitempty"`

	// 排序的方向，倒序或者正序
	SortDir *[]ListInternetBandwidthsRequestSortDir `json:"sort_dir,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Size *[]int32 `json:"size,omitempty"`

	Name *[]string `json:"name,omitempty"`

	NameLike *string `json:"name_like,omitempty"`

	AccessSite *[]string `json:"access_site,omitempty"`

	Status *[]ListInternetBandwidthsRequestStatus `json:"status,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	Type *[]string `json:"type,omitempty"`
}

func (o ListInternetBandwidthsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternetBandwidthsRequest struct{}"
	}

	return strings.Join([]string{"ListInternetBandwidthsRequest", string(data)}, " ")
}

type ListInternetBandwidthsRequestFields struct {
	value string
}

type ListInternetBandwidthsRequestFieldsEnum struct {
	ID                    ListInternetBandwidthsRequestFields
	NAME                  ListInternetBandwidthsRequestFields
	ISP                   ListInternetBandwidthsRequestFields
	INGRESS_SIZE          ListInternetBandwidthsRequestFields
	ACCESS_SITE           ListInternetBandwidthsRequestFields
	SIZE                  ListInternetBandwidthsRequestFields
	DESCRIPTION           ListInternetBandwidthsRequestFields
	CHARGE_MODE           ListInternetBandwidthsRequestFields
	RATIO_95PEAK          ListInternetBandwidthsRequestFields
	FREEZEN_INFO          ListInternetBandwidthsRequestFields
	DOMAIN_ID             ListInternetBandwidthsRequestFields
	STATUS                ListInternetBandwidthsRequestFields
	CREATED_AT            ListInternetBandwidthsRequestFields
	UPDATED_AT            ListInternetBandwidthsRequestFields
	IS_PRE_PAID           ListInternetBandwidthsRequestFields
	TYPE                  ListInternetBandwidthsRequestFields
	TAGS                  ListInternetBandwidthsRequestFields
	SYS_TAGS              ListInternetBandwidthsRequestFields
	ENTERPRISE_PROJECT_ID ListInternetBandwidthsRequestFields
}

func GetListInternetBandwidthsRequestFieldsEnum() ListInternetBandwidthsRequestFieldsEnum {
	return ListInternetBandwidthsRequestFieldsEnum{
		ID: ListInternetBandwidthsRequestFields{
			value: "id",
		},
		NAME: ListInternetBandwidthsRequestFields{
			value: "name",
		},
		ISP: ListInternetBandwidthsRequestFields{
			value: "isp",
		},
		INGRESS_SIZE: ListInternetBandwidthsRequestFields{
			value: "ingress_size",
		},
		ACCESS_SITE: ListInternetBandwidthsRequestFields{
			value: "access_site",
		},
		SIZE: ListInternetBandwidthsRequestFields{
			value: "size",
		},
		DESCRIPTION: ListInternetBandwidthsRequestFields{
			value: "description",
		},
		CHARGE_MODE: ListInternetBandwidthsRequestFields{
			value: "charge_mode",
		},
		RATIO_95PEAK: ListInternetBandwidthsRequestFields{
			value: "ratio_95peak",
		},
		FREEZEN_INFO: ListInternetBandwidthsRequestFields{
			value: "freezen_info",
		},
		DOMAIN_ID: ListInternetBandwidthsRequestFields{
			value: "domain_id",
		},
		STATUS: ListInternetBandwidthsRequestFields{
			value: "status",
		},
		CREATED_AT: ListInternetBandwidthsRequestFields{
			value: "created_at",
		},
		UPDATED_AT: ListInternetBandwidthsRequestFields{
			value: "updated_at",
		},
		IS_PRE_PAID: ListInternetBandwidthsRequestFields{
			value: "is_pre_paid",
		},
		TYPE: ListInternetBandwidthsRequestFields{
			value: "type",
		},
		TAGS: ListInternetBandwidthsRequestFields{
			value: "tags",
		},
		SYS_TAGS: ListInternetBandwidthsRequestFields{
			value: "sys_tags",
		},
		ENTERPRISE_PROJECT_ID: ListInternetBandwidthsRequestFields{
			value: "enterprise_project_id",
		},
	}
}

func (c ListInternetBandwidthsRequestFields) Value() string {
	return c.value
}

func (c ListInternetBandwidthsRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInternetBandwidthsRequestFields) UnmarshalJSON(b []byte) error {
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

type ListInternetBandwidthsRequestExtFields struct {
	value string
}

type ListInternetBandwidthsRequestExtFieldsEnum struct {
	GEIP_COUNT         ListInternetBandwidthsRequestExtFields
	GEIP_SEGMENT_COUNT ListInternetBandwidthsRequestExtFields
}

func GetListInternetBandwidthsRequestExtFieldsEnum() ListInternetBandwidthsRequestExtFieldsEnum {
	return ListInternetBandwidthsRequestExtFieldsEnum{
		GEIP_COUNT: ListInternetBandwidthsRequestExtFields{
			value: "geip_count",
		},
		GEIP_SEGMENT_COUNT: ListInternetBandwidthsRequestExtFields{
			value: "geip_segment_count",
		},
	}
}

func (c ListInternetBandwidthsRequestExtFields) Value() string {
	return c.value
}

func (c ListInternetBandwidthsRequestExtFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInternetBandwidthsRequestExtFields) UnmarshalJSON(b []byte) error {
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

type ListInternetBandwidthsRequestSortKey struct {
	value string
}

type ListInternetBandwidthsRequestSortKeyEnum struct {
	ID         ListInternetBandwidthsRequestSortKey
	SIZE       ListInternetBandwidthsRequestSortKey
	CREATED_AT ListInternetBandwidthsRequestSortKey
	UPDATED_AT ListInternetBandwidthsRequestSortKey
}

func GetListInternetBandwidthsRequestSortKeyEnum() ListInternetBandwidthsRequestSortKeyEnum {
	return ListInternetBandwidthsRequestSortKeyEnum{
		ID: ListInternetBandwidthsRequestSortKey{
			value: "id",
		},
		SIZE: ListInternetBandwidthsRequestSortKey{
			value: "size",
		},
		CREATED_AT: ListInternetBandwidthsRequestSortKey{
			value: "created_at",
		},
		UPDATED_AT: ListInternetBandwidthsRequestSortKey{
			value: "updated_at",
		},
	}
}

func (c ListInternetBandwidthsRequestSortKey) Value() string {
	return c.value
}

func (c ListInternetBandwidthsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInternetBandwidthsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListInternetBandwidthsRequestSortDir struct {
	value string
}

type ListInternetBandwidthsRequestSortDirEnum struct {
	ASC  ListInternetBandwidthsRequestSortDir
	DESC ListInternetBandwidthsRequestSortDir
}

func GetListInternetBandwidthsRequestSortDirEnum() ListInternetBandwidthsRequestSortDirEnum {
	return ListInternetBandwidthsRequestSortDirEnum{
		ASC: ListInternetBandwidthsRequestSortDir{
			value: "asc",
		},
		DESC: ListInternetBandwidthsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListInternetBandwidthsRequestSortDir) Value() string {
	return c.value
}

func (c ListInternetBandwidthsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInternetBandwidthsRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListInternetBandwidthsRequestStatus struct {
	value string
}

type ListInternetBandwidthsRequestStatusEnum struct {
	FREEZED ListInternetBandwidthsRequestStatus
	NORMAL  ListInternetBandwidthsRequestStatus
}

func GetListInternetBandwidthsRequestStatusEnum() ListInternetBandwidthsRequestStatusEnum {
	return ListInternetBandwidthsRequestStatusEnum{
		FREEZED: ListInternetBandwidthsRequestStatus{
			value: "freezed",
		},
		NORMAL: ListInternetBandwidthsRequestStatus{
			value: "normal",
		},
	}
}

func (c ListInternetBandwidthsRequestStatus) Value() string {
	return c.value
}

func (c ListInternetBandwidthsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInternetBandwidthsRequestStatus) UnmarshalJSON(b []byte) error {
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
