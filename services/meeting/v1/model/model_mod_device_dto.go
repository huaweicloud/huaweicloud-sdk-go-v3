package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 修改终端的参数DTO对象
type ModDeviceDto struct {
	// 终端名称，建议为具体位置。 maxLength：64 minLength：0

	Name *string `json:"name,omitempty"`
	// 投影码生成模式，默认为自动 - 0、自动(该模式下根据消息上报的IP地址内部控制复杂度：   私网地址配置成简单模式；公网地址配置成复杂模式) - 1、简单 - 2、复杂

	PrjCodeMode *ModDeviceDtoPrjCodeMode `json:"prjCodeMode,omitempty"`
	// 部门编号，默认为根部门 默认值：1 maxLength：32 minLength：0

	DeptCode *string `json:"deptCode,omitempty"`
	// 手机号，必须加上国家码。 例如中国大陆手机为“+86xxxxxxxxxxx”，当填写手机号时 “country”参数必填。 手机号只允许输入纯数字。 maxLength：32 minLength：0 说明： - 手机号或者邮箱至少填写一个。

	Phone *string `json:"phone,omitempty"`
	// 若smsNumber为手机号,则需带上手机号所属的国家。 例如国家为中国大陆则country参数取值为chinaPR 国家和国家码的对应关系请参考：https://support.huaweicloud.com/api-meeting/meeting_21_0109.html

	Country *string `json:"country,omitempty"`
	// 统一邮箱格式 maxLength：255 minLength：0

	Email *string `json:"email,omitempty"`
	// 终端描述 maxLength：128 minLength：0

	Description *string `json:"description,omitempty"`
	// 终端状态。 * 0、正常 * 1、冻结

	Status *ModDeviceDtoStatus `json:"status,omitempty"`
}

func (o ModDeviceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModDeviceDto struct{}"
	}

	return strings.Join([]string{"ModDeviceDto", string(data)}, " ")
}

type ModDeviceDtoPrjCodeMode struct {
	value int32
}

type ModDeviceDtoPrjCodeModeEnum struct {
	E_0 ModDeviceDtoPrjCodeMode
	E_1 ModDeviceDtoPrjCodeMode
	E_2 ModDeviceDtoPrjCodeMode
}

func GetModDeviceDtoPrjCodeModeEnum() ModDeviceDtoPrjCodeModeEnum {
	return ModDeviceDtoPrjCodeModeEnum{
		E_0: ModDeviceDtoPrjCodeMode{
			value: 0,
		}, E_1: ModDeviceDtoPrjCodeMode{
			value: 1,
		}, E_2: ModDeviceDtoPrjCodeMode{
			value: 2,
		},
	}
}

func (c ModDeviceDtoPrjCodeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModDeviceDtoPrjCodeMode) UnmarshalJSON(b []byte) error {
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

type ModDeviceDtoStatus struct {
	value int32
}

type ModDeviceDtoStatusEnum struct {
	E_0 ModDeviceDtoStatus
	E_1 ModDeviceDtoStatus
}

func GetModDeviceDtoStatusEnum() ModDeviceDtoStatusEnum {
	return ModDeviceDtoStatusEnum{
		E_0: ModDeviceDtoStatus{
			value: 0,
		}, E_1: ModDeviceDtoStatus{
			value: 1,
		},
	}
}

func (c ModDeviceDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModDeviceDtoStatus) UnmarshalJSON(b []byte) error {
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
