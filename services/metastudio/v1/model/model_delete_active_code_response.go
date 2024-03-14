package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteActiveCodeResponse Response Object
type DeleteActiveCodeResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteActiveCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteActiveCodeResponse struct{}"
	}

	return strings.Join([]string{"DeleteActiveCodeResponse", string(data)}, " ")
}
