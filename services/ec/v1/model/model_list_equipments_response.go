package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEquipmentsResponse Response Object
type ListEquipmentsResponse struct {

	// 设备基本信息列表
	Equipments     *[]EquipmentItem `json:"equipments,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListEquipmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEquipmentsResponse struct{}"
	}

	return strings.Join([]string{"ListEquipmentsResponse", string(data)}, " ")
}
