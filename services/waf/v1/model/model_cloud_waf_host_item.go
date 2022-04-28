package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CloudWafHostItem struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 域名id
	Hostid *string `json:"hostid,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// WAF部署模式
	Type *int32 `json:"type,omitempty"`

	// 是否开启了代理
	Proxy *bool `json:"proxy,omitempty"`

	Flag *Flag `json:"flag,omitempty"`

	// 创建的云模式防护域名
	Hostname *string `json:"hostname,omitempty"`

	// cname前缀
	AccessCode *string `json:"access_code,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 创建防护域名的时间
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	// 接入状态
	AccessStatus *int32 `json:"access_status,omitempty"`

	// 是否使用独享ip
	ExclusiveIp *bool `json:"exclusive_ip,omitempty"`

	// 付费模式，目前只支持prePaid预付款模式
	PaidType *CloudWafHostItemPaidType `json:"paid_type,omitempty"`
}

func (o CloudWafHostItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudWafHostItem struct{}"
	}

	return strings.Join([]string{"CloudWafHostItem", string(data)}, " ")
}

type CloudWafHostItemPaidType struct {
	value string
}

type CloudWafHostItemPaidTypeEnum struct {
	PRE_PAID CloudWafHostItemPaidType
}

func GetCloudWafHostItemPaidTypeEnum() CloudWafHostItemPaidTypeEnum {
	return CloudWafHostItemPaidTypeEnum{
		PRE_PAID: CloudWafHostItemPaidType{
			value: "prePaid",
		},
	}
}

func (c CloudWafHostItemPaidType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudWafHostItemPaidType) UnmarshalJSON(b []byte) error {
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
