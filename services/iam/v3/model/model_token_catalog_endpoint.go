package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type TokenCatalogEndpoint struct {
	// 终端节点的URL。

	Url string `json:"url"`
	// 终端节点所属区域。

	Region string `json:"region"`
	// 终端节点所属区域ID。

	RegionId string `json:"region_id"`
	// 接口类型，描述接口在该终端节点的可见性。值为“public”，表示该接口为公开接口。

	Interface string `json:"interface"`
	// 终端节点ID。

	Id string `json:"id"`
}

func (o TokenCatalogEndpoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TokenCatalogEndpoint struct{}"
	}

	return strings.Join([]string{"TokenCatalogEndpoint", string(data)}, " ")
}
