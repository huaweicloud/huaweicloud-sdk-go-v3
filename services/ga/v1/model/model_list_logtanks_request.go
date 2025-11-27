package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListLogtanksRequest Request Object
type ListLogtanksRequest struct {

	// 分页查询每页的资源个数。如果不设置，则默认为500。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询的起始的资源ID，表示上一页最后一条查询资源记录的ID。不指定时表示查询第一页。 必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 资源ID。
	Id *string `json:"id,omitempty"`

	// 配置状态，可选范围: - ACTIVE：运行中 - PENDING：待定 - ERROR：错误 - DELETING：正在删除
	Status *ListLogtanksRequestStatus `json:"status,omitempty"`

	// 资源ID列表。
	ResourceIds *[]string `json:"resource_ids,omitempty"`

	// 云日志的资源类型。 取值范围： LISTENER：监听器
	ResourceType *ListLogtanksRequestResourceType `json:"resource_type,omitempty"`
}

func (o ListLogtanksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogtanksRequest struct{}"
	}

	return strings.Join([]string{"ListLogtanksRequest", string(data)}, " ")
}

type ListLogtanksRequestStatus struct {
	value string
}

type ListLogtanksRequestStatusEnum struct {
	ACTIVE   ListLogtanksRequestStatus
	PENDING  ListLogtanksRequestStatus
	ERROR    ListLogtanksRequestStatus
	DELETING ListLogtanksRequestStatus
}

func GetListLogtanksRequestStatusEnum() ListLogtanksRequestStatusEnum {
	return ListLogtanksRequestStatusEnum{
		ACTIVE: ListLogtanksRequestStatus{
			value: "ACTIVE",
		},
		PENDING: ListLogtanksRequestStatus{
			value: "PENDING",
		},
		ERROR: ListLogtanksRequestStatus{
			value: "ERROR",
		},
		DELETING: ListLogtanksRequestStatus{
			value: "DELETING",
		},
	}
}

func (c ListLogtanksRequestStatus) Value() string {
	return c.value
}

func (c ListLogtanksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLogtanksRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListLogtanksRequestResourceType struct {
	value string
}

type ListLogtanksRequestResourceTypeEnum struct {
	LISTENER ListLogtanksRequestResourceType
}

func GetListLogtanksRequestResourceTypeEnum() ListLogtanksRequestResourceTypeEnum {
	return ListLogtanksRequestResourceTypeEnum{
		LISTENER: ListLogtanksRequestResourceType{
			value: "LISTENER",
		},
	}
}

func (c ListLogtanksRequestResourceType) Value() string {
	return c.value
}

func (c ListLogtanksRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLogtanksRequestResourceType) UnmarshalJSON(b []byte) error {
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
