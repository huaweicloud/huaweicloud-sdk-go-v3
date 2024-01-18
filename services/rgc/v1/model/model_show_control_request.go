package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowControlRequest Request Object
type ShowControlRequest struct {

	// 控制策略ID。
	ControlId string `json:"control_id"`
}

func (o ShowControlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowControlRequest struct{}"
	}

	return strings.Join([]string{"ShowControlRequest", string(data)}, " ")
}
