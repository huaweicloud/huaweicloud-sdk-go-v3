package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 站点覆盖策略
type CoverageSite struct {
	// 站点名称。 具体信息可通过调用“查询边缘站点列表”来查询(注意：本字段区分大小写)。

	Site string `json:"site"`
	// 租户需求数量列表。表示租户发放资源站点的运营商和发放的资源组的数量。

	Demands []Demand `json:"demands"`
}

func (o CoverageSite) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CoverageSite struct{}"
	}

	return strings.Join([]string{"CoverageSite", string(data)}, " ")
}
