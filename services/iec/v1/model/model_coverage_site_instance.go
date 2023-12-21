package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CoverageSiteInstance 站点覆盖策略
type CoverageSiteInstance struct {

	// 站点ID。 具体信息可通过调用“查询边缘站点列表”来查询。
	Site string `json:"site"`

	// 租户需求数量列表。表示租户发放资源站点的运营商和发放的资源组的数量。
	Demands []DemandInstance `json:"demands"`
}

func (o CoverageSiteInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CoverageSiteInstance struct{}"
	}

	return strings.Join([]string{"CoverageSiteInstance", string(data)}, " ")
}
