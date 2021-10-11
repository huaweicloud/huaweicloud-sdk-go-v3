package model

import (
	"encoding/json"

	"strings"
)

// 创建路由请求结构体
type CreateRouterReqDto struct {
	// 路由ID，节点下唯一

	RouteId string `json:"route_id"`
	// sql參數

	Sql string `json:"sql"`
}

func (o CreateRouterReqDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRouterReqDto struct{}"
	}

	return strings.Join([]string{"CreateRouterReqDto", string(data)}, " ")
}
