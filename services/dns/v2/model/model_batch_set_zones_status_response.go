package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetZonesStatusResponse Response Object
type BatchSetZonesStatusResponse struct {
	Links *Link `json:"links,omitempty"`

	// 设置zone的列表响应。
	Zones *[]ZoneData `json:"zones,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchSetZonesStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetZonesStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchSetZonesStatusResponse", string(data)}, " ")
}
