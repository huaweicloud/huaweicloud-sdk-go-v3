package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowControlsForAccountRequest Request Object
type ShowControlsForAccountRequest struct {

	// 纳管账号ID。
	ManagedAccountId string `json:"managed_account_id"`

	// 控制策略ID。
	ControlId string `json:"control_id"`
}

func (o ShowControlsForAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowControlsForAccountRequest struct{}"
	}

	return strings.Join([]string{"ShowControlsForAccountRequest", string(data)}, " ")
}
