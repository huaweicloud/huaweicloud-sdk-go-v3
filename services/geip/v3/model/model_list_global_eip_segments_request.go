package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGlobalEipSegmentsRequest Request Object
type ListGlobalEipSegmentsRequest struct {

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *int32 `json:"offset,omitempty"`

	// 分页起始点
	Marker *string `json:"marker,omitempty"`

	// 翻页方向
	PageReverse *bool `json:"page_reverse,omitempty"`

	Fields *[]ListGlobalEipSegmentsRequestFields `json:"fields,omitempty"`

	// 按照sort_key指定的字段排序
	SortKey *[]ListGlobalEipSegmentsRequestSortKey `json:"sort_key,omitempty"`

	// 排序的方向，倒序或者正序
	SortDir *[]ListGlobalEipSegmentsRequestSortDir `json:"sort_dir,omitempty"`

	Id *[]string `json:"id,omitempty"`

	InternetBandwidthId *[]string `json:"internet_bandwidth_id,omitempty"`

	Name *[]string `json:"name,omitempty"`

	NameLike *string `json:"name_like,omitempty"`

	AccessSite *[]string `json:"access_site,omitempty"`

	GeipPoolName *[]string `json:"geip_pool_name,omitempty"`

	Isp *[]string `json:"isp,omitempty"`

	IpVersion *[]ListGlobalEipSegmentsRequestIpVersion `json:"ip_version,omitempty"`

	Cidr *[]string `json:"cidr,omitempty"`

	CidrV6 *[]string `json:"cidr_v6,omitempty"`

	Freezen *[]bool `json:"freezen,omitempty"`

	InternetBandwidthIsNull *[]bool `json:"internet_bandwidth_is_null,omitempty"`

	Status *[]ListGlobalEipSegmentsRequestStatus `json:"status,omitempty"`

	AssociateInstanceRegion *[]string `json:"associate_instance.region,omitempty"`

	AssociateInstanceInstanceType *[]string `json:"associate_instance.instance_type,omitempty"`

	AssociateInstancePublicBorderGroup *[]string `json:"associate_instance.public_border_group,omitempty"`

	AssociateInstanceInstanceSite *[]string `json:"associate_instance.instance_site,omitempty"`

	AssociateInstanceInstanceId *[]string `json:"associate_instance.instance_id,omitempty"`

	AssociateInstanceProjectId *[]string `json:"associate_instance.project_id,omitempty"`

	AssociateInstanceServiceId *[]string `json:"associate_instance.service_id,omitempty"`

	AssociateInstanceServiceType *[]string `json:"associate_instance.service_type,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	Tags *[]string `json:"tags,omitempty"`
}

func (o ListGlobalEipSegmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalEipSegmentsRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalEipSegmentsRequest", string(data)}, " ")
}

type ListGlobalEipSegmentsRequestFields struct {
	value string
}

type ListGlobalEipSegmentsRequestFieldsEnum struct {
	ID                    ListGlobalEipSegmentsRequestFields
	NAME                  ListGlobalEipSegmentsRequestFields
	DESCRIPTION           ListGlobalEipSegmentsRequestFields
	DOMAIN_ID             ListGlobalEipSegmentsRequestFields
	ACCESS_SITE           ListGlobalEipSegmentsRequestFields
	GEIP_POOL_NAME        ListGlobalEipSegmentsRequestFields
	ISP                   ListGlobalEipSegmentsRequestFields
	IP_VERSION            ListGlobalEipSegmentsRequestFields
	CIDR                  ListGlobalEipSegmentsRequestFields
	CIDR_V6               ListGlobalEipSegmentsRequestFields
	FREEZEN               ListGlobalEipSegmentsRequestFields
	FREEZEN_INFO          ListGlobalEipSegmentsRequestFields
	STATUS                ListGlobalEipSegmentsRequestFields
	CREATED_AT            ListGlobalEipSegmentsRequestFields
	UPDATED_AT            ListGlobalEipSegmentsRequestFields
	INTERNET_BANDWIDTH    ListGlobalEipSegmentsRequestFields
	ASSOCIATE_INSTANCE    ListGlobalEipSegmentsRequestFields
	IS_PRE_PAID           ListGlobalEipSegmentsRequestFields
	TAGS                  ListGlobalEipSegmentsRequestFields
	SYS_TAGS              ListGlobalEipSegmentsRequestFields
	ENTERPRISE_PROJECT_ID ListGlobalEipSegmentsRequestFields
}

func GetListGlobalEipSegmentsRequestFieldsEnum() ListGlobalEipSegmentsRequestFieldsEnum {
	return ListGlobalEipSegmentsRequestFieldsEnum{
		ID: ListGlobalEipSegmentsRequestFields{
			value: "id",
		},
		NAME: ListGlobalEipSegmentsRequestFields{
			value: "name",
		},
		DESCRIPTION: ListGlobalEipSegmentsRequestFields{
			value: "description",
		},
		DOMAIN_ID: ListGlobalEipSegmentsRequestFields{
			value: "domain_id",
		},
		ACCESS_SITE: ListGlobalEipSegmentsRequestFields{
			value: "access_site",
		},
		GEIP_POOL_NAME: ListGlobalEipSegmentsRequestFields{
			value: "geip_pool_name",
		},
		ISP: ListGlobalEipSegmentsRequestFields{
			value: "isp",
		},
		IP_VERSION: ListGlobalEipSegmentsRequestFields{
			value: "ip_version",
		},
		CIDR: ListGlobalEipSegmentsRequestFields{
			value: "cidr",
		},
		CIDR_V6: ListGlobalEipSegmentsRequestFields{
			value: "cidr_v6",
		},
		FREEZEN: ListGlobalEipSegmentsRequestFields{
			value: "freezen",
		},
		FREEZEN_INFO: ListGlobalEipSegmentsRequestFields{
			value: "freezen_info",
		},
		STATUS: ListGlobalEipSegmentsRequestFields{
			value: "status",
		},
		CREATED_AT: ListGlobalEipSegmentsRequestFields{
			value: "created_at",
		},
		UPDATED_AT: ListGlobalEipSegmentsRequestFields{
			value: "updated_at",
		},
		INTERNET_BANDWIDTH: ListGlobalEipSegmentsRequestFields{
			value: "internet_bandwidth",
		},
		ASSOCIATE_INSTANCE: ListGlobalEipSegmentsRequestFields{
			value: "associate_instance",
		},
		IS_PRE_PAID: ListGlobalEipSegmentsRequestFields{
			value: "is_pre_paid",
		},
		TAGS: ListGlobalEipSegmentsRequestFields{
			value: "tags",
		},
		SYS_TAGS: ListGlobalEipSegmentsRequestFields{
			value: "sys_tags",
		},
		ENTERPRISE_PROJECT_ID: ListGlobalEipSegmentsRequestFields{
			value: "enterprise_project_id",
		},
	}
}

func (c ListGlobalEipSegmentsRequestFields) Value() string {
	return c.value
}

func (c ListGlobalEipSegmentsRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalEipSegmentsRequestFields) UnmarshalJSON(b []byte) error {
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

type ListGlobalEipSegmentsRequestSortKey struct {
	value string
}

type ListGlobalEipSegmentsRequestSortKeyEnum struct {
	ID         ListGlobalEipSegmentsRequestSortKey
	CIDR       ListGlobalEipSegmentsRequestSortKey
	CREATED_AT ListGlobalEipSegmentsRequestSortKey
	UPDATED_AT ListGlobalEipSegmentsRequestSortKey
}

func GetListGlobalEipSegmentsRequestSortKeyEnum() ListGlobalEipSegmentsRequestSortKeyEnum {
	return ListGlobalEipSegmentsRequestSortKeyEnum{
		ID: ListGlobalEipSegmentsRequestSortKey{
			value: "id",
		},
		CIDR: ListGlobalEipSegmentsRequestSortKey{
			value: "cidr",
		},
		CREATED_AT: ListGlobalEipSegmentsRequestSortKey{
			value: "created_at",
		},
		UPDATED_AT: ListGlobalEipSegmentsRequestSortKey{
			value: "updated_at",
		},
	}
}

func (c ListGlobalEipSegmentsRequestSortKey) Value() string {
	return c.value
}

func (c ListGlobalEipSegmentsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalEipSegmentsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListGlobalEipSegmentsRequestSortDir struct {
	value string
}

type ListGlobalEipSegmentsRequestSortDirEnum struct {
	ASC  ListGlobalEipSegmentsRequestSortDir
	DESC ListGlobalEipSegmentsRequestSortDir
}

func GetListGlobalEipSegmentsRequestSortDirEnum() ListGlobalEipSegmentsRequestSortDirEnum {
	return ListGlobalEipSegmentsRequestSortDirEnum{
		ASC: ListGlobalEipSegmentsRequestSortDir{
			value: "asc",
		},
		DESC: ListGlobalEipSegmentsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListGlobalEipSegmentsRequestSortDir) Value() string {
	return c.value
}

func (c ListGlobalEipSegmentsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalEipSegmentsRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListGlobalEipSegmentsRequestIpVersion struct {
	value int32
}

type ListGlobalEipSegmentsRequestIpVersionEnum struct {
	E_4 ListGlobalEipSegmentsRequestIpVersion
	E_6 ListGlobalEipSegmentsRequestIpVersion
}

func GetListGlobalEipSegmentsRequestIpVersionEnum() ListGlobalEipSegmentsRequestIpVersionEnum {
	return ListGlobalEipSegmentsRequestIpVersionEnum{
		E_4: ListGlobalEipSegmentsRequestIpVersion{
			value: 4,
		}, E_6: ListGlobalEipSegmentsRequestIpVersion{
			value: 6,
		},
	}
}

func (c ListGlobalEipSegmentsRequestIpVersion) Value() int32 {
	return c.value
}

func (c ListGlobalEipSegmentsRequestIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalEipSegmentsRequestIpVersion) UnmarshalJSON(b []byte) error {
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

type ListGlobalEipSegmentsRequestStatus struct {
	value string
}

type ListGlobalEipSegmentsRequestStatusEnum struct {
	IDLE  ListGlobalEipSegmentsRequestStatus
	INUSE ListGlobalEipSegmentsRequestStatus
}

func GetListGlobalEipSegmentsRequestStatusEnum() ListGlobalEipSegmentsRequestStatusEnum {
	return ListGlobalEipSegmentsRequestStatusEnum{
		IDLE: ListGlobalEipSegmentsRequestStatus{
			value: "idle",
		},
		INUSE: ListGlobalEipSegmentsRequestStatus{
			value: "inuse",
		},
	}
}

func (c ListGlobalEipSegmentsRequestStatus) Value() string {
	return c.value
}

func (c ListGlobalEipSegmentsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalEipSegmentsRequestStatus) UnmarshalJSON(b []byte) error {
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
