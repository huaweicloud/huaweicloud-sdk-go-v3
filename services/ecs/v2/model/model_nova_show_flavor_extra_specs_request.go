package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaShowFlavorExtraSpecsRequest Request Object
type NovaShowFlavorExtraSpecsRequest struct {

	// 规格id。
	FlavorId string `json:"flavor_id"`
}

func (o NovaShowFlavorExtraSpecsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaShowFlavorExtraSpecsRequest struct{}"
	}

	return strings.Join([]string{"NovaShowFlavorExtraSpecsRequest", string(data)}, " ")
}
