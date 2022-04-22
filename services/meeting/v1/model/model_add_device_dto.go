package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 终端信息。
type AddDeviceDto struct {

	// 终端名称，可以自定义，建议为具体位置，方便识别。 maxLength：64 minLength：1
	Name string `json:"name"`

	// 终端型号，枚举类型。当前支持TE系列和部分第三方硬件终端，具体的终端类型可以通过获取所有终端类型接口查询。 maxLength：128 minLength：1
	Model string `json:"model"`

	// 终端SN码，仅可包含数字、字母和下划线。 maxLength：30 minLength：1
	Sn *string `json:"sn,omitempty"`

	// 投影码生成模式，默认为自动 - 0、自动(该模式下根据消息上报的IP地址内部控制复杂度：   私网地址配置成简单模式；公网地址配置成复杂模式) - 1、简单 - 2、复杂
	PrjCodeMode *AddDeviceDtoPrjCodeMode `json:"prjCodeMode,omitempty"`

	// 部门编号，默认为根部门 默认值：1 maxLength：32 minLength：1
	DeptCode *string `json:"deptCode,omitempty"`

	// 手机号，必须加上国家码，例如中国大陆手机为“+86xxxxxxxxxxx”。当填写手机号时 “country”参数必填。 手机号只允许输入纯数字。 maxLength：32 minLength：0 说明： - 手机号或者邮箱至少填写一个。
	Phone *string `json:"phone,omitempty"`

	// 若smsNumber为手机号,则需带上手机号所属的国家。 例如国家为中国大陆则country参数取值为chinaPR 国家和国家码的对应关系请参考：https://support.huaweicloud.com/api-meeting/meeting_21_0109.html
	Country *string `json:"country,omitempty"`

	// 统一邮箱格式 maxLength：255 minLength：0
	Email *string `json:"email,omitempty"`

	// 终端描述 maxLength：128 minLength：0
	Description *string `json:"description,omitempty"`

	// 终端状态 * 0、正常 * 1、冻结\" 默认值：0。
	Status *AddDeviceDtoStatus `json:"status,omitempty"`

	// 是否发送邮件和短信通知。 * 0：不发送 * 不填或者其他值就发送
	SendNotify *string `json:"sendNotify,omitempty"`
}

func (o AddDeviceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDeviceDto struct{}"
	}

	return strings.Join([]string{"AddDeviceDto", string(data)}, " ")
}

type AddDeviceDtoPrjCodeMode struct {
	value int32
}

type AddDeviceDtoPrjCodeModeEnum struct {
	E_0 AddDeviceDtoPrjCodeMode
	E_1 AddDeviceDtoPrjCodeMode
	E_2 AddDeviceDtoPrjCodeMode
}

func GetAddDeviceDtoPrjCodeModeEnum() AddDeviceDtoPrjCodeModeEnum {
	return AddDeviceDtoPrjCodeModeEnum{
		E_0: AddDeviceDtoPrjCodeMode{
			value: 0,
		}, E_1: AddDeviceDtoPrjCodeMode{
			value: 1,
		}, E_2: AddDeviceDtoPrjCodeMode{
			value: 2,
		},
	}
}

func (c AddDeviceDtoPrjCodeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddDeviceDtoPrjCodeMode) UnmarshalJSON(b []byte) error {
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

type AddDeviceDtoStatus struct {
	value int32
}

type AddDeviceDtoStatusEnum struct {
	E_0 AddDeviceDtoStatus
	E_1 AddDeviceDtoStatus
}

func GetAddDeviceDtoStatusEnum() AddDeviceDtoStatusEnum {
	return AddDeviceDtoStatusEnum{
		E_0: AddDeviceDtoStatus{
			value: 0,
		}, E_1: AddDeviceDtoStatus{
			value: 1,
		},
	}
}

func (c AddDeviceDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddDeviceDtoStatus) UnmarshalJSON(b []byte) error {
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
