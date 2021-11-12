package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 保护实例变更规格请求数据结构
type ResizeProtectedInstanceRequestParams struct {
	// 变更规格后，生产站点云服务器和容灾站点云服务器的flavor ID。可通过查询云服务器规格变更支持列表接口获取。 说明：系统支持同时变更生产站点云服务器和容灾站点云服务器的规格。如需同时变更，请使用flavorRef参数，变更规格后，生产站点云服务器和容灾站点云服务器的规格相同。

	FlavorRef *string `json:"flavorRef,omitempty"`
	// 变更规格后，生产站点云服务器的flavor ID。可通过查询云服务器规格变更支持列表接口获取。 说明：系统支持仅变更生产站点云服务器的规格。此时，请使用production_flavorRef参数。当flavorRef参数有值时，production_flavorRef参数不生效。

	ProductionFlavorRef *string `json:"production_flavorRef,omitempty"`
	// 变更规格后，容灾站点云服务器的flavor ID。可通过查询云服务器规格变更支持列表接口获取。 说明：系统支持仅变更容灾站点云服务器的规格。此时，请使用dr_flavorRef参数。当flavorRef参数有值时，dr_flavorRef参数不生效。

	DrFlavorRef *string `json:"dr_flavorRef,omitempty"`
	// 新生产站点专属主机ID。 说明：生产站点云服务器在专属主机上时，变更规格需要指定此参数。可以指定为生产站点云服务器当前所在专属主机ID或其他专属主机ID。

	ProductionDedicatedHostId *string `json:"production_dedicated_host_id,omitempty"`
	// 新容灾站点专属主机ID。 说明：容灾站点云服务器在专属主机上时，变更规格需要指定此参数。可以指定为容灾站点云服务器当前所在专属主机ID或其他专属主机ID。

	DrDedicatedHostId *string `json:"dr_dedicated_host_id,omitempty"`
}

func (o ResizeProtectedInstanceRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeProtectedInstanceRequestParams struct{}"
	}

	return strings.Join([]string{"ResizeProtectedInstanceRequestParams", string(data)}, " ")
}
