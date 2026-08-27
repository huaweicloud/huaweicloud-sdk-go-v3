package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportHostsDetailRequest Request Object
type ExportHostsDetailRequest struct {

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
	State *ExportHostsDetailRequestState `json:"state,omitempty"`

	// 每页显示的数量。
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 上一页显示的最后记录的id，与offset同时使用时不生效。
	Marker *string `json:"marker,omitempty"`

	// 过滤指定时间起状态变更的专属主机。 日期和时间戳的格式为ISO 8601：CCYY-MM-DDThh:mm:ss±hh:mm 如果包含“hh:mm”值，则将时区作为UTC的偏移量返回。例如，“2015-08-27T09:49:58-05:00”。如果您省略时区，则假定为UTC时区。
	ChangesSince *string `json:"changes_since,omitempty"`

	// 排序字段名称，需要结合sort_type字段一起使用。 - instance_total 云办公主机上的实例总数 - available_vcpus 云办公主机可用的vCPU核数 - available_memory 云办公主机可用的内存大小
	SortField *ExportHostsDetailRequestSortField `json:"sort_field,omitempty"`

	// 排序类型，默认升序，需要结合sort_field字段一起使用。 - ASC 升序。 - DESC 降序。
	SortType *ExportHostsDetailRequestSortType `json:"sort_type,omitempty"`

	// 语言。  - zh_CN：中文 - en_US：英文
	Language ExportHostsDetailRequestLanguage `json:"language"`
}

func (o ExportHostsDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportHostsDetailRequest struct{}"
	}

	return strings.Join([]string{"ExportHostsDetailRequest", string(data)}, " ")
}

type ExportHostsDetailRequestState struct {
	value string
}

type ExportHostsDetailRequestStateEnum struct {
	AVAILABLE ExportHostsDetailRequestState
	FAULT     ExportHostsDetailRequestState
	RELEASED  ExportHostsDetailRequestState
}

func GetExportHostsDetailRequestStateEnum() ExportHostsDetailRequestStateEnum {
	return ExportHostsDetailRequestStateEnum{
		AVAILABLE: ExportHostsDetailRequestState{
			value: "available",
		},
		FAULT: ExportHostsDetailRequestState{
			value: "fault",
		},
		RELEASED: ExportHostsDetailRequestState{
			value: "released",
		},
	}
}

func (c ExportHostsDetailRequestState) Value() string {
	return c.value
}

func (c ExportHostsDetailRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportHostsDetailRequestState) UnmarshalJSON(b []byte) error {
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

type ExportHostsDetailRequestSortField struct {
	value string
}

type ExportHostsDetailRequestSortFieldEnum struct {
	INSTANCE_TOTAL   ExportHostsDetailRequestSortField
	AVAILABLE_VCPUS  ExportHostsDetailRequestSortField
	AVAILABLE_MEMORY ExportHostsDetailRequestSortField
}

func GetExportHostsDetailRequestSortFieldEnum() ExportHostsDetailRequestSortFieldEnum {
	return ExportHostsDetailRequestSortFieldEnum{
		INSTANCE_TOTAL: ExportHostsDetailRequestSortField{
			value: "instance_total",
		},
		AVAILABLE_VCPUS: ExportHostsDetailRequestSortField{
			value: "available_vcpus",
		},
		AVAILABLE_MEMORY: ExportHostsDetailRequestSortField{
			value: "available_memory",
		},
	}
}

func (c ExportHostsDetailRequestSortField) Value() string {
	return c.value
}

func (c ExportHostsDetailRequestSortField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportHostsDetailRequestSortField) UnmarshalJSON(b []byte) error {
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

type ExportHostsDetailRequestSortType struct {
	value string
}

type ExportHostsDetailRequestSortTypeEnum struct {
	ASC  ExportHostsDetailRequestSortType
	DESC ExportHostsDetailRequestSortType
}

func GetExportHostsDetailRequestSortTypeEnum() ExportHostsDetailRequestSortTypeEnum {
	return ExportHostsDetailRequestSortTypeEnum{
		ASC: ExportHostsDetailRequestSortType{
			value: "ASC",
		},
		DESC: ExportHostsDetailRequestSortType{
			value: "DESC",
		},
	}
}

func (c ExportHostsDetailRequestSortType) Value() string {
	return c.value
}

func (c ExportHostsDetailRequestSortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportHostsDetailRequestSortType) UnmarshalJSON(b []byte) error {
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

type ExportHostsDetailRequestLanguage struct {
	value string
}

type ExportHostsDetailRequestLanguageEnum struct {
	ZH_CN ExportHostsDetailRequestLanguage
	EN_US ExportHostsDetailRequestLanguage
}

func GetExportHostsDetailRequestLanguageEnum() ExportHostsDetailRequestLanguageEnum {
	return ExportHostsDetailRequestLanguageEnum{
		ZH_CN: ExportHostsDetailRequestLanguage{
			value: "zh_CN",
		},
		EN_US: ExportHostsDetailRequestLanguage{
			value: "en_US",
		},
	}
}

func (c ExportHostsDetailRequestLanguage) Value() string {
	return c.value
}

func (c ExportHostsDetailRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportHostsDetailRequestLanguage) UnmarshalJSON(b []byte) error {
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
