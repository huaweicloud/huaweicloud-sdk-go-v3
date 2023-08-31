package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEquipmentInfoBody struct {

	// 设备名称
	Name string `json:"name"`
}

func (o UpdateEquipmentInfoBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentInfoBody struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentInfoBody", string(data)}, " ")
}
