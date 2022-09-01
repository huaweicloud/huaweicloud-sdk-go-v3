package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSpanSearchRequest struct {

	// 应用id
	XBusinessId int64 `json:"x-business-id" xml:"x-business-id"`

	Body *TraceSearchParam `json:"body,omitempty" xml:"body"`
}

func (o ShowSpanSearchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpanSearchRequest struct{}"
	}

	return strings.Join([]string{"ShowSpanSearchRequest", string(data)}, " ")
}
