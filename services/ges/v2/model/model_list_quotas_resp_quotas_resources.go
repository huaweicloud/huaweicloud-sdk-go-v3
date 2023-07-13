package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListQuotasRespQuotasResources struct {

	// 类型。  取值范围：  - \"graph\" - \"backup\" - \"metadata\"
	Type *string `json:"type,omitempty"`

	// 图的可用个数。
	Available *int32 `json:"available,omitempty"`

	// 边的可用个数。type为graph时此值有效。
	EdgeVolume *int32 `json:"edge_volume,omitempty"`
}

func (o ListQuotasRespQuotasResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotasRespQuotasResources struct{}"
	}

	return strings.Join([]string{"ListQuotasRespQuotasResources", string(data)}, " ")
}
