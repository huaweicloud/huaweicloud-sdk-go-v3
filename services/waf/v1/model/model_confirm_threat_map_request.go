package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmThreatMapRequest Request Object
type ConfirmThreatMapRequest struct {
}

func (o ConfirmThreatMapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmThreatMapRequest struct{}"
	}

	return strings.Join([]string{"ConfirmThreatMapRequest", string(data)}, " ")
}
