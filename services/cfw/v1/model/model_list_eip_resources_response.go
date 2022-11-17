package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ListEipResourcesResponse struct {

	// 弹性公网ID
	Id *string `json:"id,omitempty"`

	// 弹性公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// EIP防护状态
	Status *ListEipResourcesResponseStatus `json:"status,omitempty"`

	// 弹性公网IP,IPV6
	PublicIpv6 *string `json:"public_ipv6,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 设备id
	DeviceId *string `json:"device_id,omitempty"`

	// 设备名称
	DeviceName *string `json:"device_name,omitempty"`

	// 设备拥有者
	DeviceOwner *string `json:"device_owner,omitempty"`

	// 关联实例类型
	AssociateInstanceType *string `json:"associate_instance_type,omitempty"`
	HttpStatusCode        int     `json:"-"`
}

func (o ListEipResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEipResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListEipResourcesResponse", string(data)}, " ")
}

type ListEipResourcesResponseStatus struct {
	value int32
}

type ListEipResourcesResponseStatusEnum struct {
	E_0 ListEipResourcesResponseStatus
	E_1 ListEipResourcesResponseStatus
}

func GetListEipResourcesResponseStatusEnum() ListEipResourcesResponseStatusEnum {
	return ListEipResourcesResponseStatusEnum{
		E_0: ListEipResourcesResponseStatus{
			value: 0,
		}, E_1: ListEipResourcesResponseStatus{
			value: 1,
		},
	}
}

func (c ListEipResourcesResponseStatus) Value() int32 {
	return c.value
}

func (c ListEipResourcesResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEipResourcesResponseStatus) UnmarshalJSON(b []byte) error {
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
