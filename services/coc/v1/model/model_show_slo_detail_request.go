package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSloDetailRequest Request Object
type ShowSloDetailRequest struct {

	// SLO的ID
	SloId string `json:"slo_id"`
}

func (o ShowSloDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSloDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowSloDetailRequest", string(data)}, " ")
}
