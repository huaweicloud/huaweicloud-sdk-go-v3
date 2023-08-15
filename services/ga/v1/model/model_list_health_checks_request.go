package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHealthChecksRequest Request Object
type ListHealthChecksRequest struct {

	// 分页查询每页的资源个数。如果不设置，则默认为500。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询的起始的资源ID，表示上一页最后一条查询资源记录的ID。不指定时表示查询第一页。 必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 资源ID。
	Id *string `json:"id,omitempty"`

	// 配置状态，可选范围: - ACTIVE：运行中 - PENDING：待定 - ERROR：错误 - DELETING：正在删除
	Status *ListHealthChecksRequestStatus `json:"status,omitempty"`

	// 终端节点组ID。
	EndpointGroupId *string `json:"endpoint_group_id,omitempty"`
}

func (o ListHealthChecksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthChecksRequest struct{}"
	}

	return strings.Join([]string{"ListHealthChecksRequest", string(data)}, " ")
}

type ListHealthChecksRequestStatus struct {
	value string
}

type ListHealthChecksRequestStatusEnum struct {
	ACTIVE   ListHealthChecksRequestStatus
	PENDING  ListHealthChecksRequestStatus
	ERROR    ListHealthChecksRequestStatus
	DELETING ListHealthChecksRequestStatus
}

func GetListHealthChecksRequestStatusEnum() ListHealthChecksRequestStatusEnum {
	return ListHealthChecksRequestStatusEnum{
		ACTIVE: ListHealthChecksRequestStatus{
			value: "ACTIVE",
		},
		PENDING: ListHealthChecksRequestStatus{
			value: "PENDING",
		},
		ERROR: ListHealthChecksRequestStatus{
			value: "ERROR",
		},
		DELETING: ListHealthChecksRequestStatus{
			value: "DELETING",
		},
	}
}

func (c ListHealthChecksRequestStatus) Value() string {
	return c.value
}

func (c ListHealthChecksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHealthChecksRequestStatus) UnmarshalJSON(b []byte) error {
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
