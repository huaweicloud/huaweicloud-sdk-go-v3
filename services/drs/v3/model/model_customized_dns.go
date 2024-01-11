package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CustomizedDns 客户自定义DNS服务。
type CustomizedDns struct {

	// 是否设置客户自定义DNS。
	IsSetDns bool `json:"is_set_dns"`

	// 设置客户自定义DNS的行为。 - add：新增客户自定义DNS IP。 - keep：保持客户自定义DNS IP。 - update：更新客户自定义DNS IP（当DNS IP变化时更新生效）。 - recover：还原系统默认DNS IP（还原时可能会导致域名解析失败，请谨慎操作）。
	SetDnsAction CustomizedDnsSetDnsAction `json:"set_dns_action"`

	// 设置客户自定义DNS IP。
	DnsIp string `json:"dns_ip"`
}

func (o CustomizedDns) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizedDns struct{}"
	}

	return strings.Join([]string{"CustomizedDns", string(data)}, " ")
}

type CustomizedDnsSetDnsAction struct {
	value string
}

type CustomizedDnsSetDnsActionEnum struct {
	ADD     CustomizedDnsSetDnsAction
	KEEP    CustomizedDnsSetDnsAction
	UPDATE  CustomizedDnsSetDnsAction
	RECOVER CustomizedDnsSetDnsAction
}

func GetCustomizedDnsSetDnsActionEnum() CustomizedDnsSetDnsActionEnum {
	return CustomizedDnsSetDnsActionEnum{
		ADD: CustomizedDnsSetDnsAction{
			value: "add",
		},
		KEEP: CustomizedDnsSetDnsAction{
			value: "keep",
		},
		UPDATE: CustomizedDnsSetDnsAction{
			value: "update",
		},
		RECOVER: CustomizedDnsSetDnsAction{
			value: "recover",
		},
	}
}

func (c CustomizedDnsSetDnsAction) Value() string {
	return c.value
}

func (c CustomizedDnsSetDnsAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomizedDnsSetDnsAction) UnmarshalJSON(b []byte) error {
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
