package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateDeviceRequestBody struct {

	// 父设备ID，无父设备时不填写，自动向下取整
	ParentDeviceId *int32 `json:"parent_device_id,omitempty"`

	Product *ProductReferer `json:"product"`

	// 设备密码，输入要求：至少1数字，1大写字母，1小写字母，1特殊字符(~!@#$%^&*()-_=+|[{}];:<>/?)，长度8-32个字符
	Password *string `json:"password,omitempty"`

	// 设备用户名，支持英文大小写、英文符号(-)及数字，长度10-50
	UserName *string `json:"user_name,omitempty"`

	// 设备名称，支持中文、中文标点符号（）。；，：“”、？《》及英文大小写、数字及英文符号()_,#.?'-@%&!, 长度2-64
	DeviceName string `json:"device_name"`

	// 设备物理编号，通常使用MAC或者IMEI号，支持英文大小写，数字，下划线和中划线，长度2-64
	NodeId string `json:"node_id"`

	// 应用ID
	AppId string `json:"app_id"`

	// 设备状态 0启用 1禁用，不填时默认为0启用
	Status *CreateDeviceRequestBodyStatus `json:"status,omitempty"`

	// 备注
	Description *string `json:"description,omitempty"`

	// 标签
	Tags *[]string `json:"tags,omitempty"`
}

func (o CreateDeviceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeviceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDeviceRequestBody", string(data)}, " ")
}

type CreateDeviceRequestBodyStatus struct {
	value int32
}

type CreateDeviceRequestBodyStatusEnum struct {
	E_0 CreateDeviceRequestBodyStatus
	E_1 CreateDeviceRequestBodyStatus
}

func GetCreateDeviceRequestBodyStatusEnum() CreateDeviceRequestBodyStatusEnum {
	return CreateDeviceRequestBodyStatusEnum{
		E_0: CreateDeviceRequestBodyStatus{
			value: 0,
		}, E_1: CreateDeviceRequestBodyStatus{
			value: 1,
		},
	}
}

func (c CreateDeviceRequestBodyStatus) Value() int32 {
	return c.value
}

func (c CreateDeviceRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDeviceRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
