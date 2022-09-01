package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CatalogEndpoints struct {

	// 终端节点ID。
	Id string `json:"id" xml:"id"`

	// 终端节点平面，public表示为公开。
	Interface string `json:"interface" xml:"interface"`

	// 终端节点所属区域。
	Region string `json:"region" xml:"region"`

	// 终端节点所属区域的ID。
	RegionId string `json:"region_id" xml:"region_id"`

	// 终端节点的地址。
	Url string `json:"url" xml:"url"`
}

func (o CatalogEndpoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogEndpoints struct{}"
	}

	return strings.Join([]string{"CatalogEndpoints", string(data)}, " ")
}
