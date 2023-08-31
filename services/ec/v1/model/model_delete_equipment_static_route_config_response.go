package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEquipmentStaticRouteConfigResponse Response Object
type DeleteEquipmentStaticRouteConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEquipmentStaticRouteConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEquipmentStaticRouteConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteEquipmentStaticRouteConfigResponse", string(data)}, " ")
}
