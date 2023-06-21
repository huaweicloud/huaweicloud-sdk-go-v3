package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 增加OpenTSDB节点的请求体，其中节点个数为指定增加TSD节点的个数，如请求体的node_num为2，那么会扩容两个tsd节点
type AddComponentReq struct {

	// 节点个数, 范围是[2,10]
	NodeNum int32 `json:"node_num"`
}

func (o AddComponentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddComponentReq struct{}"
	}

	return strings.Join([]string{"AddComponentReq", string(data)}, " ")
}
