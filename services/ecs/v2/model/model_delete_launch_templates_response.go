package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLaunchTemplatesResponse Response Object
type DeleteLaunchTemplatesResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteLaunchTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLaunchTemplatesResponse struct{}"
	}

	return strings.Join([]string{"DeleteLaunchTemplatesResponse", string(data)}, " ")
}
