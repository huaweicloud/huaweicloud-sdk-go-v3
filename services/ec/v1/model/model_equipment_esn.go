package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EquipmentEsn struct {

	// esn
	Esn string `json:"esn"`
}

func (o EquipmentEsn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EquipmentEsn struct{}"
	}

	return strings.Join([]string{"EquipmentEsn", string(data)}, " ")
}
