package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Endpoint struct {

	// 终端节点所属服务的ID。
	ServiceId string `json:"service_id" xml:"service_id"`

	// 终端节点所属区域的ID。
	RegionId string `json:"region_id" xml:"region_id"`

	Links *Links `json:"links" xml:"links"`

	// 终端节点ID。
	Id string `json:"id" xml:"id"`

	// 终端节点平面。
	Interface string `json:"interface" xml:"interface"`

	// 终端节点所属区域。
	Region string `json:"region" xml:"region"`

	// 终端节点的地址。
	Url string `json:"url" xml:"url"`

	// 终端节点是否可用。
	Enabled bool `json:"enabled" xml:"enabled"`
}

func (o Endpoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Endpoint struct{}"
	}

	return strings.Join([]string{"Endpoint", string(data)}, " ")
}
