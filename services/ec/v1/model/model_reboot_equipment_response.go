package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootEquipmentResponse Response Object
type RebootEquipmentResponse struct {

	// 成功信息
	SuccessMsg     *string `json:"success_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RebootEquipmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootEquipmentResponse struct{}"
	}

	return strings.Join([]string{"RebootEquipmentResponse", string(data)}, " ")
}
