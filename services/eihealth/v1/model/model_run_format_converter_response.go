package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunFormatConverterResponse Response Object
type RunFormatConverterResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunFormatConverterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunFormatConverterResponse struct{}"
	}

	return strings.Join([]string{"RunFormatConverterResponse", string(data)}, " ")
}
