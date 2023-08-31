package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEquipmentEsnResponse Response Object
type UpdateEquipmentEsnResponse struct {

	// esn
	Esn            *string `json:"esn,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEquipmentEsnResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentEsnResponse struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentEsnResponse", string(data)}, " ")
}
