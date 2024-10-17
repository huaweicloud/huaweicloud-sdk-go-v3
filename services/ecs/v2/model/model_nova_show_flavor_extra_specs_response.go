package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaShowFlavorExtraSpecsResponse Response Object
type NovaShowFlavorExtraSpecsResponse struct {

	// 描述弹性云服务器规格的键值对。
	ExtraSpecs     map[string]string `json:"extra_specs,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o NovaShowFlavorExtraSpecsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaShowFlavorExtraSpecsResponse struct{}"
	}

	return strings.Join([]string{"NovaShowFlavorExtraSpecsResponse", string(data)}, " ")
}
