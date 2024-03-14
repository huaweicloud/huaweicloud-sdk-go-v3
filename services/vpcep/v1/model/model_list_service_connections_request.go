package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListServiceConnectionsRequest Request Object
type ListServiceConnectionsRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json， 如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType *string `json:"Content-Type,omitempty"`

	// 终端节点服务的ID。
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	// 终端节点的ID，唯一标识。
	Id *string `json:"id,omitempty"`

	// 终端节点的报文标识。
	MarkerId *string `json:"marker_id,omitempty"`

	// 终端节点的连接状态。  - pendingAcceptance：待接受  - accepted：已接受  - rejected：已拒绝  - failed：失败
	Status *ListServiceConnectionsRequestStatus `json:"status,omitempty"`

	// 查询结果中终端节点列表的排序字段，取值为：  - create_at：终端节点的创建时间  - update_at：终端节点的更新时间 默认值为create_at。
	SortKey *ListServiceConnectionsRequestSortKey `json:"sort_key,omitempty"`

	// 查询结果中终端节点列表的排序方式，取值为：  - desc：降序排序  - asc：升序排序 默认值为desc。
	SortDir *ListServiceConnectionsRequestSortDir `json:"sort_dir,omitempty"`

	// 查询返回终端节点服务的连接列表限制每页个数，即每页返回的个数。 取值范围：0~1000，取值一般为10，20或者50，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量。 偏移量为一个大于0小于终端节点服务总个数的整数， 表示从偏移量后面的终端节点服务开始查询。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListServiceConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListServiceConnectionsRequest", string(data)}, " ")
}

type ListServiceConnectionsRequestStatus struct {
	value string
}

type ListServiceConnectionsRequestStatusEnum struct {
	PENDING_ACCEPTANCE ListServiceConnectionsRequestStatus
	ACCEPTED           ListServiceConnectionsRequestStatus
	REJECTED           ListServiceConnectionsRequestStatus
	FAILED             ListServiceConnectionsRequestStatus
}

func GetListServiceConnectionsRequestStatusEnum() ListServiceConnectionsRequestStatusEnum {
	return ListServiceConnectionsRequestStatusEnum{
		PENDING_ACCEPTANCE: ListServiceConnectionsRequestStatus{
			value: "pendingAcceptance",
		},
		ACCEPTED: ListServiceConnectionsRequestStatus{
			value: "accepted",
		},
		REJECTED: ListServiceConnectionsRequestStatus{
			value: "rejected",
		},
		FAILED: ListServiceConnectionsRequestStatus{
			value: "failed",
		},
	}
}

func (c ListServiceConnectionsRequestStatus) Value() string {
	return c.value
}

func (c ListServiceConnectionsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceConnectionsRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListServiceConnectionsRequestSortKey struct {
	value string
}

type ListServiceConnectionsRequestSortKeyEnum struct {
	CREATE_AT ListServiceConnectionsRequestSortKey
	UPDATE_AT ListServiceConnectionsRequestSortKey
}

func GetListServiceConnectionsRequestSortKeyEnum() ListServiceConnectionsRequestSortKeyEnum {
	return ListServiceConnectionsRequestSortKeyEnum{
		CREATE_AT: ListServiceConnectionsRequestSortKey{
			value: "create_at",
		},
		UPDATE_AT: ListServiceConnectionsRequestSortKey{
			value: "update_at",
		},
	}
}

func (c ListServiceConnectionsRequestSortKey) Value() string {
	return c.value
}

func (c ListServiceConnectionsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceConnectionsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListServiceConnectionsRequestSortDir struct {
	value string
}

type ListServiceConnectionsRequestSortDirEnum struct {
	ASC  ListServiceConnectionsRequestSortDir
	DESC ListServiceConnectionsRequestSortDir
}

func GetListServiceConnectionsRequestSortDirEnum() ListServiceConnectionsRequestSortDirEnum {
	return ListServiceConnectionsRequestSortDirEnum{
		ASC: ListServiceConnectionsRequestSortDir{
			value: "asc",
		},
		DESC: ListServiceConnectionsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListServiceConnectionsRequestSortDir) Value() string {
	return c.value
}

func (c ListServiceConnectionsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceConnectionsRequestSortDir) UnmarshalJSON(b []byte) error {
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
