package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRouteRequestBody struct {

	// 操作类型。add_ip为增加集群路由，del_ip为删除集群路由。
	Configtype string `json:"configtype"`

	// 路由ip地址，公网源数据所在的服务器ip。不能以0开头。
	Configkey string `json:"configkey"`

	// 路由子网掩码。如果上面ip取的是16位，子网掩码可以填255.255.0.0，24位的话子网掩码填255.255.255.0
	Configvalue string `json:"configvalue"`
}

func (o UpdateRouteRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRouteRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRouteRequestBody", string(data)}, " ")
}
