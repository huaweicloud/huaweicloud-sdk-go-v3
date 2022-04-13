package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListEndpointServiceRequest struct {
	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。

	ContentType string `json:"Content-Type"`
	// 终端节点服务的名称，支持大小写，前后模糊匹配。

	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`
	// 终端节点服务的ID，唯一标识。

	Id *string `json:"id,omitempty"`
	// 终端节点服务的状态。 ● creating：创建中 ● available：可连接 ● failed：失败 ● deleting：删除中

	Status *ListEndpointServiceRequestStatus `json:"status,omitempty"`
	// 查询结果中终端节点服务列表的排序字段，取值为： ● create_at：终端节点服务的创建时间 ● update_at：终端节点服务的更新时间 默认值为create_at。

	SortKey *ListEndpointServiceRequestSortKey `json:"sort_key,omitempty"`
	// 查询结果中终端节点服务列表的排 序方式，取值为： ● desc：降序排序 ● asc：升序排序 默认值为desc。

	SortDir *ListEndpointServiceRequestSortDir `json:"sort_dir,omitempty"`
	// 查询返回的终端节点服务数量限制，即每页返回的终端节点服务的个数。 取值范围：0~1000，取值一般为10，20或者50，默认为10。

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量。 偏移量为一个大于0小于终端节点服务总个数的整数，表示从偏移量后面的终端节点服务开始查询。

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListEndpointServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointServiceRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointServiceRequest", string(data)}, " ")
}

type ListEndpointServiceRequestStatus struct {
	value string
}

type ListEndpointServiceRequestStatusEnum struct {
	CREATING  ListEndpointServiceRequestStatus
	AVAILABLE ListEndpointServiceRequestStatus
	FAILED    ListEndpointServiceRequestStatus
	DELETING  ListEndpointServiceRequestStatus
}

func GetListEndpointServiceRequestStatusEnum() ListEndpointServiceRequestStatusEnum {
	return ListEndpointServiceRequestStatusEnum{
		CREATING: ListEndpointServiceRequestStatus{
			value: "creating",
		},
		AVAILABLE: ListEndpointServiceRequestStatus{
			value: "available",
		},
		FAILED: ListEndpointServiceRequestStatus{
			value: "failed",
		},
		DELETING: ListEndpointServiceRequestStatus{
			value: "deleting",
		},
	}
}

func (c ListEndpointServiceRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEndpointServiceRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListEndpointServiceRequestSortKey struct {
	value string
}

type ListEndpointServiceRequestSortKeyEnum struct {
	CREATE_AT ListEndpointServiceRequestSortKey
	UPDATE_AT ListEndpointServiceRequestSortKey
}

func GetListEndpointServiceRequestSortKeyEnum() ListEndpointServiceRequestSortKeyEnum {
	return ListEndpointServiceRequestSortKeyEnum{
		CREATE_AT: ListEndpointServiceRequestSortKey{
			value: "create_at",
		},
		UPDATE_AT: ListEndpointServiceRequestSortKey{
			value: "update_at",
		},
	}
}

func (c ListEndpointServiceRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEndpointServiceRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListEndpointServiceRequestSortDir struct {
	value string
}

type ListEndpointServiceRequestSortDirEnum struct {
	ASC  ListEndpointServiceRequestSortDir
	DESC ListEndpointServiceRequestSortDir
}

func GetListEndpointServiceRequestSortDirEnum() ListEndpointServiceRequestSortDirEnum {
	return ListEndpointServiceRequestSortDirEnum{
		ASC: ListEndpointServiceRequestSortDir{
			value: "asc",
		},
		DESC: ListEndpointServiceRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListEndpointServiceRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEndpointServiceRequestSortDir) UnmarshalJSON(b []byte) error {
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
