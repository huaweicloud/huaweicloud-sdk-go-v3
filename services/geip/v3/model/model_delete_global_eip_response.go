package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGlobalEipResponse Response Object
type DeleteGlobalEipResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGlobalEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGlobalEipResponse struct{}"
	}

	return strings.Join([]string{"DeleteGlobalEipResponse", string(data)}, " ")
}
