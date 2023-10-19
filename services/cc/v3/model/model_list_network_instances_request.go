package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListNetworkInstancesRequest Request Object
type ListNetworkInstancesRequest struct {

	// 每页返回的个数。 取值范围：1~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向翻页。 翻页过程中，查询条件不能修改，包括过滤条件，排序条件，limit。
	Marker *string `json:"marker,omitempty"`

	// 根据id查询，可查询多个id。
	Id *[]string `json:"id,omitempty"`

	// 根据名字查询，可查询多个名字。
	Name *[]string `json:"name,omitempty"`

	// 根据描述查询，可查询多个描述。
	Description *[]string `json:"description,omitempty"`

	// 根据云连接的ID过滤列表。
	CloudConnectionId *[]string `json:"cloud_connection_id,omitempty"`

	// 根据状态过滤网络实例列表。ACTIVE：表示状态可用。
	Status *[]ListNetworkInstancesRequestStatus `json:"status,omitempty"`

	// 根据类型过滤网络实例列表。
	Type *[]string `json:"type,omitempty"`

	// 根据网络实例ID过滤网络实例列表。
	InstanceId *[]string `json:"instance_id,omitempty"`

	// 根据网络实例所在的Region过滤网络实例列表。
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
