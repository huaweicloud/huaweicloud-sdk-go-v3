package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Cancel2DDigitalHumanVideoResponse Response Object
type Cancel2DDigitalHumanVideoResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o Cancel2DDigitalHumanVideoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cancel2DDigitalHumanVideoResponse struct{}"
	}

	return strings.Join([]string{"Cancel2DDigitalHumanVideoResponse", string(data)}, " ")
}
