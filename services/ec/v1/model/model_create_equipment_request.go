package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEquipmentRequest Request Object
type CreateEquipmentRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	Body *EquipmentActivate `json:"body,omitempty"`
}

func (o CreateEquipmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEquipmentRequest struct{}"
	}

	return strings.Join([]string{"CreateEquipmentRequest", string(data)}, " ")
}
