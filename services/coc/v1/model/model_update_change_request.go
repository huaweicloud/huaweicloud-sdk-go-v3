package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateChangeRequest Request Object
type UpdateChangeRequest struct {

	// 变更单id
	ChangeId string `json:"change_id"`

	Body *CocUpdateChangeRequestBodyV2 `json:"body,omitempty"`
}

func (o UpdateChangeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateChangeRequest struct{}"
	}

	return strings.Join([]string{"UpdateChangeRequest", string(data)}, " ")
}
