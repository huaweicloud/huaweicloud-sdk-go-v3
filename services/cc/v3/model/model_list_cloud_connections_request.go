package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCloudConnectionsRequest Request Object
type ListCloudConnectionsRequest struct {

	// 每页返回的个数。 取值范围：1~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向后翻页。 翻页过程中，查询条件不能修改，包括过滤条件、排序条件、limit。
	Marker *string `json:"marker,omitempty"`

	// 根据ID查询，可查询多个ID。
	Id *[]string `json:"id,omitempty"`

	// 根据名字查询，可查询多个名字。
	Name *[]string `json:"name,omitempty"`

	// 根据描述查询，可查询多个描述。
	Description *[]string `json:"description,omitempty"`

	// 根据企业项目ID过滤列表。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 根据状态过滤云连接实例列表。ACTIVE：表示状态可用。
	Status *[]ListCloudConnectionsRequestStatus `json:"status,omitempty"`

	// 根据类型过滤云连接实例列表。
	Type *[]string `json:"type,omitempty"`

	// 根据使用场景过滤云连接实例列表。
	UsedScene *[]string `json:"used_scene,omitempty"`
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
