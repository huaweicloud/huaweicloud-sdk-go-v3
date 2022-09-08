package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Indicates the resource type
type Quota struct {

	// 类型。  取值范围： - \"graph\" - \"backup\" - \"metadata\"
	Type string `json:"type"`

	// 图的可用个数。
	Available int32 `json:"available"`

	// 边的可用个数。type为graph时此值有效。
	EdgeVolume *int32 `json:"edgeVolume,omitempty"`
}

func (o Quota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quota struct{}"
	}

	return strings.Join([]string{"Quota", string(data)}, " ")
}
