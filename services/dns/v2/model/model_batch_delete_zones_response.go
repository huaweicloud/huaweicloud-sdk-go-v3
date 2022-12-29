package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteZonesResponse struct {

	// 待删除zone类型，当前仅支持 public 或 private
	Zones *[]ZoneData `json:"zones,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteZonesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteZonesResponse", string(data)}, " ")
}
