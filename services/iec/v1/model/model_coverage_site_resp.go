package model

import (
	"encoding/json"

	"strings"
)

// 站点覆盖策略
type CoverageSiteResp struct {
	// 站点名称。 具体信息可通过调用“查询边缘站点列表”来查询(注意：本字段区分大小写)。

	Site string `json:"site"`
	// 租户需求数量列表。表示租户发放资源站点的运营商和发放的资源组的数量。

	Demands []DemandResp `json:"demands"`
	// 覆盖区域的国际化信息。

	I18nSite *string `json:"i18n_site,omitempty"`
}

func (o CoverageSiteResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CoverageSiteResp struct{}"
	}

	return strings.Join([]string{"CoverageSiteResp", string(data)}, " ")
}
