package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMcpServerResponse Response Object
type DeleteMcpServerResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMcpServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMcpServerResponse struct{}"
	}

	return strings.Join([]string{"DeleteMcpServerResponse", string(data)}, " ")
}
