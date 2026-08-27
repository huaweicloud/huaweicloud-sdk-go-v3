package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHostsDetailRequest Request Object
type ListHostsDetailRequest struct {

	// 云办公主机名称。
	Name *string `json:"name,omitempty"`

	// 云办公主机所属区域。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 云办公主机的id。
	HostId *string `json:"host_id,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 类型。
	HostType *string `json:"host_type,omitempty"`

	// 云办公主机类型名称。
	HostTypeName *string `json:"host_type_name,omitempty"`

	// 云办公主机状态，available-可用的，fault-错误的，released-释放的。
	State *ListHostsDetailRequestState `json:"state,omitempty"`

	// 每页显示的数量。
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 上一页显示的最后记录的id，与offset同时使用时不生效。
	Marker *string `json:"marker,omitempty"`

	// 过滤指定时间起状态变更的专属主机。 日期和时间戳的格式为ISO 8601：CCYY-MM-DDThh:mm:ss±hh:mm 如果包含“hh:mm”值，则将时区作为UTC的偏移量返回。例如，“2015-08-27T09:49:58-05:00”。如果您省略时区，则假定为UTC时区。
	ChangesSince *string `json:"changes_since,omitempty"`

	// 排序字段名称，需要结合sort_type字段一起使用。 - instance_total 云办公主机上的实例总数 - available_vcpus 云办公主机可用的vCPU核数 - available_memory 云办公主机可用的内存大小
	SortField *ListHostsDetailRequestSortField `json:"sort_field,omitempty"`

	// 排序类型，默认升序，需要结合sort_field字段一起使用。 - ASC 升序。 - DESC 降序。
	SortType *ListHostsDetailRequestSortType `json:"sort_type,omitempty"`
}

func (o ListHostsDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsDetailRequest struct{}"
	}

	return strings.Join([]string{"ListHostsDetailRequest", string(data)}, " ")
}

type ListHostsDetailRequestState struct {
	value string
}

type ListHostsDetailRequestStateEnum struct {
	AVAILABLE ListHostsDetailRequestState
	FAULT     ListHostsDetailRequestState
	RELEASED  ListHostsDetailRequestState
}

func GetListHostsDetailRequestStateEnum() ListHostsDetailRequestStateEnum {
	return ListHostsDetailRequestStateEnum{
		AVAILABLE: ListHostsDetailRequestState{
			value: "available",
		},
		FAULT: ListHostsDetailRequestState{
			value: "fault",
		},
		RELEASED: ListHostsDetailRequestState{
			value: "released",
		},
	}
}

func (c ListHostsDetailRequestState) Value() string {
	return c.value
}

func (c ListHostsDetailRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHostsDetailRequestState) UnmarshalJSON(b []byte) error {
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

type ListHostsDetailRequestSortField struct {
	value string
}

type ListHostsDetailRequestSortFieldEnum struct {
	INSTANCE_TOTAL   ListHostsDetailRequestSortField
	AVAILABLE_VCPUS  ListHostsDetailRequestSortField
	AVAILABLE_MEMORY ListHostsDetailRequestSortField
}

func GetListHostsDetailRequestSortFieldEnum() ListHostsDetailRequestSortFieldEnum {
	return ListHostsDetailRequestSortFieldEnum{
		INSTANCE_TOTAL: ListHostsDetailRequestSortField{
			value: "instance_total",
		},
		AVAILABLE_VCPUS: ListHostsDetailRequestSortField{
			value: "available_vcpus",
		},
		AVAILABLE_MEMORY: ListHostsDetailRequestSortField{
			value: "available_memory",
		},
	}
}

func (c ListHostsDetailRequestSortField) Value() string {
	return c.value
}

func (c ListHostsDetailRequestSortField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHostsDetailRequestSortField) UnmarshalJSON(b []byte) error {
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

type ListHostsDetailRequestSortType struct {
	value string
}

type ListHostsDetailRequestSortTypeEnum struct {
	ASC  ListHostsDetailRequestSortType
	DESC ListHostsDetailRequestSortType
}

func GetListHostsDetailRequestSortTypeEnum() ListHostsDetailRequestSortTypeEnum {
	return ListHostsDetailRequestSortTypeEnum{
		ASC: ListHostsDetailRequestSortType{
			value: "ASC",
		},
		DESC: ListHostsDetailRequestSortType{
			value: "DESC",
		},
	}
}

func (c ListHostsDetailRequestSortType) Value() string {
	return c.value
}

func (c ListHostsDetailRequestSortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHostsDetailRequestSortType) UnmarshalJSON(b []byte) error {
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
