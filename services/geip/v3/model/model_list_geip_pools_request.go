package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGeipPoolsRequest Request Object
type ListGeipPoolsRequest struct {

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *int32 `json:"offset,omitempty"`

	// 分页起始点
	Marker *string `json:"marker,omitempty"`

	// 翻页方向
	PageReverse *bool `json:"page_reverse,omitempty"`

	Fields *[]ListGeipPoolsRequestFields `json:"fields,omitempty"`

	// 按照sort_key指定的字段排序
	SortKey *[]ListGeipPoolsRequestSortKey `json:"sort_key,omitempty"`

	// 排序的方向，倒序或者正序
	SortDir *[]ListGeipPoolsRequestSortDir `json:"sort_dir,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Code *[]string `json:"code,omitempty"`

	AccessSite *[]string `json:"access_site,omitempty"`

	Isp *[]string `json:"isp,omitempty"`

	IpVersion *[]ListGeipPoolsRequestIpVersion `json:"ip_version,omitempty"`

	Status *[]ListGeipPoolsRequestStatus `json:"status,omitempty"`

	Type *[]ListGeipPoolsRequestType `json:"type,omitempty"`
}

func (o ListGeipPoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeipPoolsRequest struct{}"
	}

	return strings.Join([]string{"ListGeipPoolsRequest", string(data)}, " ")
}

type ListGeipPoolsRequestFields struct {
	value string
}

type ListGeipPoolsRequestFieldsEnum struct {
	ID          ListGeipPoolsRequestFields
	NAME        ListGeipPoolsRequestFields
	ISP         ListGeipPoolsRequestFields
	ACCESS_SITE ListGeipPoolsRequestFields
	TYPE        ListGeipPoolsRequestFields
	IP_VERSION  ListGeipPoolsRequestFields
	EN_NAME     ListGeipPoolsRequestFields
	CN_NAME     ListGeipPoolsRequestFields
	STATUS      ListGeipPoolsRequestFields
	CREATED_AT  ListGeipPoolsRequestFields
	UPDATED_AT  ListGeipPoolsRequestFields
	DESCRIPTION ListGeipPoolsRequestFields
}

func GetListGeipPoolsRequestFieldsEnum() ListGeipPoolsRequestFieldsEnum {
	return ListGeipPoolsRequestFieldsEnum{
		ID: ListGeipPoolsRequestFields{
			value: "id",
		},
		NAME: ListGeipPoolsRequestFields{
			value: "name",
		},
		ISP: ListGeipPoolsRequestFields{
			value: "isp",
		},
		ACCESS_SITE: ListGeipPoolsRequestFields{
			value: "access_site",
		},
		TYPE: ListGeipPoolsRequestFields{
			value: "type",
		},
		IP_VERSION: ListGeipPoolsRequestFields{
			value: "ip_version",
		},
		EN_NAME: ListGeipPoolsRequestFields{
			value: "en_name",
		},
		CN_NAME: ListGeipPoolsRequestFields{
			value: "cn_name",
		},
		STATUS: ListGeipPoolsRequestFields{
			value: "status",
		},
		CREATED_AT: ListGeipPoolsRequestFields{
			value: "created_at",
		},
		UPDATED_AT: ListGeipPoolsRequestFields{
			value: "updated_at",
		},
		DESCRIPTION: ListGeipPoolsRequestFields{
			value: "description",
		},
	}
}

func (c ListGeipPoolsRequestFields) Value() string {
	return c.value
}

func (c ListGeipPoolsRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGeipPoolsRequestFields) UnmarshalJSON(b []byte) error {
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

type ListGeipPoolsRequestSortKey struct {
	value string
}

type ListGeipPoolsRequestSortKeyEnum struct {
	ID          ListGeipPoolsRequestSortKey
	NAME        ListGeipPoolsRequestSortKey
	TYPE        ListGeipPoolsRequestSortKey
	ACCESS_SITE ListGeipPoolsRequestSortKey
	CREATED_AT  ListGeipPoolsRequestSortKey
	UPDATED_AT  ListGeipPoolsRequestSortKey
}

func GetListGeipPoolsRequestSortKeyEnum() ListGeipPoolsRequestSortKeyEnum {
	return ListGeipPoolsRequestSortKeyEnum{
		ID: ListGeipPoolsRequestSortKey{
			value: "id",
		},
		NAME: ListGeipPoolsRequestSortKey{
			value: "name",
		},
		TYPE: ListGeipPoolsRequestSortKey{
			value: "type",
		},
		ACCESS_SITE: ListGeipPoolsRequestSortKey{
			value: "access_site",
		},
		CREATED_AT: ListGeipPoolsRequestSortKey{
			value: "created_at",
		},
		UPDATED_AT: ListGeipPoolsRequestSortKey{
			value: "updated_at",
		},
	}
}

func (c ListGeipPoolsRequestSortKey) Value() string {
	return c.value
}

func (c ListGeipPoolsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGeipPoolsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListGeipPoolsRequestSortDir struct {
	value string
}

type ListGeipPoolsRequestSortDirEnum struct {
	ASC  ListGeipPoolsRequestSortDir
	DESC ListGeipPoolsRequestSortDir
}

func GetListGeipPoolsRequestSortDirEnum() ListGeipPoolsRequestSortDirEnum {
	return ListGeipPoolsRequestSortDirEnum{
		ASC: ListGeipPoolsRequestSortDir{
			value: "asc",
		},
		DESC: ListGeipPoolsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListGeipPoolsRequestSortDir) Value() string {
	return c.value
}

func (c ListGeipPoolsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGeipPoolsRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListGeipPoolsRequestIpVersion struct {
	value string
}

type ListGeipPoolsRequestIpVersionEnum struct {
	E_4 ListGeipPoolsRequestIpVersion
	E_6 ListGeipPoolsRequestIpVersion
}

func GetListGeipPoolsRequestIpVersionEnum() ListGeipPoolsRequestIpVersionEnum {
	return ListGeipPoolsRequestIpVersionEnum{
		E_4: ListGeipPoolsRequestIpVersion{
			value: "4",
		},
		E_6: ListGeipPoolsRequestIpVersion{
			value: "6",
		},
	}
}

func (c ListGeipPoolsRequestIpVersion) Value() string {
	return c.value
}

func (c ListGeipPoolsRequestIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGeipPoolsRequestIpVersion) UnmarshalJSON(b []byte) error {
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

type ListGeipPoolsRequestStatus struct {
	value string
}

type ListGeipPoolsRequestStatusEnum struct {
	ACTIVE   ListGeipPoolsRequestStatus
	INACTIVE ListGeipPoolsRequestStatus
	SOLDOUT  ListGeipPoolsRequestStatus
}

func GetListGeipPoolsRequestStatusEnum() ListGeipPoolsRequestStatusEnum {
	return ListGeipPoolsRequestStatusEnum{
		ACTIVE: ListGeipPoolsRequestStatus{
			value: "active",
		},
		INACTIVE: ListGeipPoolsRequestStatus{
			value: "inactive",
		},
		SOLDOUT: ListGeipPoolsRequestStatus{
			value: "soldout",
		},
	}
}

func (c ListGeipPoolsRequestStatus) Value() string {
	return c.value
}

func (c ListGeipPoolsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGeipPoolsRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListGeipPoolsRequestType struct {
	value string
}

type ListGeipPoolsRequestTypeEnum struct {
	GEIP         ListGeipPoolsRequestType
	GEIP_SEGMENT ListGeipPoolsRequestType
}

func GetListGeipPoolsRequestTypeEnum() ListGeipPoolsRequestTypeEnum {
	return ListGeipPoolsRequestTypeEnum{
		GEIP: ListGeipPoolsRequestType{
			value: "GEIP",
		},
		GEIP_SEGMENT: ListGeipPoolsRequestType{
			value: "GEIP_SEGMENT",
		},
	}
}

func (c ListGeipPoolsRequestType) Value() string {
	return c.value
}

func (c ListGeipPoolsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGeipPoolsRequestType) UnmarshalJSON(b []byte) error {
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
