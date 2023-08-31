package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEquipmentLanConfigResponse Response Object
type DeleteEquipmentLanConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEquipmentLanConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEquipmentLanConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteEquipmentLanConfigResponse", string(data)}, " ")
}
