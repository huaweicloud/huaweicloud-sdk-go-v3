package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 负载均衡器规格信息。
type Flavor struct {
	// 规格ID。

	Id string `json:"id"`

	Info *FlavorInfo `json:"info"`
	// 规格名称。

	Name string `json:"name"`
	// 共享。

	Shared bool `json:"shared"`
	// 项目ID。

	ProjectId string `json:"project_id"`
	// L4和L7 分别表示四层和七层flavor。查询支持按type过滤。

	Type string `json:"type"`
	// 是否售罄。

	FlavorSoldOut bool `json:"flavor_sold_out"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
