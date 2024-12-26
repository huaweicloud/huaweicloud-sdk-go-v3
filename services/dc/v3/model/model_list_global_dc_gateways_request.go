package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGlobalDcGatewaysRequest Request Object
type ListGlobalDcGatewaysRequest struct {

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// 上一页最后一条资源记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 返回结果按照升序(asc)或降序(desc)排列，默认为asc
	SortDir *[]ListGlobalDcGatewaysRequestSortDir `json:"sort_dir,omitempty"`

	// 根据资源ID过滤实例
	Id *[]string `json:"id,omitempty"`

	// 根据名字过滤查询，可查询多个名字。
	Name *[]string `json:"name,omitempty"`

	// 根据企业项目ID过滤资源实例
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 站点网络ID
	SiteNetworkId *[]string `json:"site_network_id,omitempty"`

	// 云连接ID
	CloudConnectionId *[]string `json:"cloud_connection_id,omitempty"`

	// 根椐资源状态过滤实例
	Status *[]string `json:"status,omitempty"`

	// 全球中心网络ID
	GlobalCenterNetworkId *[]string `json:"global_center_network_id,omitempty"`
}

func (o ListGlobalDcGatewaysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalDcGatewaysRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalDcGatewaysRequest", string(data)}, " ")
}

type ListGlobalDcGatewaysRequestSortDir struct {
	value string
}

type ListGlobalDcGatewaysRequestSortDirEnum struct {
	ASC  ListGlobalDcGatewaysRequestSortDir
	DESC ListGlobalDcGatewaysRequestSortDir
}

func GetListGlobalDcGatewaysRequestSortDirEnum() ListGlobalDcGatewaysRequestSortDirEnum {
	return ListGlobalDcGatewaysRequestSortDirEnum{
		ASC: ListGlobalDcGatewaysRequestSortDir{
			value: "asc",
		},
		DESC: ListGlobalDcGatewaysRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListGlobalDcGatewaysRequestSortDir) Value() string {
	return c.value
}

func (c ListGlobalDcGatewaysRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalDcGatewaysRequestSortDir) UnmarshalJSON(b []byte) error {
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
