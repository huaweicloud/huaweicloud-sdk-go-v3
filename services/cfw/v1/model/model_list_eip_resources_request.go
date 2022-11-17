package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListEipResourcesRequest struct {

	// 防护对象ID
	ObjectId string `json:"object_id"`

	// 弹性公网ID/弹性公网IP
	KeyWord *string `json:"key_word,omitempty"`

	// 防护状态 null-全部 0-开启防护 1-关闭防护
	Status *ListEipResourcesRequestStatus `json:"status,omitempty"`

	// 是否同步租户EIP数据 0-不同步 1-同步
	Sync *ListEipResourcesRequestSync `json:"sync,omitempty"`

	// 每页显示个数
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 设备键
	DeviceKey *string `json:"device_key,omitempty"`

	// 地址类型0 ipv4,1 ipv6
	AddressType *int32 `json:"address_type,omitempty"`
}

func (o ListEipResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEipResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListEipResourcesRequest", string(data)}, " ")
}

type ListEipResourcesRequestStatus struct {
	value string
}

type ListEipResourcesRequestStatusEnum struct {
	NULL ListEipResourcesRequestStatus
	E_0  ListEipResourcesRequestStatus
	E_1  ListEipResourcesRequestStatus
}

func GetListEipResourcesRequestStatusEnum() ListEipResourcesRequestStatusEnum {
	return ListEipResourcesRequestStatusEnum{
		NULL: ListEipResourcesRequestStatus{
			value: "null",
		},
		E_0: ListEipResourcesRequestStatus{
			value: "0",
		},
		E_1: ListEipResourcesRequestStatus{
			value: "1",
		},
	}
}

func (c ListEipResourcesRequestStatus) Value() string {
	return c.value
}

func (c ListEipResourcesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEipResourcesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListEipResourcesRequestSync struct {
	value int32
}

type ListEipResourcesRequestSyncEnum struct {
	E_0 ListEipResourcesRequestSync
	E_1 ListEipResourcesRequestSync
}

func GetListEipResourcesRequestSyncEnum() ListEipResourcesRequestSyncEnum {
	return ListEipResourcesRequestSyncEnum{
		E_0: ListEipResourcesRequestSync{
			value: 0,
		}, E_1: ListEipResourcesRequestSync{
			value: 1,
		},
	}
}

func (c ListEipResourcesRequestSync) Value() int32 {
	return c.value
}

func (c ListEipResourcesRequestSync) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEipResourcesRequestSync) UnmarshalJSON(b []byte) error {
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
