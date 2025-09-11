package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlavorCapacityResponse Response Object
type ShowFlavorCapacityResponse struct {
	Resources      *[]QueryFlavorCapacityAzInfo `json:"resources,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowFlavorCapacityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlavorCapacityResponse struct{}"
	}

	return strings.Join([]string{"ShowFlavorCapacityResponse", string(data)}, " ")
}
