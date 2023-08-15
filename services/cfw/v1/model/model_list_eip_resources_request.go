package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEipResourcesRequest Request Object
type ListEipResourcesRequest struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用查询防火墙实例接口获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。具体可参考APIExlorer和帮助中心FAQ。
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

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 设备键
	DeviceKey *string `json:"device_key,omitempty"`

	// 地址类型0 ipv4,1 ipv6
	AddressType *int32 `json:"address_type,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用查询防火墙实例接口获得。具体可参考APIExlorer和帮助中心FAQ。默认情况下，fw_instance_Id为空时，返回帐号下第一个墙的信息；fw_instance_Id非空时，返回与fw_instance_Id对应墙的信息。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 所绑定防火墙id防火墙名称
	FwKeyWord *string `json:"fw_key_word,omitempty"`

	// 弹性公网ip的企业项目id
	EpsId *string `json:"eps_id,omitempty"`
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
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
