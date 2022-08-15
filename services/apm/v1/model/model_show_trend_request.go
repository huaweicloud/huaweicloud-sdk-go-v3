package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTrendRequest struct {

	// 应用id
	XBusinessId int64 `json:"x-business-id"`

	Body *TrendParam `json:"body,omitempty"`
}

func (o ShowTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowTrendRequest", string(data)}, " ")
}
