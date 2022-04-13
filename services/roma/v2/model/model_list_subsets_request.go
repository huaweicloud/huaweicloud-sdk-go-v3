package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListSubsetsRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 设备ID

	DeviceId int32 `json:"device_id"`
	// 每页显示条目数量，最大数量999，超过999后只返回999

	Limit *int32 `json:"limit,omitempty"`
	// 设备名称

	DeviceName *string `json:"device_name,omitempty"`
	// 设备状态 0-启动 1-停用

	Status *ListSubsetsRequestStatus `json:"status,omitempty"`
	// 设备状态 0-未连接 1-在线 2-离线

	OnlineStatus *ListSubsetsRequestOnlineStatus `json:"online_status,omitempty"`
	// 偏移量，表示从此偏移量开始查询， offset大于等于0

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListSubsetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubsetsRequest struct{}"
	}

	return strings.Join([]string{"ListSubsetsRequest", string(data)}, " ")
}

type ListSubsetsRequestStatus struct {
	value int32
}

type ListSubsetsRequestStatusEnum struct {
	E_0 ListSubsetsRequestStatus
	E_1 ListSubsetsRequestStatus
}

func GetListSubsetsRequestStatusEnum() ListSubsetsRequestStatusEnum {
	return ListSubsetsRequestStatusEnum{
		E_0: ListSubsetsRequestStatus{
			value: 0,
		}, E_1: ListSubsetsRequestStatus{
			value: 1,
		},
	}
}

func (c ListSubsetsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSubsetsRequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ListSubsetsRequestOnlineStatus struct {
	value int32
}

type ListSubsetsRequestOnlineStatusEnum struct {
	E_0 ListSubsetsRequestOnlineStatus
	E_1 ListSubsetsRequestOnlineStatus
	E_2 ListSubsetsRequestOnlineStatus
}

func GetListSubsetsRequestOnlineStatusEnum() ListSubsetsRequestOnlineStatusEnum {
	return ListSubsetsRequestOnlineStatusEnum{
		E_0: ListSubsetsRequestOnlineStatus{
			value: 0,
		}, E_1: ListSubsetsRequestOnlineStatus{
			value: 1,
		}, E_2: ListSubsetsRequestOnlineStatus{
			value: 2,
		},
	}
}

func (c ListSubsetsRequestOnlineStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSubsetsRequestOnlineStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
