package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIegRequest Request Object
type UpdateIegRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	Body *UpdateIegRequestBody `json:"body,omitempty"`
}

func (o UpdateIegRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIegRequest struct{}"
	}

	return strings.Join([]string{"UpdateIegRequest", string(data)}, " ")
}
