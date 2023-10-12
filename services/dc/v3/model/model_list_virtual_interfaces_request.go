package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListVirtualInterfacesRequest Request Object
type ListVirtualInterfacesRequest struct {

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条资源记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// 返回结果按照升序(asc)或降序(desc)排列，默认为asc
	SortDir *[]ListVirtualInterfacesRequestSortDir `json:"sort_dir,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 根据企业项目ID过滤资源实例
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 根据资源ID过滤实例
	Id *[]string `json:"id,omitempty"`

	// 根椐资源状态过淲实例
	Status *[]string `json:"status,omitempty"`

	// 根椐物理专线ID过滤查询实例信息
	DirectConnectId *[]string `json:"direct_connect_id,omitempty"`

	// 根椐虚拟网关ID过滤查询实例信息
	VgwId *[]string `json:"vgw_id,omitempty"`
}

func (o ListVirtualInterfacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVirtualInterfacesRequest struct{}"
	}

	return strings.Join([]string{"ListVirtualInterfacesRequest", string(data)}, " ")
}

type ListVirtualInterfacesRequestSortDir struct {
	value string
}

type ListVirtualInterfacesRequestSortDirEnum struct {
	ASC  ListVirtualInterfacesRequestSortDir
	DESC ListVirtualInterfacesRequestSortDir
}

func GetListVirtualInterfacesRequestSortDirEnum() ListVirtualInterfacesRequestSortDirEnum {
	return ListVirtualInterfacesRequestSortDirEnum{
		ASC: ListVirtualInterfacesRequestSortDir{
			value: "asc",
		},
		DESC: ListVirtualInterfacesRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListVirtualInterfacesRequestSortDir) Value() string {
	return c.value
}

func (c ListVirtualInterfacesRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVirtualInterfacesRequestSortDir) UnmarshalJSON(b []byte) error {
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
