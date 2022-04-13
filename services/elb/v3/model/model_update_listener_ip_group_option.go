package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// listener对象中的ipgroup信息
type UpdateListenerIpGroupOption struct {
	// 监听器关联的访问控制组的id。 创建时必选，更新时非必选。 指定的ipgroup必须已存在，不能指定为null，否则与enable_ipgroup冲突。

	IpgroupId *string `json:"ipgroup_id,omitempty"`
	// 访问控制组的状态。 True:开启访问控制； False：关闭访问控制； 开启访问控制的监听器，允许直接删除。

	EnableIpgroup *bool `json:"enable_ipgroup,omitempty"`
	// 访问控制组的类型。 white:白名单，只允许指定ip访问； black:黑名单，不允许指定ip访问；

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

func (c UpdateListenerIpGroupOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateListenerIpGroupOptionType) UnmarshalJSON(b []byte) error {
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
