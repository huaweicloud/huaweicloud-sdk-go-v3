package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlavorCapacityRequest Request Object
type ShowFlavorCapacityRequest struct {
	FlavorId string `json:"flavor_id"`
}

func (o ShowFlavorCapacityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlavorCapacityRequest struct{}"
	}

	return strings.Join([]string{"ShowFlavorCapacityRequest", string(data)}, " ")
}
