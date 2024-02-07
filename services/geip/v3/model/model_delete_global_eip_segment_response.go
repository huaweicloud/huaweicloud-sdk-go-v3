package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGlobalEipSegmentResponse Response Object
type DeleteGlobalEipSegmentResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGlobalEipSegmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGlobalEipSegmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteGlobalEipSegmentResponse", string(data)}, " ")
}
