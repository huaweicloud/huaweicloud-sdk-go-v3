package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChangeOrderRequest Request Object
type CreateChangeOrderRequest struct {

	// 桌面id。
	DesktopId string `json:"desktop_id"`

	Body *CreateChangeOrderReq `json:"body,omitempty"`
}

func (o CreateChangeOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChangeOrderRequest struct{}"
	}

	return strings.Join([]string{"CreateChangeOrderRequest", string(data)}, " ")
}
