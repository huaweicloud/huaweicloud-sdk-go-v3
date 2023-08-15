package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSpanSearchRequest Request Object
type ShowSpanSearchRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *TraceSearchParam `json:"body,omitempty"`
}

func (o ShowSpanSearchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpanSearchRequest struct{}"
	}

	return strings.Join([]string{"ShowSpanSearchRequest", string(data)}, " ")
}
