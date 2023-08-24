package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEndpointConnectionsRequest Request Object
type ListEndpointConnectionsRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// 终端节点的ID，唯一标识
	Id *string `json:"id,omitempty"`

	// 终端节点的报文标识
	MarkerId *int32 `json:"marker_id,omitempty"`

	// 终端节点的连接状态 - pendingAcceptance 待接受 - accepted 已接受 - rejected 已拒绝 - failed 失败
	Status *ListEndpointConnectionsRequestStatus `json:"status,omitempty"`
}

func (o ListEndpointConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointConnectionsRequest", string(data)}, " ")
}

type ListEndpointConnectionsRequestStatus struct {
	value string
}

type ListEndpointConnectionsRequestStatusEnum struct {
	PENDING_ACCEPTANCE ListEndpointConnectionsRequestStatus
	ACCEPTED           ListEndpointConnectionsRequestStatus
	REJECTED           ListEndpointConnectionsRequestStatus
	FAILED             ListEndpointConnectionsRequestStatus
}

func GetListEndpointConnectionsRequestStatusEnum() ListEndpointConnectionsRequestStatusEnum {
	return ListEndpointConnectionsRequestStatusEnum{
		PENDING_ACCEPTANCE: ListEndpointConnectionsRequestStatus{
			value: "pendingAcceptance",
		},
		ACCEPTED: ListEndpointConnectionsRequestStatus{
			value: "accepted",
		},
		REJECTED: ListEndpointConnectionsRequestStatus{
			value: "rejected",
		},
		FAILED: ListEndpointConnectionsRequestStatus{
			value: "failed",
		},
	}
}

func (c ListEndpointConnectionsRequestStatus) Value() string {
	return c.value
}

func (c ListEndpointConnectionsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEndpointConnectionsRequestStatus) UnmarshalJSON(b []byte) error {
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
