package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIRackRequest Request Object
type UpdateIRackRequest struct {

	// iRack 唯一标识符
	IrackId string `json:"irack_id"`

	Body *IRackRequest `json:"body,omitempty"`
}

func (o UpdateIRackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIRackRequest struct{}"
	}

	return strings.Join([]string{"UpdateIRackRequest", string(data)}, " ")
}
