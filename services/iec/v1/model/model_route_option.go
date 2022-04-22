package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 路由对象
type RouteOption struct {

	// 路由的类型  取值范围：     1）ecs：弹性云服务器     2）vip：虚拟IP
	Type RouteOptionType `json:"type"`

	// 路由的目的网段  约束：合法的CIDR格式, 目的地址不可更新
	Destination string `json:"destination"`

	// 路由下一跳对象的ID  取值范围：     1）当type为ecs时，传入ecs实例ID；     2）当type为vip时，取值为vip的ip地址；
	Nexthop string `json:"nexthop"`

	// 路由的描述信息  取值范围：0-255个字符，不能包含“<”和“>”
	Description *string `json:"description,omitempty"`
}

func (o RouteOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouteOption struct{}"
	}

	return strings.Join([]string{"RouteOption", string(data)}, " ")
}

type RouteOptionType struct {
	value string
}

type RouteOptionTypeEnum struct {
	ECS RouteOptionType
	VIP RouteOptionType
}

func GetRouteOptionTypeEnum() RouteOptionTypeEnum {
	return RouteOptionTypeEnum{
		ECS: RouteOptionType{
			value: "ecs",
		},
		VIP: RouteOptionType{
			value: "vip",
		},
	}
}

func (c RouteOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RouteOptionType) UnmarshalJSON(b []byte) error {
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
