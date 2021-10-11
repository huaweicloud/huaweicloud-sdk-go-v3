package model

import (
	"encoding/json"

	"strings"
)

type MemberInfo struct {
	// 后端服务器地址  后端实例类型为ip时生效

	Host *string `json:"host,omitempty"`
	// 权重值。  允许您对云服务器进行评级，权重值越大，转发到该云服务的请求数量越多。权重只对加权轮询和加权最小连接算法生效  仅VPC通道类型为2时有效。

	Weight *int32 `json:"weight,omitempty"`
	// 后端云服务器的编号。  后端实例类型为instance时生效，支持英文，数字，“-”,“_”，1 ~ 64字符。

	EcsId *string `json:"ecs_id,omitempty"`
	// 后端云服务器的名称。  后端实例类型为instance时生效，支持汉字，英文，数字，“-”,“_”,“.”，1 ~ 64字符。

	EcsName *string `json:"ecs_name,omitempty"`
}

func (o MemberInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MemberInfo struct{}"
	}

	return strings.Join([]string{"MemberInfo", string(data)}, " ")
}
