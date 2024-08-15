package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateListenerIpGroupOption listener对象中的ipgroup信息
type UpdateListenerIpGroupOption struct {

	// 参数解释：监听器关联的访问控制组的id。 创建时必选，更新时非必选。  约束限制：指定的ipgroup必须已存在，不能指定为null，否则与enable_ipgroup冲突。
	IpgroupId *string `json:"ipgroup_id,omitempty"`

	// 参数解释：访问控制组的状态。 开启访问控制的监听器，允许直接删除。  取值范围： - true:开启访问控制。 - flase：关闭访问控制。
	EnableIpgroup *bool `json:"enable_ipgroup,omitempty"`

	// 参数解释：访问控制组的类型。  取值范围： - white:白名单，只允许指定ip访问。 - black:黑名单，不允许指定ip访问。
	Type *UpdateListenerIpGroupOptionType `json:"type,omitempty"`
}

func (o UpdateListenerIpGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerIpGroupOption struct{}"
	}

	return strings.Join([]string{"UpdateListenerIpGroupOption", string(data)}, " ")
}

type UpdateListenerIpGroupOptionType struct {
	value string
}

type UpdateListenerIpGroupOptionTypeEnum struct {
	WHITE UpdateListenerIpGroupOptionType
	BLACK UpdateListenerIpGroupOptionType
}

func GetUpdateListenerIpGroupOptionTypeEnum() UpdateListenerIpGroupOptionTypeEnum {
	return UpdateListenerIpGroupOptionTypeEnum{
		WHITE: UpdateListenerIpGroupOptionType{
			value: "white",
		},
		BLACK: UpdateListenerIpGroupOptionType{
			value: "black",
		},
	}
}

func (c UpdateListenerIpGroupOptionType) Value() string {
	return c.value
}

func (c UpdateListenerIpGroupOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateListenerIpGroupOptionType) UnmarshalJSON(b []byte) error {
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
