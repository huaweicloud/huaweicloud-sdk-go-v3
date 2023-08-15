package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEndpointsRequest Request Object
type ListEndpointsRequest struct {

	// 终端节点组ID。
	EndpointGroupId string `json:"endpoint_group_id"`

	// 分页查询每页的资源个数。如果不设置，则默认为500。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询的起始的资源ID，表示上一页最后一条查询资源记录的ID。不指定时表示查询第一页。 必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 资源ID。
	Id *string `json:"id,omitempty"`

	// 配置状态，可选范围: - ACTIVE：运行中 - PENDING：待定 - ERROR：错误 - DELETING：正在删除
	Status *ListEndpointsRequestStatus `json:"status,omitempty"`
}

func (o ListEndpointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointsRequest", string(data)}, " ")
}

type ListEndpointsRequestStatus struct {
	value string
}

type ListEndpointsRequestStatusEnum struct {
	ACTIVE   ListEndpointsRequestStatus
	PENDING  ListEndpointsRequestStatus
	ERROR    ListEndpointsRequestStatus
	DELETING ListEndpointsRequestStatus
}

func GetListEndpointsRequestStatusEnum() ListEndpointsRequestStatusEnum {
	return ListEndpointsRequestStatusEnum{
		ACTIVE: ListEndpointsRequestStatus{
			value: "ACTIVE",
		},
		PENDING: ListEndpointsRequestStatus{
			value: "PENDING",
		},
		ERROR: ListEndpointsRequestStatus{
			value: "ERROR",
		},
		DELETING: ListEndpointsRequestStatus{
			value: "DELETING",
		},
	}
}

func (c ListEndpointsRequestStatus) Value() string {
	return c.value
}

func (c ListEndpointsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEndpointsRequestStatus) UnmarshalJSON(b []byte) error {
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
