package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RouterRespDto 创建路由请求结构体
type RouterRespDto struct {

	// 路由ID，节点下唯一
	RouteId string `json:"route_id"`

	// sql參數
	Sql *string `json:"sql,omitempty"`
}

func (o RouterRespDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouterRespDto struct{}"
	}

	return strings.Join([]string{"RouterRespDto", string(data)}, " ")
}
