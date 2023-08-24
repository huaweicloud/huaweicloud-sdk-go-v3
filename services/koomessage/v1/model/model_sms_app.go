package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SmsApp struct {

	// 应用ID，用于获取、修改应用的唯一标识。
	AppId *string `json:"app_id,omitempty"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 状态。
	Status *SmsAppStatus `json:"status,omitempty"`

	// 地域。
	Region *string `json:"region,omitempty"`

	// 创建时间，格式：yyyy-MM-dd'T'HH:mm:ss。
	CreateTime *string `json:"create_time,omitempty"`

	// 上行短信地址。
	UpLinkAddr *string `json:"up_link_addr,omitempty"`
}

func (o SmsApp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsApp struct{}"
	}

	return strings.Join([]string{"SmsApp", string(data)}, " ")
}

type SmsAppStatus struct {
	value string
}

type SmsAppStatusEnum struct {
	CREATED   SmsAppStatus
	SUSPENDED SmsAppStatus
	LAUNCHED  SmsAppStatus
}

func GetSmsAppStatusEnum() SmsAppStatusEnum {
	return SmsAppStatusEnum{
		CREATED: SmsAppStatus{
			value: "CREATED：待上线。应用暂未创建成功，请稍候",
		},
		SUSPENDED: SmsAppStatus{
			value: "SUSPENDED：暂停。无法发起业务请求。当客户所发短信内容触发业务违规，或客户申请退订短信业务时，运营经理会将客户短信应用暂停",
		},
		LAUNCHED: SmsAppStatus{
			value: "LAUNCHED：正常。应用添加成功，可以正常使用",
		},
	}
}

func (c SmsAppStatus) Value() string {
	return c.value
}

func (c SmsAppStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmsAppStatus) UnmarshalJSON(b []byte) error {
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
