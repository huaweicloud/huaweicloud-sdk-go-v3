package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListCloudConnectionsRequest struct {

	// 分页查询时，每页返回的个数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询时，上一页最后一条记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 根据ID过滤云连接实例列表。
	Id *[]string `json:"id,omitempty"`

	// 根据名称过滤云连接实例列表。
	Name *[]string `json:"name,omitempty"`

	// 根据描述过滤云连接实例列表。
	Description *[]string `json:"description,omitempty"`

	// 根据状态过滤云连接实例列表。ACTIVE：表示状态可用。
	Status *[]ListCloudConnectionsRequestStatus `json:"status,omitempty"`

	// 根据企业项目ID过滤云连接实例列表。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 根据类型过滤云连接实例列表。
	Type *[]string `json:"type,omitempty"`
}

func (o ListCloudConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionsRequest", string(data)}, " ")
}

type ListCloudConnectionsRequestStatus struct {
	value string
}

type ListCloudConnectionsRequestStatusEnum struct {
	ACTIVE ListCloudConnectionsRequestStatus
}

func GetListCloudConnectionsRequestStatusEnum() ListCloudConnectionsRequestStatusEnum {
	return ListCloudConnectionsRequestStatusEnum{
		ACTIVE: ListCloudConnectionsRequestStatus{
			value: "ACTIVE",
		},
	}
}

func (c ListCloudConnectionsRequestStatus) Value() string {
	return c.value
}

func (c ListCloudConnectionsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCloudConnectionsRequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
