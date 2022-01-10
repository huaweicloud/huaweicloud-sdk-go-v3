package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点数据源请求内容
type NodeContentReq struct {
	// 节点实例ID

	SiteId string `json:"site_id"`
	// SQL列表，将指定边缘平台节点的数字孪生模型实例数据转发到中心平台节点。

	Sqllist []string `json:"sqllist"`
}

func (o NodeContentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeContentReq struct{}"
	}

	return strings.Join([]string{"NodeContentReq", string(data)}, " ")
}
