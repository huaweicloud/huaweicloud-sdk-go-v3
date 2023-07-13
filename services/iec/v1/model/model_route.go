package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Route 路由对象
type Route struct {

	// 路由的类型  取值范围：     1）ecs：弹性云服务器     2）vip：虚拟IP     3）local：系统路由，不可修改和删除
	Type *RouteType `json:"type,omitempty"`

	// 路由的目的网段  约束：合法的CIDR格式
	Destination *string `json:"destination,omitempty"`

	// 路由下一跳对象的ID  取值范围：     1）当type为ecs时，传入ecs实例ID；     2）当type为vip时，取值为vip对应的IP地址；
	Nexthop *string `json:"nexthop,omitempty"`

	// 路由的描述信息  取值范围：0-255个字符，不能包含“<”和“>”
	Description *string `json:"description,omitempty"`
}

func (o Route) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Route struct{}"
	}

	return strings.Join([]string{"Route", string(data)}, " ")
}

type RouteType struct {
	value string
}

type RouteTypeEnum struct {
	ECS   RouteType
	VIP   RouteType
	LOCAL RouteType
}

func GetRouteTypeEnum() RouteTypeEnum {
	return RouteTypeEnum{
		ECS: RouteType{
			value: "ecs",
		},
		VIP: RouteType{
			value: "vip",
		},
		LOCAL: RouteType{
			value: "local",
		},
	}
}

func (c RouteType) Value() string {
	return c.value
}

func (c RouteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RouteType) UnmarshalJSON(b []byte) error {
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
