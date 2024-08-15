package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateListenerIpGroupOption 监听器对象中的控制组（ipgroup）信息，可以不传或传null或{}，表示监听器不绑定访问控制组。 若需要绑定访问控制组，则ipgroup_id是必须的。  [不支持该字段，请勿使用。](tag:hcso_dt)
type CreateListenerIpGroupOption struct {

	// 参数解释：监听器关联的访问控制组的id。  约束限制： - 当关联的ipgroup中的ip_list为[]，且类型为白名单时，表示禁止所有ip的访问。 - 当关联的ipgroup中的ip_list为[]，且类型为黑名单时，表示允许所有ip的访问。
	IpgroupId string `json:"ipgroup_id"`

	// 参数解释：访问控制组的状态。  取值范围： - true：开启访问控制，默认值。 - flase：关闭访问控制。
	EnableIpgroup *bool `json:"enable_ipgroup,omitempty"`

	// 参数解释：访问控制组的类型。  取值范围： - white：白名单，只允许指定ip访问。 - black：黑名单，不允许指定ip访问。  默认取值：white
	Type *CreateListenerIpGroupOptionType `json:"type,omitempty"`
}

func (o CreateListenerIpGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerIpGroupOption struct{}"
	}

	return strings.Join([]string{"CreateListenerIpGroupOption", string(data)}, " ")
}

type CreateListenerIpGroupOptionType struct {
	value string
}

type CreateListenerIpGroupOptionTypeEnum struct {
	WHITE CreateListenerIpGroupOptionType
	BLACK CreateListenerIpGroupOptionType
}

func GetCreateListenerIpGroupOptionTypeEnum() CreateListenerIpGroupOptionTypeEnum {
	return CreateListenerIpGroupOptionTypeEnum{
		WHITE: CreateListenerIpGroupOptionType{
			value: "white",
		},
		BLACK: CreateListenerIpGroupOptionType{
			value: "black",
		},
	}
}

func (c CreateListenerIpGroupOptionType) Value() string {
	return c.value
}

func (c CreateListenerIpGroupOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateListenerIpGroupOptionType) UnmarshalJSON(b []byte) error {
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
