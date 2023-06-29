package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModDeviceDto 修改的终端参数。
type ModDeviceDto struct {

	// 终端名称，建议为具体位置。
	Name *string `json:"name,omitempty"`

	// 投影码生成模式，默认为自动。 * 0：自动(该模式下根据消息上报的IP地址内部控制复杂度。   私网地址配置成简单模式，公网地址配置成复杂模式) * 1：简单 * 2：复杂
	PrjCodeMode *ModDeviceDtoPrjCodeMode `json:"prjCodeMode,omitempty"`

	// 部门编码，默认为根部门。 默认值：1。
	DeptCode *string `json:"deptCode,omitempty"`

	// 手机号，必须加上国家码。 例如中国大陆手机为“+86xxxxxxxxxxx”，当填写手机号时 “country”参数必填。 手机号只允许输入纯数字。 > 手机号或者邮箱至少填写一个。
	Phone *string `json:"phone,omitempty"`

	// [[手机号所属的国家](https://support.huaweicloud.com/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hws)[[手机号所属的国家](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hk) 。
	Country *string `json:"country,omitempty"`

	// 邮箱地址。
	Email *string `json:"email,omitempty"`

	// 终端描述。
	Description *string `json:"description,omitempty"`

	// 终端状态。 * 0：正常 * 1：冻结
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

func (c ModDeviceDtoPrjCodeMode) Value() int32 {
	return c.value
}

func (c ModDeviceDtoPrjCodeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModDeviceDtoPrjCodeMode) UnmarshalJSON(b []byte) error {
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

func (c ModDeviceDtoStatus) Value() int32 {
	return c.value
}

func (c ModDeviceDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModDeviceDtoStatus) UnmarshalJSON(b []byte) error {
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
