package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowClobDetailRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *GetClobDetailParam `json:"body,omitempty"`
}

func (o ShowClobDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClobDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowClobDetailRequest", string(data)}, " ")
}
