package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGlobalEipsRequest Request Object
type ListGlobalEipsRequest struct {

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *int32 `json:"offset,omitempty"`

	// 分页起始点
	Marker *string `json:"marker,omitempty"`

	// 翻页方向
	PageReverse *bool `json:"page_reverse,omitempty"`

	Fields *[]ListGlobalEipsRequestFields `json:"fields,omitempty"`

	// 按照sort_key指定的字段排序
	SortKey *[]ListGlobalEipsRequestSortKey `json:"sort_key,omitempty"`

	// 排序的方向，倒序或者正序
	SortDir *[]ListGlobalEipsRequestSortDir `json:"sort_dir,omitempty"`

	Id *[]string `json:"id,omitempty"`

	InternetBandwidthId *[]string `json:"internet_bandwidth_id,omitempty"`

	Name *[]string `json:"name,omitempty"`

	NameLike *string `json:"name_like,omitempty"`

	AccessSite *[]string `json:"access_site,omitempty"`

	GeipPoolName *[]string `json:"geip_pool_name,omitempty"`

	Isp *[]string `json:"isp,omitempty"`

	IpVersion *[]ListGlobalEipsRequestIpVersion `json:"ip_version,omitempty"`

	IpAddress *[]string `json:"ip_address,omitempty"`

	Ipv6Address *[]string `json:"ipv6_address,omitempty"`

	Freezen *[]bool `json:"freezen,omitempty"`

	Polluted *[]bool `json:"polluted,omitempty"`

	InternetBandwidthIsNull *[]bool `json:"internet_bandwidth_is_null,omitempty"`

	GcbBandwidthIsNull *[]bool `json:"gcb_bandwidth_is_null,omitempty"`

	Status *[]ListGlobalEipsRequestStatus `json:"status,omitempty"`

	AssociateInstanceInfoRegion *[]string `json:"associate_instance_info.region,omitempty"`

	AssociateInstanceInfoInstanceType *[]string `json:"associate_instance_info.instance_type,omitempty"`

	AssociateInstanceInfoPublicBorderGroup *[]string `json:"associate_instance_info.public_border_group,omitempty"`

	AssociateInstanceInfoInstanceSite *[]string `json:"associate_instance_info.instance_site,omitempty"`

	AssociateInstanceInfoInstanceId *[]string `json:"associate_instance_info.instance_id,omitempty"`

	AssociateInstanceInfoProjectId *[]string `json:"associate_instance_info.project_id,omitempty"`

	AssociateInstanceInfoServiceId *[]string `json:"associate_instance_info.service_id,omitempty"`

	AssociateInstanceInfoServiceType *[]string `json:"associate_instance_info.service_type,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	Tags *[]string `json:"tags,omitempty"`
}

func (o ListGlobalEipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalEipsRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalEipsRequest", string(data)}, " ")
}

type ListGlobalEipsRequestFields struct {
	value string
}

type ListGlobalEipsRequestFieldsEnum struct {
	ID                               ListGlobalEipsRequestFields
	NAME                             ListGlobalEipsRequestFields
	DESCRIPTION                      ListGlobalEipsRequestFields
	DOMAIN_ID                        ListGlobalEipsRequestFields
	ACCESS_SITE                      ListGlobalEipsRequestFields
	GEIP_POOL_NAME                   ListGlobalEipsRequestFields
	ISP                              ListGlobalEipsRequestFields
	IP_VERSION                       ListGlobalEipsRequestFields
	IP_ADDRESS                       ListGlobalEipsRequestFields
	IPV6_ADDRESS                     ListGlobalEipsRequestFields
	FREEZEN                          ListGlobalEipsRequestFields
	FREEZEN_INFO                     ListGlobalEipsRequestFields
	POLLUTED                         ListGlobalEipsRequestFields
	STATUS                           ListGlobalEipsRequestFields
	CREATED_AT                       ListGlobalEipsRequestFields
	UPDATED_AT                       ListGlobalEipsRequestFields
	INTERNET_BANDWIDTH_INFO          ListGlobalEipsRequestFields
	GLOBAL_CONNECTION_BANDWIDTH_INFO ListGlobalEipsRequestFields
	ASSOCIATE_INSTANCE_INFO          ListGlobalEipsRequestFields
	IS_PRE_PAID                      ListGlobalEipsRequestFields
	TAGS                             ListGlobalEipsRequestFields
	SYS_TAGS                         ListGlobalEipsRequestFields
	ENTERPRISE_PROJECT_ID            ListGlobalEipsRequestFields
}

func GetListGlobalEipsRequestFieldsEnum() ListGlobalEipsRequestFieldsEnum {
	return ListGlobalEipsRequestFieldsEnum{
		ID: ListGlobalEipsRequestFields{
			value: "id",
		},
		NAME: ListGlobalEipsRequestFields{
			value: "name",
		},
		DESCRIPTION: ListGlobalEipsRequestFields{
			value: "description",
		},
		DOMAIN_ID: ListGlobalEipsRequestFields{
			value: "domain_id",
		},
		ACCESS_SITE: ListGlobalEipsRequestFields{
			value: "access_site",
		},
		GEIP_POOL_NAME: ListGlobalEipsRequestFields{
			value: "geip_pool_name",
		},
		ISP: ListGlobalEipsRequestFields{
			value: "isp",
		},
		IP_VERSION: ListGlobalEipsRequestFields{
			value: "ip_version",
		},
		IP_ADDRESS: ListGlobalEipsRequestFields{
			value: "ip_address",
		},
		IPV6_ADDRESS: ListGlobalEipsRequestFields{
			value: "ipv6_address",
		},
		FREEZEN: ListGlobalEipsRequestFields{
			value: "freezen",
		},
		FREEZEN_INFO: ListGlobalEipsRequestFields{
			value: "freezen_info",
		},
		POLLUTED: ListGlobalEipsRequestFields{
			value: "polluted",
		},
		STATUS: ListGlobalEipsRequestFields{
			value: "status",
		},
		CREATED_AT: ListGlobalEipsRequestFields{
			value: "created_at",
		},
		UPDATED_AT: ListGlobalEipsRequestFields{
			value: "updated_at",
		},
		INTERNET_BANDWIDTH_INFO: ListGlobalEipsRequestFields{
			value: "internet_bandwidth_info",
		},
		GLOBAL_CONNECTION_BANDWIDTH_INFO: ListGlobalEipsRequestFields{
			value: "global_connection_bandwidth_info",
		},
		ASSOCIATE_INSTANCE_INFO: ListGlobalEipsRequestFields{
			value: "associate_instance_info",
		},
		IS_PRE_PAID: ListGlobalEipsRequestFields{
			value: "is_pre_paid",
		},
		TAGS: ListGlobalEipsRequestFields{
			value: "tags",
		},
		SYS_TAGS: ListGlobalEipsRequestFields{
			value: "sys_tags",
		},
		ENTERPRISE_PROJECT_ID: ListGlobalEipsRequestFields{
			value: "enterprise_project_id",
		},
	}
}

func (c ListGlobalEipsRequestFields) Value() string {
	return c.value
}

func (c ListGlobalEipsRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalEipsRequestFields) UnmarshalJSON(b []byte) error {
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

type ListGlobalEipsRequestSortKey struct {
	value string
}

type ListGlobalEipsRequestSortKeyEnum struct {
	ID         ListGlobalEipsRequestSortKey
	IP_ADDRESS ListGlobalEipsRequestSortKey
	CREATED_AT ListGlobalEipsRequestSortKey
	UPDATED_AT ListGlobalEipsRequestSortKey
}

func GetListGlobalEipsRequestSortKeyEnum() ListGlobalEipsRequestSortKeyEnum {
	return ListGlobalEipsRequestSortKeyEnum{
		ID: ListGlobalEipsRequestSortKey{
			value: "id",
		},
		IP_ADDRESS: ListGlobalEipsRequestSortKey{
			value: "ip_address",
		},
		CREATED_AT: ListGlobalEipsRequestSortKey{
			value: "created_at",
		},
		UPDATED_AT: ListGlobalEipsRequestSortKey{
			value: "updated_at",
		},
	}
}

func (c ListGlobalEipsRequestSortKey) Value() string {
	return c.value
}

func (c ListGlobalEipsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalEipsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListGlobalEipsRequestSortDir struct {
	value string
}

type ListGlobalEipsRequestSortDirEnum struct {
	ASC  ListGlobalEipsRequestSortDir
	DESC ListGlobalEipsRequestSortDir
}

func GetListGlobalEipsRequestSortDirEnum() ListGlobalEipsRequestSortDirEnum {
	return ListGlobalEipsRequestSortDirEnum{
		ASC: ListGlobalEipsRequestSortDir{
			value: "asc",
		},
		DESC: ListGlobalEipsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListGlobalEipsRequestSortDir) Value() string {
	return c.value
}

func (c ListGlobalEipsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalEipsRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListGlobalEipsRequestIpVersion struct {
	value int32
}

type ListGlobalEipsRequestIpVersionEnum struct {
	E_4 ListGlobalEipsRequestIpVersion
	E_6 ListGlobalEipsRequestIpVersion
}

func GetListGlobalEipsRequestIpVersionEnum() ListGlobalEipsRequestIpVersionEnum {
	return ListGlobalEipsRequestIpVersionEnum{
		E_4: ListGlobalEipsRequestIpVersion{
			value: 4,
		}, E_6: ListGlobalEipsRequestIpVersion{
			value: 6,
		},
	}
}

func (c ListGlobalEipsRequestIpVersion) Value() int32 {
	return c.value
}

func (c ListGlobalEipsRequestIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalEipsRequestIpVersion) UnmarshalJSON(b []byte) error {
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

type ListGlobalEipsRequestStatus struct {
	value string
}

type ListGlobalEipsRequestStatusEnum struct {
	IDLE           ListGlobalEipsRequestStatus
	INUSE          ListGlobalEipsRequestStatus
	PENDING_CREATE ListGlobalEipsRequestStatus
	PENDING_UPDATE ListGlobalEipsRequestStatus
}

func GetListGlobalEipsRequestStatusEnum() ListGlobalEipsRequestStatusEnum {
	return ListGlobalEipsRequestStatusEnum{
		IDLE: ListGlobalEipsRequestStatus{
			value: "idle",
		},
		INUSE: ListGlobalEipsRequestStatus{
			value: "inuse",
		},
		PENDING_CREATE: ListGlobalEipsRequestStatus{
			value: "pending_create",
		},
		PENDING_UPDATE: ListGlobalEipsRequestStatus{
			value: "pending_update",
		},
	}
}

func (c ListGlobalEipsRequestStatus) Value() string {
	return c.value
}

func (c ListGlobalEipsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalEipsRequestStatus) UnmarshalJSON(b []byte) error {
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
