package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEquipmentInterfaceNameResponse Response Object
type ListEquipmentInterfaceNameResponse struct {

	// 设备接口名字列表
	InterfaceNames *[]string `json:"interface_names,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListEquipmentInterfaceNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEquipmentInterfaceNameResponse struct{}"
	}

	return strings.Join([]string{"ListEquipmentInterfaceNameResponse", string(data)}, " ")
}
