package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListNetworkInstancesRequest struct {

	// 分页查询时，每页返回的个数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询时，上一页最后一条记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 根据ID过滤网络实例列表。
	Id *[]string `json:"id,omitempty"`

	// 根据名称过滤网络实例列表。
	Name *[]string `json:"name,omitempty"`

	// 根据描述过滤网络实例列表。
	Description *[]string `json:"description,omitempty"`

	// 根据状态过滤网络实例列表。ACTIVE：表示状态可用。
	Status *[]ListNetworkInstancesRequestStatus `json:"status,omitempty"`

	// 根据类型过滤网络实例列表。
	Type *[]string `json:"type,omitempty"`

	// 根据云连接实例ID过滤网络实例列表。
	CloudConnectionId *[]string `json:"cloud_connection_id,omitempty"`

	// 根据VPC或者VGW的ID过滤网络实例列表。
	InstanceId *[]string `json:"instance_id,omitempty"`

	// 根据VPC或者VGW所在的Region过滤网络实例列表。
	RegionId *[]string `json:"region_id,omitempty"`
}

func (o ListNetworkInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetworkInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListNetworkInstancesRequest", string(data)}, " ")
}

type ListNetworkInstancesRequestStatus struct {
	value string
}

type ListNetworkInstancesRequestStatusEnum struct {
	ACTIVE ListNetworkInstancesRequestStatus
}

func GetListNetworkInstancesRequestStatusEnum() ListNetworkInstancesRequestStatusEnum {
	return ListNetworkInstancesRequestStatusEnum{
		ACTIVE: ListNetworkInstancesRequestStatus{
			value: "ACTIVE",
		},
	}
}

func (c ListNetworkInstancesRequestStatus) Value() string {
	return c.value
}

func (c ListNetworkInstancesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNetworkInstancesRequestStatus) UnmarshalJSON(b []byte) error {
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
