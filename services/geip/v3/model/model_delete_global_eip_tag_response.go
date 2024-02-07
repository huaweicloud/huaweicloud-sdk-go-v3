package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGlobalEipTagResponse Response Object
type DeleteGlobalEipTagResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGlobalEipTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGlobalEipTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteGlobalEipTagResponse", string(data)}, " ")
}
