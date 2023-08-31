package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEquipmentResponse Response Object
type DeleteEquipmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEquipmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEquipmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteEquipmentResponse", string(data)}, " ")
}
