package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListFirewallUsingGetRequest struct {

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 每页显示个数
	Limit int32 `json:"limit"`

	// 服务类型 0 南北向防火墙 1 东西向防火墙
	ServiceType ListFirewallUsingGetRequestServiceType `json:"service_type"`
}

func (o ListFirewallUsingGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFirewallUsingGetRequest struct{}"
	}

	return strings.Join([]string{"ListFirewallUsingGetRequest", string(data)}, " ")
}

type ListFirewallUsingGetRequestServiceType struct {
	value int32
}

type ListFirewallUsingGetRequestServiceTypeEnum struct {
	E_0 ListFirewallUsingGetRequestServiceType
	E_1 ListFirewallUsingGetRequestServiceType
}

func GetListFirewallUsingGetRequestServiceTypeEnum() ListFirewallUsingGetRequestServiceTypeEnum {
	return ListFirewallUsingGetRequestServiceTypeEnum{
		E_0: ListFirewallUsingGetRequestServiceType{
			value: 0,
		}, E_1: ListFirewallUsingGetRequestServiceType{
			value: 1,
		},
	}
}

func (c ListFirewallUsingGetRequestServiceType) Value() int32 {
	return c.value
}

func (c ListFirewallUsingGetRequestServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFirewallUsingGetRequestServiceType) UnmarshalJSON(b []byte) error {
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
