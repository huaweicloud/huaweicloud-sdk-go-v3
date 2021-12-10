package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListTasksDetailsRequest struct {
	// 服务名称

	ServiceName string `json:"service_name"`
	// 目标服务作业对应的服务版本号

	ServiceVersion *string `json:"service_version,omitempty"`
	// 目标服务作业的状态

	State *ListTasksDetailsRequestState `json:"state,omitempty"`
	// 目标服务作业的名称，支持模糊匹配

	NameLike *string `json:"name_like,omitempty"`
	// 目标服务作业的ID，支持模糊匹配

	IdLike *string `json:"id_like,omitempty"`
	// 目标服务作业的创建起始时间

	CreatedSince *int64 `json:"created_since,omitempty"`
	// 目标服务作业的创建截止时间

	CreatedUntil *int64 `json:"created_until,omitempty"`
	// 展示服务作业时的排序字段和顺序

	Order *ListTasksDetailsRequestOrder `json:"order,omitempty"`
	// 首个展示的服务作业的偏移量

	Offset *int32 `json:"offset,omitempty"`
	// 展示服务作业的数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTasksDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListTasksDetailsRequest", string(data)}, " ")
}

type ListTasksDetailsRequestState struct {
	value string
}

type ListTasksDetailsRequestStateEnum struct {
	PENDING       ListTasksDetailsRequestState
	RECOVERING    ListTasksDetailsRequestState
	STARTING      ListTasksDetailsRequestState
	UPGRADING     ListTasksDetailsRequestState
	CREATE_FAILED ListTasksDetailsRequestState
	START_FAILED  ListTasksDetailsRequestState
	RUNNING       ListTasksDetailsRequestState
	STOPPING      ListTasksDetailsRequestState
	STOPPED       ListTasksDetailsRequestState
	ABNORMAL      ListTasksDetailsRequestState
	SUCCEEDED     ListTasksDetailsRequestState
	FAILED        ListTasksDetailsRequestState
	DELETING      ListTasksDetailsRequestState
	FREEZING      ListTasksDetailsRequestState
	FROZEN        ListTasksDetailsRequestState
}

func GetListTasksDetailsRequestStateEnum() ListTasksDetailsRequestStateEnum {
	return ListTasksDetailsRequestStateEnum{
		PENDING: ListTasksDetailsRequestState{
			value: "PENDING",
		},
		RECOVERING: ListTasksDetailsRequestState{
			value: "RECOVERING",
		},
		STARTING: ListTasksDetailsRequestState{
			value: "STARTING",
		},
		UPGRADING: ListTasksDetailsRequestState{
			value: "UPGRADING",
		},
		CREATE_FAILED: ListTasksDetailsRequestState{
			value: "CREATE_FAILED",
		},
		START_FAILED: ListTasksDetailsRequestState{
			value: "START_FAILED",
		},
		RUNNING: ListTasksDetailsRequestState{
			value: "RUNNING",
		},
		STOPPING: ListTasksDetailsRequestState{
			value: "STOPPING",
		},
		STOPPED: ListTasksDetailsRequestState{
			value: "STOPPED",
		},
		ABNORMAL: ListTasksDetailsRequestState{
			value: "ABNORMAL",
		},
		SUCCEEDED: ListTasksDetailsRequestState{
			value: "SUCCEEDED",
		},
		FAILED: ListTasksDetailsRequestState{
			value: "FAILED",
		},
		DELETING: ListTasksDetailsRequestState{
			value: "DELETING",
		},
		FREEZING: ListTasksDetailsRequestState{
			value: "FREEZING",
		},
		FROZEN: ListTasksDetailsRequestState{
			value: "FROZEN",
		},
	}
}

func (c ListTasksDetailsRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksDetailsRequestState) UnmarshalJSON(b []byte) error {
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

type ListTasksDetailsRequestOrder struct {
	value string
}

type ListTasksDetailsRequestOrderEnum struct {
	NAMEASC        ListTasksDetailsRequestOrder
	NAMEDESC       ListTasksDetailsRequestOrder
	CREATED_ATASC  ListTasksDetailsRequestOrder
	CREATED_ATDESC ListTasksDetailsRequestOrder
	UPDATED_ATASC  ListTasksDetailsRequestOrder
	UPDATED_ATDESC ListTasksDetailsRequestOrder
}

func GetListTasksDetailsRequestOrderEnum() ListTasksDetailsRequestOrderEnum {
	return ListTasksDetailsRequestOrderEnum{
		NAMEASC: ListTasksDetailsRequestOrder{
			value: "name:ASC",
		},
		NAMEDESC: ListTasksDetailsRequestOrder{
			value: "name:DESC",
		},
		CREATED_ATASC: ListTasksDetailsRequestOrder{
			value: "created_at:ASC",
		},
		CREATED_ATDESC: ListTasksDetailsRequestOrder{
			value: "created_at:DESC",
		},
		UPDATED_ATASC: ListTasksDetailsRequestOrder{
			value: "updated_at:ASC",
		},
		UPDATED_ATDESC: ListTasksDetailsRequestOrder{
			value: "updated_at:DESC",
		},
	}
}

func (c ListTasksDetailsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksDetailsRequestOrder) UnmarshalJSON(b []byte) error {
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
